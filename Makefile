# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Database connection string
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Default value for n if not provided (e.g., make up n=1)
n ?= 1

# --- Migration Commands ---

# Run all pending up migrations
migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

# Run n up migrations
# Usage: make up n=3
up:
	migrate -path migrations -database "$(DB_URL)" -verbose up $(n)

# Rollback n migrations
# Usage: make down n=1
down:
	migrate -path migrations -database "$(DB_URL)" -verbose down $(n)

# Fix "dirty" database state
# Usage: make force v=1
force:
	migrate -path migrations -database "$(DB_URL)" force $(v)

# Check current migration version
version:
	migrate -path migrations -database "$(DB_URL)" version

# Create a new migration file
# Usage: make new-migration name=add_users_table
new-migration:
	migrate create -ext sql -dir migrations -seq $(name)

.PHONY: migrate-up up down force version new-migration
