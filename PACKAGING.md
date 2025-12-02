# DMMVC Packaging Guide

This document describes how DMMVC is packaged for distribution and use.

## Package Structure

DMMVC can be used in multiple ways:

### 1. As a Go Module (Library)

```bash
go get github.com/dedomorozoff/dmmvc@latest
```

Import and use in your project:

```go
import (
    "github.com/dedomorozoff/dmmvc/internal/database"
    "github.com/dedomorozoff/dmmvc/internal/logger"
)
```

### 2. As a CLI Tool

Install globally:

```bash
# Via go install
go install github.com/dedomorozoff/dmmvc/cmd/cli@latest

# Via installation script
./scripts/install.sh      # Linux/macOS
scripts\install.bat       # Windows
```

Use CLI commands:

```bash
dmmvc make:crud Product
dmmvc make:controller About
dmmvc make:model Category
```

### 3. As a Project Template

Create new projects based on DMMVC:

```bash
./scripts/create-project.sh my-app      # Linux/macOS
scripts\create-project.bat my-app       # Windows
```

### 4. As a Docker Image

Build and run in containers:

```bash
docker build -t dmmvc .
docker run -p 8080:8080 dmmvc
```

## Installation Scripts

### Cross-Platform Scripts

| Purpose | Linux/macOS | Windows |
|---------|-------------|---------|
| Install CLI | `install.sh` | `install.bat` |
| Create Project | `create-project.sh` | `create-project.bat` |
| Quick Start | `quickstart.sh` | `quickstart.bat` |
| Test Runner | `test.sh` | `test.bat` |
| Docker Build | `docker-build.sh` | `docker-build.bat` |

### Linux/macOS Only

- `dev.sh` - Development helper with multiple commands
- `uninstall.sh` - Remove CLI from system

## Distribution Methods

### 1. Source Distribution

Users clone the repository and build from source:

```bash
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc
./scripts/install.sh
```

**Advantages:**
- Always latest code
- Can modify and customize
- Full source access

**Use cases:**
- Development
- Contributing
- Learning

### 2. Binary Distribution

Pre-built binaries for different platforms:

```bash
# Download binary
wget https://github.com/dedomorozoff/dmmvc/releases/latest/download/dmmvc-linux-amd64
chmod +x dmmvc-linux-amd64
sudo mv dmmvc-linux-amd64 /usr/local/bin/dmmvc
```

**Advantages:**
- No build required
- Fast installation
- Smaller download

**Use cases:**
- Production deployment
- CI/CD pipelines
- Quick installation

### 3. Package Managers

Future support for package managers:

```bash
# Homebrew (macOS/Linux)
brew install dmmvc

# Chocolatey (Windows)
choco install dmmvc

# Snap (Linux)
snap install dmmvc

# APT (Debian/Ubuntu)
apt install dmmvc

# YUM (RedHat/CentOS)
yum install dmmvc
```

### 4. Docker Hub

Pull pre-built images:

```bash
docker pull dedomorozoff/dmmvc:latest
docker run -p 8080:8080 dedomorozoff/dmmvc
```

## Build Process

### CLI Tool

```bash
# Build for current platform
go build -o dmmvc cmd/cli/main.go

# Build for specific platform
GOOS=linux GOARCH=amd64 go build -o dmmvc-linux-amd64 cmd/cli/main.go
GOOS=windows GOARCH=amd64 go build -o dmmvc-windows-amd64.exe cmd/cli/main.go
GOOS=darwin GOARCH=amd64 go build -o dmmvc-darwin-amd64 cmd/cli/main.go
```

### Server

```bash
# Build server
go build -o server cmd/server/main.go

# Build with optimizations
go build -ldflags="-s -w" -o server cmd/server/main.go
```

### Docker Image

```bash
# Build image
docker build -t dmmvc:latest .

# Multi-platform build
docker buildx build --platform linux/amd64,linux/arm64 -t dmmvc:latest .
```

## Release Process

### 1. Version Tagging

```bash
# Create version tag
git tag -a v1.0.0 -m "Release version 1.0.0"
git push origin v1.0.0
```

