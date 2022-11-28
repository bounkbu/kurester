FROM golang:1.19.3-alpine3.15 AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

EXPOSE 8888
CMD ["/app/main"]
