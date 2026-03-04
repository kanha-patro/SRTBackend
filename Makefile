# Makefile

# Makefile for Real-Time Shuttle Tracking Platform

# Variables
APP_NAME := srt
GO := go
GOBUILD := $(GO) build
GOTEST := $(GO) test
GOCLEAN := $(GO) clean
GOMOD := $(GO) mod
GOMIGRATE := migrate
MIGRATIONS_DIR := migrations

# Build the application
build:
	$(GOBUILD) -o $(APP_NAME) ./cmd/server/main.go

# Run the application
run: build
	./$(APP_NAME)

# Test the application
test:
	$(GOTEST) ./...

# Clean the build
clean:
	$(GOCLEAN)
	rm -f $(APP_NAME)

# Run database migrations
migrate:
	$(GOMIGRATE) -path $(MIGRATIONS_DIR) -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up

# Help
help:
	@echo "Makefile commands:"
	@echo "  build       - Build the application"
	@echo "  run         - Run the application"
	@echo "  test        - Run tests"
	@echo "  clean       - Clean the build"
	@echo "  migrate     - Run database migrations"
	@echo "  help        - Show this help message"