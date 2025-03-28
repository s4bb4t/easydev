FROM alpine:latest

RUN apk add --no-cache nginx

COPY ./nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
EXPOSE 443

CMD ["sh", "-c", "nginx"]
