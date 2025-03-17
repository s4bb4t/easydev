FROM alpine:latest

RUN apk add --no-cache nginx

COPY ./nginx.conf /etc/nginx/nginx.conf
COPY certs /usr/local/bin/certs

EXPOSE 80
EXPOSE 443

CMD ["sh", "-c", "nginx"]
