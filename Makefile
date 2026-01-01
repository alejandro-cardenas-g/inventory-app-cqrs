include .envrc

MIGRATIONS_PATH = ./internal/store/persistence/migrations

.PHONY: migration
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migration-up
migration-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(POSTGRES_URL) up

.PHONY: migration-down
migration-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(POSTGRES_URL) down

.PHONY: run
run:
	@go run cmd/api/main.go

.PHONY: run-worker
run-worker:
	@go run cmd/worker/main.go

.PHONY: sqlc
sqlc:
	@sqlc generate