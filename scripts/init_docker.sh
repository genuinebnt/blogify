#!/usr/bin/env bash
set -x
set -eo pipefail

export DB_USER=${POSTGRES_USER:=postgres}
export DB_PASSWORD=${POSTGRES_PASSWORD:=password}
export DB_NAME=${POSTGRES_DB:=blogify}
export DB_PORT=${POSTGRES_PORT:=5432}
export DB_HOST=${POSTGRES_HOST:=localhost}

export PORT=3000
export ENVIRONMENT=development

export DATABASE_URL=postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable

docker compose up -d

export PGPASSWORD=${DB_PASSWORD}
until psql -h ${DB_HOST} -U ${DB_USER} -p ${DB_PORT} -d 'postgres' -c '\q'; do
    >&2 echo "Postgres is still unavailable - sleeping"
    sleep 1
done

migrate -path=./migrations -database ${DATABASE_URL} up
