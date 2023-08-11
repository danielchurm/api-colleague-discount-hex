PWD=$(shell pwd)

DOCKER_TEST_INCLUDE_RUNTIME=1

-include smartshop-services-tools/docker/Makefile

## TODO: Need to define unique SERVICE_NAME (used for docker images etc.)
SERVICE_NAME?=api_go_template

PATH:=$(PWD)/bin:${PATH}:${HOME}/go/bin:/usr/local/bin
export PATH

SHELL:=env PATH=$(PATH) /bin/bash

export BUILDKIT_PROGRESS=plain
export DOCKER_BUILDKIT=1

export GOPRIVATE=github.com/JSainsburyPLC


MIGRATE_ARCH?=$(shell go env GOARCH)
MIGRATE_OS?=$(shell go env GOOS)
MIGRATE_VERSION?=v4.15.2

MOCKGEN_VERSION?=v1.6.0

DB_USER?=api_go_template
DB_NAME?=api_go_template
DB_PASSWORD?=password
DB_HOST?=api-go-template-postgres.db.internal
DB_PORT?=5432
DB_SSL_MODE?=disable
DB_TYPE?=postgres
DB_URL=${DB_TYPE}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}

.PHONY: first_time_setup
first_time_setup: docker_deps tools deps mocks

.PHONY: docker_deps
docker_deps:
	git submodule init
	git submodule update
	$(MAKE) pre_commit_install

.PHONY: deps
deps:
	go mod download

.PHONY: tools
tools:
	# Install mockgen
	go install github.com/golang/mock/mockgen@$(MOCKGEN_VERSION)
	# Download and extract the golang-migrate/migrate binary
	curl -L https://github.com/golang-migrate/migrate/releases/download/$(MIGRATE_VERSION)/migrate.$(MIGRATE_OS)-$(MIGRATE_ARCH).tar.gz | tar xvz migrate

.PHONY: clean_mocks
clean_mocks:
	rm -rf mocks/

.PHONY: clean_tests
clean_tests:
	go clean -testcache

.PHONY: clean
clean: clean_mocks
	rm -f smartshop-service migrate

.PHONY: mocks
mocks: clean_mocks
	go generate -v ./...

# For local development to create a new migration use 'make migrate_new name="name_of_my_migration"
.PHONY: migrate_new
migrate_new:
	./migrate -database $(DB_URL) create -ext sql -dir migrations -seq -digits 4 $(name)

# For use in pipes, ALWAYS FORWARD
.PHONY: migrate
migrate:
	./migrate -verbose -source file://migrations -database ${DB_URL} up

# For local development with postgres running in a docker container
.PHONY: migrate_down
migrate_down:
	./migrate -source file://migrations -database ${DB_URL} down 1

# For local development with postgres running in a docker container
.PHONY: migrate_up
migrate_up:
	./migrate -source file://migrations -database ${DB_URL} up 1

.PHONY: build
build:
	go build -o smartshop-service

.PHONY: test
test: clean_tests
	go test -timeout=5s -cover -race $$(go list ./... | grep -v e2e )

.PHONY: ci_test
ci_test:
	go test -timeout=5s -cover -race ./...

.PHONY: run
run:
	go run main.go