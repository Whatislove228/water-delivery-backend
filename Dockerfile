FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /water-delivery ./cmd/api

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /water-delivery /usr/local/bin/water-delivery

EXPOSE 8080

CMD ["water-delivery"]