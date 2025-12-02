#!/bin/bash
# DMMVC Development Helper Script

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Functions
print_success() {
    echo -e "${GREEN}✓${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

print_info() {
    echo -e "${YELLOW}→${NC} $1"
}

# Check if .env exists
check_env() {
    if [ ! -f ".env" ]; then
        print_error ".env file not found"
        echo "Creating .env from .env.example..."
        if [ -f ".env.example" ]; then
            cp .env.example .env
            print_success ".env created"
        else
            print_error ".env.example not found"
            exit 1
        fi
    fi
}

# Install dependencies
install_deps() {
    print_info "Installing dependencies..."
    go mod download
    go mod tidy
    print_success "Dependencies installed"
}

# Build CLI
build_cli() {
    print_info "Building CLI..."
    go build -o dmmvc cmd/cli/main.go
    print_success "CLI built: ./dmmvc"
}

# Build server
build_server() {
    print_info "Building server..."
    go build -o server cmd/server/main.go
    print_success "Server built: ./server"
}

# Run server
run_server() {
    check_env
    print_info "Starting server..."
    go run cmd/server/main.go
}

# Run tests
run_tests() {
    print_info "Running tests..."
    go test ./... -v
}

# Format code
format_code() {
    print_info "Formatting code..."
    go fmt ./...
    print_success "Code formatted"
}

# Lint code
lint_code() {
    print_info "Linting code..."
    if command -v golangci-lint &> /dev/null; then
        golangci-lint run
        print_success "Linting complete"
    else
        print_error "golangci-lint not installed"
        echo "Install: https://golangci-lint.run/usage/install/"
    fi
}

# Generate Swagger docs
gen_swagger() {
    print_info "Generating Swagger documentation..."
    if command -v swag &> /dev/null; then
        swag init -g cmd/server/main.go -o docs/swagger --parseDependency --parseInternal
        print_success "Swagger docs generated"
    else
        print_error "swag not installed"
        echo "Install: go install github.com/swaggo/swag/cmd/swag@latest"
    fi
}

# Clean build artifacts
clean() {
    print_info "Cleaning build artifacts..."
    rm -f dmmvc server
    rm -f *.exe
    rm -f *.db
    rm -f *.log
    print_success "Clean complete"
}

# Show help
show_help() {
    echo "DMMVC Development Helper"
    echo ""
    echo "Usage: ./dev.sh [command]"
    echo ""
    echo "Commands:"
    echo "  install     - Install dependencies"
    echo "  build       - Build CLI and server"
    echo "  cli         - Build CLI only"
    echo "  server      - Build server only"
    echo "  run         - Run development server"
    echo "  test        - Run tests"
    echo "  fmt         - Format code"
    echo "  lint        - Lint code"
    echo "  swagger     - Generate Swagger docs"
    echo "  clean       - Clean build artifacts"
    echo "  help        - Show this help"
    echo ""
}

# Main
case "${1:-help}" in
    install)
        install_deps
        ;;
    build)
        build_cli
        build_server
        ;;
    cli)
        build_cli
        ;;
    server)
        build_server
        ;;
    run)
        run_server
        ;;
    test)
        run_tests
        ;;
    fmt|format)
        format_code
        ;;
    lint)
        lint_code
        ;;
    swagger)
        gen_swagger
        ;;
    clean)
        clean
        ;;
    help|*)
        show_help
        ;;
esac
