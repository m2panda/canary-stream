FROM valkey/valkey:9.1-alpine

RUN mkdir -p /etc/valkey

WORKDIR /etc/valkey
COPY ./valkey.conf ./valkey.conf.template

WORKDIR /usr/local/bin
COPY ./entrypoint.sh .

RUN apk add --no-cache gettext
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["entrypoint.sh"]
