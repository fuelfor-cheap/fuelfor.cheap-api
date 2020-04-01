# alpine will reduce the size to roughly 10MB
FROM alpine:latest
# run gin in production
ENV GIN_MODE=release
RUN apk --no-cache add ca-certificates
WORKDIR /root

COPY main .
COPY config ./config
CMD ["./main"]