.PHONY: build
build:
	go mod tidy
	go build -o bin/main main.go

.PHONY: run
run:
	go mod tidy
	go run main.go

.PHONY: docker-build-and-push
docker-build-and-push:
	go mod tidy
	docker build -t vsrecorder/decktype-api:latest . && docker push vsrecorder/decktype-api:latest

.PHONY: deploy
deploy:
	docker compose pull && docker compose down && docker compose up -d
