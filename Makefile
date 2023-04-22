GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test
GOFLAGS := -v 
LDFLAGS := -s -w

export APPRONTO_DB_DSN='postgres://appronto:appronto@localhost:54332/appronto?sslmode=disable'
MIGRATION_NAME := $(firstword $(MAKECMDGOALS))

confirm:
 @echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

all: 
	go run ./cmd/api

build:
	$(GOBUILD) $(GOFLAGS) -ldflags '$(LDFLAGS)' -o "appronto" cmd/api/main.go

test: 
	$(GOTEST) $(GOFLAGS) ./...

migration: confirm
	@echo 'Creating migration files for ${name}..'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

migrate: cofirm
	@echo 'Creating migration files for ${name}..'
	migrate -path ./migrations -database $(APPRONTO_DB_DSN) up

rollback: confirm
	migrate -path ./migrations -database $(APPRONTO_DB_DSN) down

tidy:
	$(GOMOD) tidy
