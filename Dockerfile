FROM alpine:3.15

WORKDIR /root/
RUN apk --no-cache add ca-certificates

CMD "./api"
