## Stage 1: build
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bot ./cmd/bot

## Stage 2: runtime
FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=builder /app/bot /app/bot

ENV BOT_TOKEN=""
ENV STICKER_FILE_ID=""

ENTRYPOINT ["/app/bot"]

