FROM alpine:3.15

WORKDIR /root/
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait && apk --no-cache add ca-certificates

CMD /wait && ./api
