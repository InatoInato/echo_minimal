# Build stage
FROM golang:1.23.10 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ❗ Статическая сборка
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

# Final stage
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/app .
COPY .env .env

EXPOSE 8080

CMD ["./app"]
