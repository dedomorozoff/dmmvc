# DMMVC Makefile

.PHONY: help build install clean run test cli swagger

# Default target
help:
	@echo "DMMVC - Available commands:"
	@echo.
	@echo Build commands:
	@echo   make build       - Build the CLI tool
	@echo   make cli         - Build CLI tool only
	@echo   make server      - Build server binary
	@echo.
	@echo Installation:
	@echo   make install     - Install CLI globally (Windows specific)
	@echo   make install-go  - Install CLI via go install (cross-platform)
	@echo.
	@echo Development:
	@echo   make run         - Run the web server
	@echo   make test        - Run tests
	@echo   make swagger     - Generate Swagger documentation
	@echo.
	@echo Utilities:
	@echo   make clean       - Remove built binaries
	@echo.
	@echo Quick start:
	@echo   1. make install-go
	@echo   2. dmmvc --help
	@echo   3. make run

# Build CLI tool
build: cli

cli:
	@echo "Building DMMVC CLI..."
	@go build -o dmmvc.exe cmd/cli/main.go
	@echo "✓ CLI built successfully: dmmvc.exe"

# Install CLI globally (Windows)
install: cli
	@echo "Installing DMMVC CLI..."
	@if not defined GOPATH (echo Error: GOPATH not set && exit /b 1)
	@if not exist "%GOPATH%\bin" mkdir "%GOPATH%\bin"
	@copy /Y dmmvc.exe "%GOPATH%\bin\dmmvc.exe" >nul
	@echo ✓ CLI installed to %GOPATH%\bin\dmmvc.exe
	@echo You can now use 'dmmvc' command globally
	@echo.
	@echo Make sure %GOPATH%\bin is in your PATH

# Install using go install (cross-platform)
install-go:
	@echo "Installing DMMVC CLI via go install..."
	@go install ./cmd/cli
	@echo "✓ CLI installed successfully"
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
