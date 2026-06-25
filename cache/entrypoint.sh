#!/bin/sh

envsubst < /etc/valkey/valkey.conf.template > /usr/local/etc/valkey/valkey.conf

exec valkey-server /usr/local/etc/valkey/valkey.conf
