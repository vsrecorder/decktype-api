SHELL := /bin/bash

.PHONY: build
build:
	go mod tidy
	go build -o bin/main main.go

.PHONY: run
run:
	go mod tidy
	go run main.go

.PHONY: image
image:
	go mod tidy
	docker build -t vsrecorder/decktype-api:latest . && docker push vsrecorder/decktype-api:latest

.PHONY: deploy
deploy:
	docker compose pull && docker compose down && docker compose up -d
