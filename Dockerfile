# Builder
FROM golang:1.23.10 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o app ./cmd/main.go

# Final image
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/app .
COPY .env .env

CMD ["./app"]
