# build go binary first
FROM golang:latest AS builder
RUN mkdir -p /app
WORKDIR /app
ADD . .

# build
ENV GO111MODULE=on
RUN cd scrapper && CGO_ENABLED=0 GOOS=linux go build -o scrapper

# run binary via cronjob
FROM ubuntu:latest
RUN apt-get update && apt-get -y install cron ca-certificates
RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /app/scrapper /app
COPY --from=builder /app/config /app/config
COPY scrapper-cron /etc/cron.d/scrapper-cron
RUN chmod 0644 /etc/cron.d/scrapper-cron
RUN crontab /etc/cron.d/scrapper-cron
RUN touch /var/log/cron.log
CMD ["cron", "-f"]