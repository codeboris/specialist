.PHONY: run
.SILENT:

run: build start

build:
	go build -v ./cmd/api

start:
	./api

init-db:
	docker compose up -d --build

down-db:
	docker compose down --remove-orphans

db:
	docker compose exec postgres psql -U postgres -d restapi

migrate-add:
	migrate create -ext sql -dir migrations -seq $(NAME)

migrate-up:
	migrate -path migrations -database 'postgres://postgres:postgres@localhost:5432/restapi?sslmode=disable' up

migrate-down:
	migrate -path migrations -database 'postgres://postgres:postgres@localhost:5432/restapi?sslmode=disable' down

.DEFAULT_GOAL:= run