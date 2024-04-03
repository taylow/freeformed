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

.PHONY: docker
docker-build:
	docker build -t freeformed .
	
.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 freeformed

.PHONY: stack
stack:
	docker-compose up -d
	
.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: migrate
migrate:
	migrate -path sql/migration -database "postgresql://freeformed:freeformed@localhost:5432/freeformed?sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	migrate -path sql/migration -database "postgresql://freeformed:freeformed@localhost:5432/freeformed?sslmode=disable" down
