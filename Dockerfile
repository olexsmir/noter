FROM alpine:3.16.1

WORKDIR /root/
RUN apk --no-cache add ca-certificates

CMD "./api"
