include .envrc
MIGRATION_PATH = ./cmd/migrate/migrations

.PHONY: run
run:
	@go run ./cmd/api

.PHONY: migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATION_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATION_PATH) -database=$(DATABASE_URL) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATION_PATH) -database=$(DATABASE_URL) down $(filter-out $@,$(MAKECMDGOALS))
