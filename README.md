# A Simple Go Clean Architecture

This is an example of implementation of Clean Architecture in Golang.

## Installation

This project requires [Go](https://golang.org/) v1.20+ to run.

```bash
# Clone this project
$ git clone https://github.com/arvians-id/go-clean-architecture.git

# Move to project directory
$ cd go-clean-architecture

# Copy .env.example to .env
$ cp .env.example .env

# Install dependencies
$ go mod download
# or
$ go mod tidy
```

Please refer to the [documentation](https://github.com/golang-migrate/migrate) for installation of migrate.
```bash
# Migrate tables
$ migrate -path database/postgres/migrations -database "postgres://root:root@host.docker.internal:5432/go_clean_architecture?sslmode=disable" up
```

## Run Application

To run application, you can use makefile's command or ```go run cmd/server/main.go```
```bash
# Run database on docker
$ docker-compose up -d

# Run application
$ make run

# Execute the call with curl in another terminal
$ curl -X GET http://localhost:8080/api/users
```
