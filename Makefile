.PHONY: run
run:
	go mod tidy
	go run main.go

.PHONY: docker-build-and-push
docker-build-and-push:
	sudo docker build --no-cache -t vsrecorder/decktype-api:latest . && sudo docker push vsrecorder/decktype-api:latest