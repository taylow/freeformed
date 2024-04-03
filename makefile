.PHONY: run
run:
	go run cmd/freeformed/main.go

.PHONY: start
start:
	go run cmd/freeformed/main.go

.PHONY: test
test:
	go test -v -cover ./...

build:
	go build -o bin/freeformed cmd/freeformed/main.go

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: migrate
migrate:
	migrate -path sql/migration -database "postgresql://freeformed:freeformed@localhost:5432/freeformed?sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	migrate -path sql/migration -database "postgresql://freeformed:freeformed@localhost:5432/freeformed?sslmode=disable" down

.PHONY: stack
stack:
	docker-compose up -d