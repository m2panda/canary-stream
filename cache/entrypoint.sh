#!/bin/sh

envsubst < /etc/valkey/valkey.conf.template > /etc/valkey/valkey.conf

exec valkey-server /etc/valkey/valkey.conf