### 2. Build Binaries

```bash
# Build for all platforms
./scripts/build-release.sh
```

### 3. Create Release

- Create GitHub release
- Upload binaries
- Write changelog
- Update documentation

### 4. Publish Docker Image

```bash
# Tag and push
docker tag dmmvc:latest dedomorozoff/dmmvc:1.0.0
docker tag dmmvc:latest dedomorozoff/dmmvc:latest
docker push dedomorozoff/dmmvc:1.0.0
docker push dedomorozoff/dmmvc:latest
```

## Installation Verification

After installation, verify:

```bash
# Check CLI
dmmvc --version
dmmvc --help

# Check Go module
go list -m github.com/dedomorozoff/dmmvc

# Check Docker image
docker images | grep dmmvc
```

## Uninstallation

### Remove CLI

```bash
# Linux/macOS
rm $(go env GOPATH)/bin/dmmvc
./scripts/uninstall.sh

# Windows
del %GOPATH%\bin\dmmvc.exe
```

### Remove Project

```bash
# Linux/macOS
rm -rf /path/to/dmmvc

# Windows
rmdir /s /q C:\path\to\dmmvc
```

### Remove Docker Image

```bash
docker rmi dmmvc:latest
```

## Documentation

All packaging-related documentation:

- [Installation Guide](docs/INSTALLATION.md) - Detailed installation instructions
- [Installation Guide (RU)](docs/INSTALLATION.ru.md) - Russian version
- [Scripts README](scripts/README.md) - Script documentation
- [Cheat Sheet](CHEATSHEET.md) - Quick reference
- [Cheat Sheet (RU)](CHEATSHEET.ru.md) - Russian quick reference

## Files Created

### Scripts (Cross-Platform)

- `scripts/install.sh` / `scripts/install.bat` - Installation
- `scripts/create-project.sh` / `scripts/create-project.bat` - Project creation
- `scripts/quickstart.sh` / `scripts/quickstart.bat` - Quick start
- `scripts/test.sh` / `scripts/test.bat` - Test runner
- `scripts/docker-build.sh` / `scripts/docker-build.bat` - Docker build

### Scripts (Linux/macOS Only)

- `scripts/dev.sh` - Development helper
- `scripts/uninstall.sh` - Uninstallation

### Documentation

- `docs/INSTALLATION.md` - Installation guide (English)
- `docs/INSTALLATION.ru.md` - Installation guide (Russian)
- `scripts/README.md` - Scripts documentation
- `CHEATSHEET.md` - Quick reference (English)
- `CHEATSHEET.ru.md` - Quick reference (Russian)
- `PACKAGING.md` - This file

### Configuration

- `Makefile` - Updated with new targets
- `README.md` - Updated with installation info
- `README.ru.md` - Updated with installation info (Russian)

## Best Practices

### For Users

1. Use installation scripts for easiest setup
2. Use `quickstart.sh` for immediate development
3. Use `dev.sh` for common development tasks
4. Keep CLI updated with `go install`

### For Contributors

1. Test scripts on all platforms
2. Update documentation when adding features
3. Follow existing script patterns
4. Add both `.sh` and `.bat` versions

### For Maintainers

1. Version all releases
2. Build binaries for all platforms
3. Update Docker images
4. Keep documentation current
5. Test installation methods

## Future Enhancements

### Planned Features

- [ ] Package manager support (Homebrew, Chocolatey, etc.)
- [ ] Auto-update mechanism
- [ ] Plugin system
- [ ] GUI installer
- [ ] Web-based project generator
- [ ] VS Code extension
- [ ] GitHub Actions templates

### Improvements

- [ ] Faster installation
- [ ] Smaller binary size
- [ ] Better error messages
- [ ] Progress indicators
- [ ] Rollback capability
- [ ] Configuration wizard

## Support

For packaging-related issues:

- **Documentation**: [docs/](docs/)
- **Issues**: https://github.com/dedomorozoff/dmmvc/issues
- **Discussions**: https://github.com/dedomorozoff/dmmvc/discussions

## License

MIT License - see [LICENSE](LICENSE) file
