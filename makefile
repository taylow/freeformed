
PACKAGES := $(shell go list ./...)
name := $(shell basename ${PWD})

all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## init: initialize project (make init module=github.com/user/project)
.PHONY: init
init:
	go mod init ${module}
	go install github.com/cosmtrek/air@latest
	asdf reshim golang

## vet: vet code
.PHONY: vet
vet:
	go vet $(PACKAGES)

## test: run unit tests
.PHONY: test
test:
	go test -race -cover $(PACKAGES)

## run: run project
.PHONY: run
run:
	go run cmd/${name}/main.go

## start: run project
.PHONY: start
start:
	go run cmd/${name}/main.go

## build: build project
.PHONY: build
build:
	go build -o bin/${name} cmd/${name}/main.go

## hotreload: build and run local project with hotreload
.PHONY: hotreload
hotreload: build
	air

## docker-build: build project into a docker container image
.PHONY: docker-build
docker-build: test
	GOPROXY=direct docker buildx build -t ${name} .

## docker-run: run project in a container
.PHONY: docker-run
docker-run:
	docker run -it --rm -p 8080:8080 ${name}

## stack: start docker stack
.PHONY: stack
stack:
	docker-compose up -d
	
## sqlc: generate sqlc code
.PHONY: sqlc
sqlc:
	sqlc generate

## migrate-up: run migrations
.PHONY: migrate-up
migrate-up:
	migrate -path sql/migration -database "postgresql://freeformed:freeformed@localhost:5432/freeformed?sslmode=disable" up

## migrate-down: rollback migrations
.PHONY: migrate-down
migrate-down:
	migrate -path sql/migration -database "postgresql://freeformed:freeformed@localhost:5432/freeformed?sslmode=disable" down

## css: build tailwindcss
.PHONY: css
css:
	tailwindcss -i css/input.css -o css/output.css --minify

## css-watch: watch build tailwindcss
.PHONY: css-watch
css-watch:
	tailwindcss -i css/input.css -o css/output.css --watch
