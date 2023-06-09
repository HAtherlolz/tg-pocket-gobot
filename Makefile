.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	sudo docker build -t tg_bot_golang:v0.1 .

start-container:
	sudo docker run --env-file .env -p 8080:8080 tg_bot_golang:v0.1
	