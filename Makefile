PROJ_PATH=${CURDIR}
PROTO_DEST=./src/proto
.DEFAULT_GOAL := help

.PHONY: install-view
install-view: ## install node_modules
	cd view && yarn

.PHONY: build-view
build-view: ## build client for production
	cd view && yarn build

.PHONY: build-api
build-api: ## build api for production
	cd api/cmd/api && go build

.PHONY: docker-build
docker-build: ## start docker compose
	docker-compose build

.PHONY: up
up: ## start docker compose
	docker-compose up -d

.PHONY: down
down: ## start docker compose
	docker-compose down

.PHONY: dockerhub-image
dockerhub-image: ## start docker with dockerhub image
	docker-compose -f docker-compose-deploy.yml up

.PHONY: dockerhub-image-down
dockerhub-image-down: ## stop docker with dockerhub image
	docker-compose -f docker-compose-deploy.yml down

.PHONY: mod-vendor
mod-vendor: ## Download, verify and vendor dependencies
	cd api && go mod tidy && go mod download && go mod verify

.PHONY: linter
linter: ## Run linter
	cd api && golangci-lint run

.PHONY: update-go-deps
update-go-deps:
	cd api && go get -u ./...

.PHONY: update-yarn-deps
go-go-deps:
	yarn upgrade --latest

# https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## Shows the help
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
        awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ''
