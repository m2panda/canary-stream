FROM timberio/vector:0.56.0-alpine

COPY ./config/clean_logs.sh /usr/local/bin/clean_logs.sh
COPY ./config/crontab.txt /etc/crontabs/root

RUN chmod +x /usr/local/bin/clean_logs.sh
RUN apk add --no-cache bash dcron

ENTRYPOINT ["sh", "-c", "crond -b && exec vector --config /etc/vector/vector.yaml"]
