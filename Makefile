run:
	go run cmd/server/main.go

migrate:
	migrate -path database/postgres/migrations -database "postgres://root:root@host.docker.internal:5432/go_clean_architecture?sslmode=disable" up