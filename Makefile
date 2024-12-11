.PHONY:
.SILENT:

run: build start

build:
	go build -v ./cmd/api

start:
	./api