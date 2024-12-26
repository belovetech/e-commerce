GOOSE_CMD=goose -dir database/schema
SQLC_CMD=sqlc
DB_TYPE=postgres


.PHONY: db-migrate, db-rollback, create-migration, db-generate, test, server

db-migrate:
	@echo "Migrating database..."
	$(GOOSE_CMD) $(DB_TYPE) $$DB_URL up

db-rollback:
	@echo "Rolling back database..."
	$(GOOSE_CMD) $(DB_TYPE) $$DB_URL down

create-migration:
	@if [ -z "$(name)" ]; then \
		echo "Please provide a name for the migration. usage: make create-migration name=migration_name"; \
		exit 1; \
	fi
	@echo "Creating migration..."
	$(GOOSE_CMD) create $(name) sql


db-generate:
	@echo "Generating database code..."
	$(SQLC_CMD) generate


test:
	@go test -v ./...


server:
	air
