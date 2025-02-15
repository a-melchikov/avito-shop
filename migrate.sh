#!/bin/bash
set -e  # Останавливаем выполнение при ошибке

echo "Выполняем миграции базы данных..."

migrate -path ./migrations -database "postgres://$DB_USER:$DB_PASSWORD@postgres:$DB_PORT/$DB_NAME?sslmode=disable" up

echo "Миграции успешно применены!"
