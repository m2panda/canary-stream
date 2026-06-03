#!/bin/bash

# Rule to stop process if error on any part of script
set -e

# Global constants
ROOT_DIR="/docker-entrypoint-initdb.d"
FILES=(
  "schema/types"
  "schema/schema"

  "seeds/status"
  "seeds/genres"
)

# Iterate each sql file name and check if
# file exist in directory, then running it with
# psql and env variables
init_sql_db() {
  for FILE_NAME in "${FILES[@]}"; do
    local SQL_FILE="$ROOT_DIR/$FILE_NAME.sql"

    if [ -f "$SQL_FILE" ]; then
      echo "Running $FILE_NAME"

      psql -v ON_ERROR_STOP=1 \
        -U "$POSTGRES_USER" \
        -d "$POSTGRES_DB" \
        -f "$SQL_FILE"
    else
      echo "File missing $FILE_NAME"
    fi
  done
}

# Main funtion
init_sql_db
