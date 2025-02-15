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
    bash \
    curl

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64 -o /usr/local/bin/migrate && \
    chmod +x /usr/local/bin/migrate

WORKDIR /app

COPY --from=builder /app/merch-shop ./
COPY .env ./.env
COPY migrations ./migrations
COPY migrate.sh ./migrate.sh

RUN chmod +x ./migrate.sh
RUN mkdir -p /app/logs && chmod 777 /app/logs

EXPOSE 8080

ENTRYPOINT ["bash", "-c", "./migrate.sh && ./merch-shop"]
