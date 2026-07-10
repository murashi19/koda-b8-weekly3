FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o ordering-system main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ordering-system /app/

CMD ["/app/ordering-system"]