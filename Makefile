# Load environment variables from .env file
include .env

MIGRATION_DIR = ./database/migration
DB_DSN = mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp\(${DB_HOST}:${DB_PORT}\)/${DB_DATABASE}

.PHONY: help migrate-up

help:
	@echo "Available targets:"
	@echo "	migrate-up		: Apply database migrations"
	@echo "	migrate-down	: Drop all migration's table"

migrate-up:
	migrate -database ${DB_DSN} -path ${MIGRATION_DIR} up

migrate-down:
	migrate -database ${DB_DSN} -path ${MIGRATION_DIR} down

migrate-drop:
	migrate -database ${DB_DSN} -path ${MIGRATION_DIR} drop

migrate-version:
	migrate -database ${DB_DSN} -path ${MIGRATION_DIR} version
