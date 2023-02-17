#!/bin/bash

set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
env=${APP_ENV:-production}

echo "$env"

if [ "$env" != "local" ]; then
    exit 1
fi

echo "Creating database ${POSTGRES_DB} and test db ....."

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE SCHEMA IF NOT EXISTS ${POSTGRES_SCHEMA};
  SET search_path To "${POSTGRES_SCHEMA}";
  CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
  CREATE DATABASE "${POSTGRES_DB}_test";
EOSQL

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "${POSTGRES_DB}_test" <<-EOSQL
  CREATE SCHEMA IF NOT EXISTS "${POSTGRES_SCHEMA}_test";
  SET search_path To "${POSTGRES_SCHEMA}_test";
  CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL