APP_NAME := $(shell grep -m1 module go.mod | awk '{print $$2}' | sed 's/module //')

# Makefile help

.PHONY: help
help: header usage options

.PHONY: header
header:
	@printf "\033[34mEnvironment\033[0m"
	@echo ""
	@printf "\033[34m---------------------------------------------------------------\033[0m"
	@echo ""
	@printf "\033[33m%-23s\033[0m" "APP_NAME"
	@printf "\033[35m%s\033[0m" $(APP_NAME)
	@echo ""
	@echo ""

.PHONY: usage
usage:
	@printf "\033[034mUsage\033[0m"
	@echo ""
	@printf "\033[34m---------------------------------------------------------------\033[0m"
	@echo ""
	@printf "\033[37m%-22s\033[0m %s\n" "make [options]"
	@echo ""

.PHONY: options
options:
	@printf "\033[34mOptions\033[0m"
	@echo ""
	@printf "\033[34m---------------------------------------------------------------\033[0m"
	@echo ""
	@perl -nle'print $& if m{^[a-zA-Z_-]+:.*?## .*$$}' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-22s\033[0m %s\n", $$1, $$2}'


# Makefile commands

.PHONY: run
run:  ## Start server
	@make start

.PHONE: run
watch: ## Start development server with hot reload
	@make start-dev

.PHONY: watch
start-dev:
	air

.PHONY: start
start:
	go run main.go

# Database commands

.PHONY: migrate
migrate: ## Run the migrations
	atlas migrate apply --env gorm

.PHONY: rollback
rollback: ## Rollback the migrations
	atlas migrate down --env gorm

.PHONY: generate-migration
generate-migration: ## Generate a new migration
	@printf "\033[33mEnter migration message: \033[0m"
	@read -r message; \
	atlas migrate diff --env gorm $$message
