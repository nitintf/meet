SHELL := /bin/bash
PROJECT_ROOT_DIR := $(shell git rev-parse --show-toplevel)

SERVER_BIN_DIR = $(PROJECT_ROOT_DIR)/bin
SERVER_BUILD_DIR = $(PROJECT_ROOT_DIR)/build

.DEFAULT_GOAL := help

.PHONY: gomod
gomod: ## Run mod tidy and vendor
	go mod tidy -compat=1.20
	go mod vendor

.PHONY: setup-local
setup-local: ## Setup App locally
	@ sh ./scripts/setup-local.sh

.PHONY: client-dev
client-dev: ## Run Client Dev on localhost
	@ cd ./client; npm run dev

.PHONY: help
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
