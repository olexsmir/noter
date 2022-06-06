FROM alpine:3.16.0

WORKDIR /root/
RUN apk --no-cache add ca-certificates

CMD "./api"
