#!/bin/bash
# DMMVC Installation Script for Linux/macOS

set -e

echo "========================================"
echo "DMMVC Framework Installation"
echo "========================================"
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "[ERROR] Go is not installed or not in PATH"
    echo "Please install Go from https://golang.org/dl/"
    exit 1
fi

echo "[1/4] Checking Go installation..."
go version
echo ""

echo "[2/4] Installing dependencies..."
go mod download
echo ""

echo "[3/4] Building CLI tool..."
go build -o dmmvc cmd/cli/main.go
echo ""

echo "[4/4] Installing CLI globally..."
GOPATH=${GOPATH:-$(go env GOPATH)}
GOBIN=${GOBIN:-$GOPATH/bin}

if [ ! -d "$GOBIN" ]; then
    mkdir -p "$GOBIN"
fi

cp dmmvc "$GOBIN/dmmvc"
chmod +x "$GOBIN/dmmvc"
echo "CLI installed to: $GOBIN/dmmvc"
echo ""

echo "========================================"
echo "Installation Complete!"
echo "========================================"
echo ""
echo "To use DMMVC CLI globally, make sure your Go bin directory is in PATH:"
echo "  export PATH=\$PATH:$GOBIN"
echo ""
echo "Add this line to your ~/.bashrc, ~/.zshrc, or ~/.profile to make it permanent"
echo ""
echo "Quick start:"
echo "  1. dmmvc --help              - Show available commands"
echo "  2. dmmvc make:crud Product   - Generate CRUD for Product"
echo "  3. go run cmd/server/main.go - Start the server"
echo ""
echo "Documentation: https://github.com/dedomorozoff/dmmvc"
echo ""
