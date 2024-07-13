FROM golang:1.20-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp

FROM alpine:3.20 AS release
WORKDIR /app
COPY --from=builder /app/myapp .
EXPOSE 3000

ENTRYPOINT ["./myapp"]