# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server .

# Runtime stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 4000
CMD ["./server"]
