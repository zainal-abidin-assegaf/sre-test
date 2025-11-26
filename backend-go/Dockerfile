FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main

FROM alpine:3.22.2

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
