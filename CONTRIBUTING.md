# Contributing to DMMVC

Thank you for your interest in contributing to DMMVC! This document provides guidelines and instructions for contributing.

## Getting Started

### Prerequisites

- Go 1.20 or higher
- Git
- Basic knowledge of Go and MVC pattern

### Setup Development Environment

```bash
# Fork and clone the repository
git clone https://github.com/YOUR_USERNAME/dmmvc
cd dmmvc

# Install dependencies
go mod download

# Run quick start
./scripts/quickstart.sh      # Linux/macOS
scripts\quickstart.bat       # Windows

# Verify installation
go test ./...
```

## Development Workflow

### 1. Create a Branch

```bash
# Create feature branch
git checkout -b feature/your-feature-name

# Or bug fix branch
git checkout -b fix/bug-description
```

### 2. Make Changes

- Write clean, readable code
- Follow Go conventions
- Add tests for new features
- Update documentation

### 3. Test Your Changes

```bash
# Run tests
go test ./...

# Run with coverage
./scripts/test.sh --coverage

# Format code
go fmt ./...

# Vet code
go vet ./...

# Lint (if golangci-lint installed)
./scripts/dev.sh lint
```

### 4. Commit Changes

```bash
# Stage changes
git add .

# Commit with descriptive message
git commit -m "feat: add new feature"
git commit -m "fix: resolve bug in controller"
git commit -m "docs: update installation guide"
```

#### Commit Message Format

Use conventional commits:

- `feat:` - New feature
- `fix:` - Bug fix
- `docs:` - Documentation changes
- `style:` - Code style changes (formatting, etc.)
- `refactor:` - Code refactoring
- `test:` - Adding or updating tests
- `chore:` - Maintenance tasks

### 5. Push and Create Pull Request

```bash
# Push to your fork
git push origin feature/your-feature-name

# Create Pull Request on GitHub
# Provide clear description of changes
```

## Code Guidelines

### Go Code Style

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting
- Use meaningful variable names
- Add comments for exported functions
- Keep functions small and focused

### Example

```go
// GetUserByID retrieves a user from the database by ID
func GetUserByID(id uint) (*models.User, error) {
    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
```

### Project Structure

```
dmmvc/
├── cmd/                    # Entry points
│   ├── cli/               # CLI tool
│   └── server/            # Web server
├── internal/              # Internal packages
│   ├── controllers/       # HTTP handlers
│   ├── models/            # Data models
│   ├── middleware/        # Middleware
│   └── ...
├── scripts/               # Utility scripts
├── docs/                  # Documentation
└── templates/             # HTML templates
```

## Adding New Features

### Adding a Controller

```go
// internal/controllers/example.go
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func ExampleController(c *gin.Context) {
    c.HTML(http.StatusOK, "pages/example.html", gin.H{
        "title": "Example",
    })
}
```

### Adding a Model

```go
// internal/models/example.go
package models

import "gorm.io/gorm"

type Example struct {
    gorm.Model
    Name        string `json:"name" gorm:"not null"`
    Description string `json:"description"`
}
```

### Adding a Route

```go
// internal/routes/routes.go
authorized.GET("/example", controllers.ExampleController)
```

### Adding a Template

```html
<!-- templates/pages/example.html -->
{{define "pages/example.html"}}
{{template "layouts/base.html" .}}
{{define "content"}}
<div class="container">
    <h1>{{.title}}</h1>
</div>
{{end}}
{{end}}
```

## Adding Scripts

When adding new utility scripts:

### 1. Create Both Versions

- `scripts/script-name.sh` for Linux/macOS
- `scripts/script-name.bat` for Windows

### 2. Follow Existing Patterns

**Shell Script Template:**

```bash
#!/bin/bash
# Script Description

set -e

echo "========================================"
echo "Script Name"
echo "========================================"
echo ""

# Your code here

echo ""
echo "✓ Complete!"
```

**Batch Script Template:**

```batch
@echo off
REM Script Description

echo ========================================
echo Script Name
echo ========================================
echo.

REM Your code here

echo.
echo [OK] Complete!
```

### 3. Make Executable (Linux/macOS)

```bash
chmod +x scripts/script-name.sh
```

### 4. Update Documentation

- Add to `scripts/README.md`
- Update main README if needed
- Add usage examples

## Documentation

### Adding Documentation

1. Create markdown file in `docs/`
2. Use clear headings and examples
3. Add to main README index
4. Create Russian version if possible

### Documentation Structure

```markdown
# Feature Name

Brief description

## Installation

Installation steps

## Usage

Usage examples

## Configuration

Configuration options

## Examples

Code examples

## Troubleshooting

Common issues and solutions
```

## Testing

### Writing Tests

```go
// internal/controllers/example_test.go
package controllers

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestExampleController(t *testing.T) {
    // Setup
    // ...
    
    // Test
    // ...
    
    // Assert
    assert.Equal(t, expected, actual)
}
```

### Running Tests

```bash
# All tests
go test ./...

# Specific package
go test ./internal/controllers

# With coverage
go test -cover ./...

# Verbose
go test -v ./...
```

## Pull Request Process

1. **Update Documentation** - Update relevant docs
2. **Add Tests** - Include tests for new features
3. **Pass CI** - Ensure all checks pass
4. **Code Review** - Address review comments
5. **Squash Commits** - Clean up commit history if needed

### PR Checklist

- [ ] Code follows project style
- [ ] Tests added/updated
- [ ] Documentation updated
- [ ] All tests pass
- [ ] No merge conflicts
- [ ] Descriptive PR title and description

## Reporting Issues

### Bug Reports

Include:

- DMMVC version
- Go version
- Operating system
- Steps to reproduce
- Expected behavior
- Actual behavior
- Error messages/logs

### Feature Requests

Include:

- Use case description
- Proposed solution
- Alternative solutions considered
- Additional context

## Community

### Communication Channels

- **GitHub Issues** - Bug reports and feature requests
- **GitHub Discussions** - Questions and discussions
- **Pull Requests** - Code contributions

### Code of Conduct

- Be respectful and inclusive
- Welcome newcomers
- Focus on constructive feedback
- Help others learn

## Recognition

Contributors will be:

- Listed in CONTRIBUTORS.md
- Mentioned in release notes
- Credited in documentation

## Questions?

If you have questions:

1. Check existing documentation
2. Search closed issues
3. Ask in GitHub Discussions
4. Open a new issue

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to DMMVC! 🎉
