FROM golang:1.23 as builder
WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o stream-key-manager main.go

FROM alpine:latest
COPY --from=builder /app/stream-key-manager /stream-key-manager
RUN chmod 755 ./stream-key-manager
CMD ["./stream-key-manager"]

EXPOSE 8000