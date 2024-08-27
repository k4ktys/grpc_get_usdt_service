FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main ./cmd/main/main.go

FROM alpine:latest

COPY --from=builder /app/main /
COPY --from=builder /app/.env /
COPY --from=builder /app/migrations /migrations

CMD ["./main"]