FROM golang:1.15.4-alpine

RUN apk update && apk add gcc musl-dev git linux-headers make

RUN mkdir /app

WORKDIR /app

ADD . /app

RUN go build

COPY docker/docker-entrypoint.sh /app/docker-entrypoint.sh

RUN chmod +x /app/docker-entrypoint.sh

ENTRYPOINT [ "/app/docker-entrypoint.sh" ]