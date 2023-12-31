
<a name="v0.2.0"></a>
## [v0.2.0](https://github.com/miauw-social/auth/compare/v0.1.0...v0.2.0)

> 2023-07-27

### Chore

* **docker:** change dockerfile
* **go:** add dependency

### Feat

* **env:** add env parser and struct to provide easier access to env vars

### Perf

* **config:** return pointer to config struct


<a name="v0.1.0"></a>
## v0.1.0

> 2023-07-22

### Build

* **docker:** add Dockerfile
* **python:** add new version of miauw-base-service
* **python:** add requirements.txt

### Chore

* **changelog:** add changelog
* **docker:** add Dockerfile for go
* **env:** add .env.example
* **git:** update .gitignore
* **git:** add certs to gitignore
* **git:** add .gitignore
* **idea:** remove .idea folder

### Docs

* **readme:** create readme.md

### Feat

* basic project in go
* jwt for verification token
* jwt for verification token
* **auth:** add create, update, login and session methods
* **handler:** implement `auth.session.get`
* **handlers:** add `Session.ID` to Redis Cache.
* **handlers:** implement `auth.sessions.get`
* **handlers:** implement `auth.sessions.exists`
* **models:** add verification to user model
* **redis:** add classes for redis and session managment using redis
* **server:** add handler hostname
* **service:** implement event decorator & add verification
* **service:** login and initial now working
* **session:** add sessin model
* **user:** add user model
* **verification:** add verification handler
* **verification:** add listener for `auth.verify`

### Fix

* **jwt:** change key to use env var
* **service:** wrong queue name
* **service:** return right types in events
* **verify-jwt:** return claims as map

### Refactor

* **deps:** add redis
* **envvars:** add envvars to databases and security
* **testing:** add db as parameter

### Pull Requests

* Merge pull request [#5](https://github.com/miauw-social/auth/issues/5) from miauw-social/development

