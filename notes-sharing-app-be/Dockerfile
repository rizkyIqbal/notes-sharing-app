# Build stage
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 4000
CMD ["./server"]