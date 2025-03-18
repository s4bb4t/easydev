FROM golang:1.23-alpine

WORKDIR /

COPY bin/app /
COPY config/prod.yaml config/

EXPOSE 8083

ENTRYPOINT ["/app"]