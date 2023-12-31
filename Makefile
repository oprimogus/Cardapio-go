include .env
export


.PHONY: fmt

lint:
	@gofmt -s -w .


install:
	go mod tidy

docker:
	docker compose -f docker-compose.yaml up -d

sqlc_generate:
	sqlc generate -f internal/infra/database/sqlc/sqlc.yaml

run:
	go run main.go

migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir ${MIGRATION_SOURCE_URL} -seq $$name

migration_up: 
	migrate -path ${MIGRATION_SOURCE_URL} -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migration_down: 
	migrate -path ${MIGRATION_SOURCE_URL} -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down 1

migration_fix: 
	@read -p "Enter migration version: " version; \
	migrate -path ${MIGRATION_SOURCE_URL} -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" force $$version
