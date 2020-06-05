FROM alpine:3.6

RUN apk add --no-cache \
        ca-certificates \
        bash \
    && rm -f /var/cache/apk/*

COPY bin/employeemanagementsystem /usr/local/bin/employeemanagementsystem

CMD ["/usr/local/bin/employeemanagementsystem"]
