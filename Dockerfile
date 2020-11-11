FROM golang:1.15.4-alpine

RUN apk update \
  && apk add gcc musl-dev git linux-headers make

RUN go get github.com/NPDSoftwareDev/consul-snapshot

COPY docker/docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

ENTRYPOINT [ "/docker-entrypoint.sh" ]
