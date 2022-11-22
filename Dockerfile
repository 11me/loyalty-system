FROM golang:1.18.8-alpine AS builder
WORKDIR /opt/build
COPY . .
RUN go mod download && \
    go build -o main .

FROM alpine:3.16.0
WORKDIR /opt/app
COPY --from=builder /opt/build/main .
EXPOSE 8080 9000
CMD ["./main"]
