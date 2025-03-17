FROM alpine:latest

COPY ./nginx.conf /etc/nginx/nginx.conf
COPY certs /usr/local/bin/certs

RUN apk add --no-cache nginx

EXPOSE 80
EXPOSE 443

CMD ["sh", "-c", "nginx"]
