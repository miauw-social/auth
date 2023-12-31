package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"miauw.social/auth/config"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
	"miauw.social/auth/database"
	"miauw.social/auth/handlers"
)

func Serve(queueName string, handler func(*gorm.DB, []byte) (handlers.Response, error)) {
	cfg := config.GetConfig()
	conn, err := amqp.Dial(cfg.RabbitMQ)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel.")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare queue.")
	messages, err := ch.Consume(
		q.Name,
		"consumerTag",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer.")
	var forever chan struct{}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range messages {
			start := time.Now()
			r, err := handler(database.Conn(), d.Body)
			took := time.Since(start).Milliseconds()
			if err != nil {
				return
			}
			if d.ReplyTo != "" {
				jsonResponse, _ := json.Marshal(r)
				headers := make(map[string]interface{})
				headers["X-Process-Time"] = took
				hostname, _ := os.Hostname()
				headers["X-Worker"] = hostname
				err := ch.PublishWithContext(ctx,
					"",
					d.ReplyTo,
					false,
					false,
					amqp.Publishing{
						Headers:       headers,
						ContentType:   "application/json",
						CorrelationId: d.CorrelationId,
						Body:          []byte(jsonResponse),
					})
				if err != nil {
					fmt.Printf("Error: %v", err)
				}

			}
			d.Ack(true)

		}
	}()

	log.Printf(" [*] Waiting for %s. To exit press Ctrl-C.", queueName)
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf(" [!] %s: %s", err, msg)
	}
}
