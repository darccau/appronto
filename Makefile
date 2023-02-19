GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test
GOFLAGS := -v 
LDFLAGS := -s -w

export APPRONTO_DB_DSN='postgres://appronto:appronto@localhost:54332/appronto?sslmode=disable'

all: 
	go run ./cmd/api

build:
	$(GOBUILD) $(GOFLAGS) -ldflags '$(LDFLAGS)' -o "appronto" cmd/api/main.go

test: 
	$(GOTEST) $(GOFLAGS) ./...

migrate:
	migrate -path ./migrations -database $(APPRONTO_DB_DSN) up

rollback:
	migrate -path ./migrations -database $(APPRONTO_DB_DSN) down

tidy:
	$(GOMOD) tidy
