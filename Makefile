.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t tg_bot_golang:v0.1 .

start-container:
	docker run --name tg_bot_golang -p 8000:8000 --env-file .env tg_bot_golang:v0.1