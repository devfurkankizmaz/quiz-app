DB_URL=postgresql://furkan:427542@127.0.0.1:5433/my_db?sslmode=disable
MIGRATION_PATH=migrations

.PHONY: migrate-up migrate-down migrate-fix dev dev-down server gen test

server:
	go run cmd/app/main.go

test:
	go test -v ./...

create-migration:
	migrate create -ext sql -dir migrations -seq init_mg

migrate-up:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose down

migrate-fix:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose force <VERSION>

dev:
	docker-compose up -d

dev-down:
	docker-compose down