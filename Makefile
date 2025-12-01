# DMMVC Makefile

.PHONY: help build install clean run test cli swagger

# Default target
help:
	@echo "DMMVC - Available commands:"
	@echo "  make build      - Build the CLI tool"
	@echo "  make install    - Build and install CLI globally"
	@echo "  make clean      - Remove built binaries"
	@echo "  make run        - Run the web server"
	@echo "  make test       - Run tests"
	@echo "  make cli        - Build CLI tool only"
	@echo "  make swagger    - Generate Swagger documentation"

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

# Generate Swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	@swag init -g cmd/server/main.go -o docs/swagger --parseDependency --parseInternal
	@echo "✓ Swagger docs generated at docs/swagger/"
	@echo "Access at: http://localhost:8080/swagger/index.html"
