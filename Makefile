PWD:=$(shell pwd)
TARGET=...

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup: ## Install depeendent tools and setup project
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.8.1
	go install github.com/cosmtrek/air@latest
	go get -u github.com/kyoh86/richgo

.PHONY: gen-api
gen-api: ## Generate router and request type structs from OpenAPI spec.
	go generate api/gen_openapi.go

.PHONY: gen-orm
gen-orm: ## Genereate facebook/ent models
	go generate ent/generate.go

.PHONY: run
run: ## Run local API server
	air -c .air.toml

.PHONY: migrate
migrate: ## Run migration
	go run ./tools/migrate

.PHONY: migrate-test
migrate-test: ## Run migration
	APP_ENV=test go run tools/migrate

.PHONY: test
test: ## Run test API server
	APP_ENV=test richgo test -cover meh/${TARGET}

.PHONY: testv
testv: ## Run test API server with verbose
	APP_ENV=test richgo test -v -cover meh/${TARGET}

.PHONY: lint
lint: ## Run linter
	find . -name \*.go -not -path "./gen/*" -exec goimports -w -local meh {} +
	golangci-lint run

.PHONY: cover
cover: ## Check coverage
	APP_ENV=test go test ./... -coverprofile cover.out
	go tool cover -html=cover.out -o cover.html

reset_local_db: ## Reset local db
	mysql -uroot -proot -hdb -P3306 -e "DROP DATABASE IF EXISTS meh"
	mysql -uroot -proot -hdb -P3306 -e "CREATE DATABASE IF NOT EXISTS meh"
	mysql -uroot -proot -hdb -P3306 -e "GRANT ALL ON meh.* TO 'dev'@'%'"
