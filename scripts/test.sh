#!/bin/bash
# DMMVC Test Runner Script

set -e

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

print_header() {
    echo ""
    echo "========================================"
    echo "$1"
    echo "========================================"
    echo ""
}

# Parse arguments
COVERAGE=false
VERBOSE=false
BENCH=false

while [[ $# -gt 0 ]]; do
    case $1 in
        -c|--coverage)
            COVERAGE=true
            shift
            ;;
        -v|--verbose)
            VERBOSE=true
            shift
            ;;
        -b|--bench)
            BENCH=true
            shift
            ;;
        -h|--help)
            echo "Usage: ./test.sh [options]"
            echo ""
            echo "Options:"
            echo "  -c, --coverage    Run tests with coverage"
            echo "  -v, --verbose     Verbose output"
            echo "  -b, --bench       Run benchmarks"
            echo "  -h, --help        Show this help"
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            echo "Use -h or --help for usage"
            exit 1
            ;;
    esac
done

print_header "DMMVC Test Runner"

# Run tests
if [ "$COVERAGE" = true ]; then
    echo "Running tests with coverage..."
    if [ "$VERBOSE" = true ]; then
        go test -v -cover -coverprofile=coverage.out ./...
    else
        go test -cover -coverprofile=coverage.out ./...
    fi
    
    echo ""
    echo "Coverage report:"
    go tool cover -func=coverage.out
    
    echo ""
    echo "Generate HTML coverage report:"
    echo "  go tool cover -html=coverage.out -o coverage.html"
    
elif [ "$BENCH" = true ]; then
    echo "Running benchmarks..."
    go test -bench=. -benchmem ./...
    
else
    echo "Running tests..."
    if [ "$VERBOSE" = true ]; then
        go test -v ./...
    else
        go test ./...
    fi
fi

echo ""
echo -e "${GREEN}✓ Tests completed${NC}"
echo ""
