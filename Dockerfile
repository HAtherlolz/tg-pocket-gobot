FROM golang:1.15-alpine3.12 AS builder

COPY . /github.com/HAtherlolz/tg-pocket-gobot/
WORKDIR /github.com/HAtherlolz/tg-pocket-gobot/

RUN go mod download
RUN go build -o ./bin/ cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/HAtherlolz/tg-pocket-gobot/bin/bot .
COPY --from=0 /github.com/HAtherlolz/tg-pocket-gobot/configs configs/

EXPOSE 8000

CMD ["./bot"]