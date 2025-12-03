# Updating Dependencies

This guide explains how to keep your DMMVC framework dependencies up to date.

## Quick Update

To update all dependencies to their latest versions:

```bash
# Update all dependencies
go get -u ./...

# Clean up and download missing dependencies
go mod tidy

# Verify everything compiles
go build -o dmmvc cmd/server/main.go
```

## Checking for Updates

### Check outdated packages

```bash
# List all dependencies
go list -m all

# Check for available updates
go list -u -m all
```

### Check specific package

```bash
go list -m -u github.com/gin-gonic/gin
```

## Updating Specific Packages

### Update a single package

```bash
go get -u github.com/gin-gonic/gin@latest
```

### Update to a specific version

```bash
go get github.com/gin-gonic/gin@v1.9.0
```

### Update major version

```bash
go get github.com/gin-gonic/gin@v2
```

## Testing After Updates

After updating dependencies, always test your application:

```bash
# Run tests
go test ./...

# Build the application
go build -o dmmvc cmd/server/main.go

# Run the server
./dmmvc
# or
go run cmd/server/main.go
```

## Common Issues

### Breaking Changes

Major version updates may include breaking changes. Always check:
- Package changelog/release notes
- Migration guides
- API documentation

### Dependency Conflicts

If you encounter conflicts:

```bash
# Remove go.sum and try again
rm go.sum
go mod tidy

# Or reset to a working state
git checkout go.mod go.sum
```

### Vendor Directory

If using vendoring:

```bash
# Update vendor directory
go mod vendor

# Verify vendor directory
go mod verify
```

## Recommended Update Schedule

- **Security updates**: Immediately
- **Minor updates**: Monthly
- **Major updates**: Quarterly (with testing)

## Key Dependencies

### Core Framework
- `github.com/gin-gonic/gin` - Web framework
- `gorm.io/gorm` - ORM
- `github.com/gin-contrib/sessions` - Session management

### Database Drivers
- `github.com/glebarez/sqlite` - SQLite (pure Go)
- `gorm.io/driver/mysql` - MySQL
- `gorm.io/driver/postgres` - PostgreSQL

### Optional Features
- `github.com/gorilla/websocket` - WebSocket support
- `github.com/redis/go-redis/v9` - Redis client
- `github.com/swaggo/gin-swagger` - Swagger documentation
- `github.com/hibiken/asynq` - Task queue
- `github.com/nfnt/resize` - Image processing

### Utilities
- `github.com/joho/godotenv` - Environment variables
- `github.com/sirupsen/logrus` - Logging
- `golang.org/x/crypto` - Cryptography

## Automated Updates

### Using Dependabot (GitHub)

Create `.github/dependabot.yml`:

```yaml
version: 2
updates:
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"
    open-pull-requests-limit: 10
```

### Using Renovate

Create `renovate.json`:

```json
{
  "extends": ["config:base"],
  "packageRules": [
    {
      "matchUpdateTypes": ["minor", "patch"],
      "automerge": true
    }
  ]
}
```

## Rollback

If an update causes issues:

```bash
# Revert to previous versions
git checkout go.mod go.sum

# Or manually edit go.mod and run
go mod tidy
```

## Best Practices

1. **Read changelogs** before updating
2. **Test thoroughly** after updates
3. **Update regularly** to avoid large jumps
4. **Keep go.sum** in version control
5. **Document breaking changes** in your CHANGELOG
6. **Use semantic versioning** for your own packages

## Resources

- [Go Modules Reference](https://go.dev/ref/mod)
- [Go Module Versioning](https://go.dev/doc/modules/version-numbers)
- [Semantic Versioning](https://semver.org/)
