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
	docker compose exec postgres psql -U postgres

.DEFAULT_GOAL:= run