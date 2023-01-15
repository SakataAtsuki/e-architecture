#!/bin/bash
# db-migration

./migrate -database "postgres://${DB_HOST}/${DB_NAME}?sslmode=disable&user=${DB_USER}&password=${DB_PASSWORD}&port=${DB_PORT}" -path=./migrations up
