FROM valkey/valkey:9.1-alpine

COPY ./valkey.conf /etc/valkey/valkey.conf.template
COPY ./entrypoint.sh /usr/local/bin/entrypoint.sh

RUN apk add --no-cache gettext
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["entrypoint.sh"]
