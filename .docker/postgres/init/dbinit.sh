#!/bin/bash
set -e
export PGPASSWORD=postgres;
psql -v ON_ERROR_STOP=1 --username "postgres" <<-EOSQL
  CREATE DATABASE api;
  GRANT ALL PRIVILEGES ON DATABASE api TO "postgres";
EOSQL