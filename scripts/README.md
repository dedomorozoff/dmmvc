# DMMVC Scripts

Utility scripts for DMMVC framework installation and development.

## Overview

| Script | Linux/macOS | Windows | Purpose |
|--------|-------------|---------|---------|
| Installation | `install.sh` | `install.bat` | Install DMMVC CLI globally |
| Create Project | `create-project.sh` | `create-project.bat` | Create new DMMVC project |
| Quick Start | `quickstart.sh` | `quickstart.bat` | Setup and run in one command |
| Development | `dev.sh` | - | Development helper tasks |
| Testing | `test.sh` | `test.bat` | Run tests with options |
| Docker Build | `docker-build.sh` | `docker-build.bat` | Build Docker image |
| Uninstall | `uninstall.sh` | - | Remove DMMVC CLI |

## Installation Scripts

### install.sh (Linux/macOS)

Automatic installation script for Unix-like systems.

```bash
chmod +x scripts/install.sh
./scripts/install.sh
```

**What it does:**
- Checks Go installation
- Downloads dependencies
- Builds CLI tool
- Installs CLI globally to `$GOPATH/bin`

### install.bat (Windows)

Automatic installation script for Windows.

```bash
scripts\install.bat
```

**What it does:**
- Checks Go installation
- Downloads dependencies
- Builds CLI tool
- Installs CLI globally to `%GOPATH%\bin`

## Project Creation Scripts

### create-project.sh (Linux/macOS)

Creates a new DMMVC-based project.

```bash
chmod +x scripts/create-project.sh
./scripts/create-project.sh my-app
```

**What it does:**
- Creates project directory structure
- Initializes Go module
- Adds DMMVC as dependency
- Creates initial files (.env, main.go, README.md)
- Sets up .gitignore

### create-project.bat (Windows)

Creates a new DMMVC-based project for Windows.

```bash
scripts\create-project.bat my-app
```

Same functionality as the shell version.

## Quick Start

### quickstart.sh / quickstart.bat

One-command setup and run script.

**Linux/macOS:**
```bash
chmod +x scripts/quickstart.sh
./scripts/quickstart.sh
```

**Windows:**
```bash
scripts\quickstart.bat
```

**What it does:**
- Creates .env if missing
- Installs dependencies
- Builds CLI tool
- Starts development server

## Development Helper

### dev.sh (Linux/macOS)

Development helper script with common tasks.

```bash
chmod +x scripts/dev.sh
./scripts/dev.sh [command]
```

**Available commands:**

- `install` - Install dependencies
- `build` - Build CLI and server
- `cli` - Build CLI only
- `server` - Build server only
- `run` - Run development server
- `test` - Run tests
- `fmt` - Format code
- `lint` - Lint code (requires golangci-lint)
- `swagger` - Generate Swagger docs (requires swag)
- `clean` - Clean build artifacts
- `help` - Show help

**Examples:**

```bash
# Install dependencies
./scripts/dev.sh install

# Run development server
./scripts/dev.sh run

# Run tests
./scripts/dev.sh test

# Format and lint
./scripts/dev.sh fmt
./scripts/dev.sh lint

# Build everything
./scripts/dev.sh build
```

## Testing

### test.sh / test.bat

Test runner with coverage and benchmark support.

**Linux/macOS:**
```bash
chmod +x scripts/test.sh
./scripts/test.sh [options]
```

**Windows:**
```bash
scripts\test.bat [options]
```

**Options:**
- `-c, --coverage` - Run with coverage report
- `-v, --verbose` - Verbose output
- `-b, --bench` - Run benchmarks
- `-h, --help` - Show help

**Examples:**
```bash
# Run tests
./scripts/test.sh

# Run with coverage
./scripts/test.sh --coverage

# Run benchmarks
./scripts/test.sh --bench

# Verbose mode
./scripts/test.sh -v -c
```

## Docker

### docker-build.sh / docker-build.bat

Build Docker image for the application.

**Linux/macOS:**
```bash
chmod +x scripts/docker-build.sh
./scripts/docker-build.sh [image-name] [tag]
```

**Windows:**
```bash
scripts\docker-build.bat [image-name] [tag]
```

**Examples:**
```bash
# Build with default name (dmmvc:latest)
./scripts/docker-build.sh

# Build with custom name
./scripts/docker-build.sh myapp v1.0

# Run built image
docker run -p 8080:8080 dmmvc:latest
```

## Uninstallation

### uninstall.sh (Linux/macOS)

Removes DMMVC CLI from the system.

```bash
chmod +x scripts/uninstall.sh
./scripts/uninstall.sh
```

**What it does:**
- Removes CLI binary from `$GOPATH/bin`
- Shows instructions for removing project directory

## Making Scripts Executable

On Linux/macOS, you need to make scripts executable before first use:

```bash
# Make all scripts executable
chmod +x scripts/*.sh

# Or individually
chmod +x scripts/install.sh
chmod +x scripts/create-project.sh
chmod +x scripts/dev.sh
chmod +x scripts/uninstall.sh
```

## Troubleshooting

### Permission Denied (Linux/macOS)

If you get "Permission denied" error:

```bash
chmod +x scripts/script-name.sh
```

### Script Not Found

Make sure you're in the project root directory:

```bash
cd /path/to/dmmvc
./scripts/install.sh
```

### Go Not Found

Install Go from https://golang.org/dl/

### GOPATH Issues

Check your GOPATH:

```bash
# Linux/macOS
go env GOPATH

# Windows
go env GOPATH
```

Add Go bin to PATH:

```bash
# Linux/macOS (add to ~/.bashrc or ~/.zshrc)
export PATH=$PATH:$(go env GOPATH)/bin

# Windows (add to system PATH)
set PATH=%PATH%;%GOPATH%\bin
```

## Platform-Specific Notes

### Linux/macOS

- Scripts use bash shell
- Require execute permissions (`chmod +x`)
- Use forward slashes for paths
- Support colored output

### Windows

- Scripts use batch/cmd
- No execute permissions needed
- Use backslashes for paths
- Limited color support

## Contributing

When adding new scripts:

1. Create both `.sh` and `.bat` versions
2. Add documentation to this README
3. Test on target platforms
4. Follow existing script patterns
5. Add error handling
6. Include helpful messages

## License

MIT License - see main project LICENSE file
