#!/bin/sh

LOGS_DIR="/var/log/canary_stream"

find "$LOGS_DIR" -mindepth 1 -maxdepth 1 -type d -mtime +4 -exec rm -rf {} \;

echo "Limpieza interna de Vector completada: $(date)"
