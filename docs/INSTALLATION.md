# DMMVC Installation Guide

## System Requirements

- **Go**: version 1.20 or higher
- **Git**: for cloning the repository
- **OS**: Linux, macOS, or Windows

## Installation Methods

### 1. Automatic Installation (Recommended)

The easiest way to install DMMVC:

#### Linux/macOS

```bash
# Clone the repository
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc

# Make script executable
chmod +x scripts/install.sh

# Run installation script
./scripts/install.sh
```

#### Windows

```bash
# Clone the repository
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc

# Run installation script
scripts\install.bat
```

The script will automatically:
- Check Go installation
- Download dependencies
- Build CLI tool
- Install it globally

### 2. Installation via go install

If you already have Go installed:

```bash
# Install latest version
go install github.com/dedomorozoff/dmmvc/cmd/cli@latest

# CLI will be available as 'cli'
# Rename to 'dmmvc' if needed
```

### 3. Installation via Makefile

```bash
# Clone the repository
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc

# Install dependencies
go mod tidy

# Build and install CLI
make install-go

# Or use platform-specific installation (Windows)
make install
```

### 4. Manual Build

```bash
# Clone the repository
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc

# Install dependencies
go mod download

# Build CLI
go build -o dmmvc cmd/cli/main.go

# Build server
go build -o server cmd/server/main.go

# Add to PATH (optional)
# Copy dmmvc to $GOPATH/bin or any directory in PATH
sudo cp dmmvc /usr/local/bin/  # Linux/macOS
```

## Verify Installation

After installation, verify that CLI works:

```bash
# Check Go version
go version

# Check DMMVC CLI installation
dmmvc --help

# If command not found, check PATH
echo $PATH  # Linux/macOS
echo %PATH% # Windows
```

## Configure PATH

If `dmmvc` command is not found, add Go bin to PATH:

### Linux/macOS - Temporary (current session)

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Linux/macOS - Permanent

Add to `~/.bashrc`, `~/.zshrc`, or `~/.profile`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Then reload:

```bash
source ~/.bashrc  # or ~/.zshrc
```

### Windows - Temporary (current session)

```cmd
set PATH=%PATH%;%GOPATH%\bin
```

### Windows - Permanent

1. Open "System" → "Advanced system settings"
2. Click "Environment Variables"
3. In "System variables" find `Path`
4. Add path: `%USERPROFILE%\go\bin` or `%GOPATH%\bin`
5. Restart terminal

## Create Your First Project

### Using Project Template

#### Linux/macOS

```bash
# Make script executable
chmod +x scripts/create-project.sh

# Create new project
./scripts/create-project.sh my-app

# Navigate to project
cd my-app

# Install dependencies
go mod tidy

# Run server
go run cmd/server/main.go
```

#### Windows

```bash
# Create new project
scripts\create-project.bat my-app

# Navigate to project
cd my-app

# Install dependencies
go mod tidy

# Run server
go run cmd/server/main.go
```

### Using Existing Repository

```bash
# Clone DMMVC as base
git clone https://github.com/dedomorozoff/dmmvc my-app
cd my-app

# Remove Git history (optional)
rm -rf .git  # Linux/macOS
rmdir /s /q .git  # Windows
git init

# Configure project
# Edit go.mod, change module name
# Edit .env, configure settings

# Install dependencies
go mod tidy

# Run
go run cmd/server/main.go
```

## Use as Library

You can use DMMVC as a library in your project:

```bash
# Create new Go project
mkdir my-app
cd my-app
go mod init my-app

# Add DMMVC as dependency
go get github.com/dedomorozoff/dmmvc@latest
```

Then import needed packages:

```go
package main

import (
    "github.com/dedomorozoff/dmmvc/internal/database"
    "github.com/dedomorozoff/dmmvc/internal/logger"
    "github.com/dedomorozoff/dmmvc/internal/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    logger.Init()
    database.Init()
    
    r := gin.Default()
    routes.Setup(r)
    
    r.Run(":8080")
}
```

## Docker Installation

### Using Ready Image

```bash
# Build image
docker build -t dmmvc .

# Run container
docker run -p 8080:8080 dmmvc
```

### Docker Compose

```bash
# Run with PostgreSQL
docker-compose -f docker/docker-compose.postgres.yml up
```

## Development Helper Script

For Linux/macOS, use the development helper script:

```bash
# Make executable
chmod +x scripts/dev.sh

# Show available commands
./scripts/dev.sh help

# Install dependencies
./scripts/dev.sh install

# Build CLI and server
./scripts/dev.sh build

# Run development server
./scripts/dev.sh run

# Run tests
./scripts/dev.sh test

# Format code
./scripts/dev.sh fmt

# Generate Swagger docs
./scripts/dev.sh swagger
```

## Update

### Update CLI

```bash
# Via go install
go install github.com/dedomorozoff/dmmvc/cmd/cli@latest

# Or rebuild locally
cd dmmvc
git pull
make install-go
```

### Update Project Dependencies

```bash
# Update all dependencies
go get -u ./...
go mod tidy

# Update specific package
go get -u github.com/gin-gonic/gin
```

## Uninstall

### Linux/macOS

```bash
# Remove CLI
rm $(go env GOPATH)/bin/dmmvc

# Or use uninstall script
chmod +x scripts/uninstall.sh
./scripts/uninstall.sh

# Remove project directory
cd ..
rm -rf dmmvc
```

### Windows

```bash
# Remove CLI
del %GOPATH%\bin\dmmvc.exe

# Or
del %USERPROFILE%\go\bin\dmmvc.exe

# Remove project directory
cd ..
rmdir /s /q dmmvc
```

## Troubleshooting

### Go not found

```bash
# Install Go from official website
# https://golang.org/dl/

# Verify installation
go version
```

### GOPATH not set

```bash
# Check GOPATH
go env GOPATH

# If empty, Go uses default value
# Linux/macOS: ~/go
# Windows: %USERPROFILE%\go
```

### Error "command not found: dmmvc"

```bash
# Check if CLI is installed
ls $(go env GOPATH)/bin/dmmvc  # Linux/macOS
dir %GOPATH%\bin\dmmvc.exe     # Windows

# Check PATH
echo $PATH   # Linux/macOS
echo %PATH%  # Windows

# Add to PATH
export PATH=$PATH:$(go env GOPATH)/bin  # Linux/macOS
set PATH=%PATH%;%GOPATH%\bin            # Windows
```

### Build errors

```bash
# Clean Go cache
go clean -cache -modcache

# Reinstall dependencies
rm go.sum
go mod download
go mod tidy
```

### Permission denied (Linux/macOS)

```bash
# Make scripts executable
chmod +x scripts/*.sh

# Or for specific script
chmod +x scripts/install.sh
```

## Additional Tools

### Recommended VS Code Extensions

- Go (official extension)
- Go Template Support
- Docker
- GitLens

### Useful Commands

```bash
# Check code
go vet ./...

# Format code
go fmt ./...

# Run tests
go test ./...

# Show dependencies
go mod graph

# Clean unused dependencies
go mod tidy
```

## Next Steps

After installation:

1. Read [Quick Start](QUICKSTART.md)
2. Learn [CLI Tools](CLI.md)
3. Check [Examples](EXAMPLES.md)
4. Configure [Database](POSTGRESQL.md)

## Support

- **Documentation**: [docs/](.)
- **Issues**: https://github.com/dedomorozoff/dmmvc/issues
- **Discussions**: https://github.com/dedomorozoff/dmmvc/discussions
