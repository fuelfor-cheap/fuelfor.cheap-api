FROM golang:latest AS builder

# build directory
RUN mkdir -p /app
WORKDIR /app
COPY . .

# build
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# run image, alpine will reduce the size to roughly 10MB
FROM alpine:latest
# run gin in production
ENV GIN_MODE=release
RUN apk --no-cache add ca-certificates
WORKDIR /root

COPY --from=builder /app/main .
COPY config ./config
CMD ["./main"]