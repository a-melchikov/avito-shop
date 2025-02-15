FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w" \
    -o merch-shop cmd/app/main.go

FROM alpine:latest

RUN apk add --no-cache \
    nss \
    freetype \
    harfbuzz \
    ca-certificates \
    ttf-freefont \
    dumb-init \
    bash

WORKDIR /app

COPY --from=builder /app/merch-shop ./
COPY .env ./.env

RUN mkdir -p /app/logs && chmod 777 /app/logs

EXPOSE 8080

CMD ["./merch-shop"]