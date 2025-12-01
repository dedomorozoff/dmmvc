# DMMVC Makefile

.PHONY: help build install clean run test cli

# Default target
help:
	@echo "DMMVC - Available commands:"
	@echo "  make build      - Build the CLI tool"
	@echo "  make install    - Build and install CLI globally"
	@echo "  make clean      - Remove built binaries"
	@echo "  make run        - Run the web server"
	@echo "  make test       - Run tests"
	@echo "  make cli        - Build CLI tool only"

# Build CLI tool
build: cli

cli:
	@echo "Building DMMVC CLI..."
	@go build -o dmmvc.exe cmd/cli/main.go
	@echo "✓ CLI built successfully: dmmvc.exe"

# Install CLI globally (Windows)
install: cli
	@echo "Installing DMMVC CLI..."
	@copy dmmvc.exe %GOPATH%\bin\dmmvc.exe
	@echo "✓ CLI installed to %GOPATH%\bin\dmmvc.exe"
	@echo "You can now use 'dmmvc' command globally"

# Clean built binaries
clean:
	@echo "Cleaning..."
	@if exist dmmvc.exe del dmmvc.exe
	@if exist server.exe del server.exe
	@echo "✓ Clean complete"

# Run web server
run:
	@echo "Starting DMMVC server..."
	@go run cmd/server/main.go

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Build server
server:
	@echo "Building server..."
	@go build -o server.exe cmd/server/main.go
	@echo "✓ Server built successfully: server.exe"
