#!/bin/bash
# DMMVC Quick Start Script

set -e

echo "========================================"
echo "DMMVC Quick Start"
echo "========================================"
echo ""

# Check if .env exists
if [ ! -f ".env" ]; then
    echo "Creating .env file..."
    if [ -f ".env.example" ]; then
        cp .env.example .env
        echo "✓ .env created from .env.example"
    else
        echo "Creating default .env..."
        cat > .env << 'EOF'
PORT=8080
GIN_MODE=debug
DB_TYPE=sqlite
DB_DSN=dmmvc.db
SESSION_SECRET=change-this-in-production
LOG_LEVEL=info
LOG_FILE=dmmvc.log
DEBUG=true
EOF
        echo "✓ Default .env created"
    fi
    echo ""
fi

# Install dependencies if needed
if [ ! -d "vendor" ] && [ ! -f "go.sum" ]; then
    echo "Installing dependencies..."
    go mod download
    go mod tidy
    echo "✓ Dependencies installed"
    echo ""
fi

# Build CLI if not exists
if [ ! -f "dmmvc" ]; then
    echo "Building CLI tool..."
    go build -o dmmvc cmd/cli/main.go
    echo "✓ CLI built"
    echo ""
fi

echo "========================================"
echo "Setup Complete!"
echo "========================================"
echo ""
echo "Available commands:"
echo "  ./dmmvc --help           - Show CLI help"
echo "  ./dmmvc make:crud User   - Generate CRUD"
echo "  go run cmd/server/main.go - Start server"
echo ""
echo "Starting development server..."
echo ""

go run cmd/server/main.go
