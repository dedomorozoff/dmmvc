# 🧭 DMMVC Navigation Guide

Quick links to find what you need.

## 🚀 Getting Started

**New to DMMVC?** Start here:
1. [README.md](README.md) - Project overview
2. [docs/QUICKSTART.md](docs/QUICKSTART.md) - Quick start guide
3. [docs/QUICKSTART_CLI.md](docs/QUICKSTART_CLI.md) - CLI quick start

## 📚 Documentation

**All documentation**: [docs/](docs/)

### Core Guides
- [Architecture](docs/ARCHITECTURE.md) - How DMMVC works
- [CLI Tool](docs/CLI.md) - Code generation
- [Examples](docs/EXAMPLES.md) - Code examples
- [File Structure](docs/FILES.md) - Project organization

### Database
- [PostgreSQL Setup](docs/POSTGRESQL.md) - PostgreSQL configuration
- [Migration Guide](docs/MIGRATION_GUIDE.md) - Switch between databases

### Deployment
- [Deployment Guide](docs/DEPLOYMENT.md) - Production deployment
- [Docker Guide](docs/DOCKER.md) - Docker setup

## 🔧 Development

### Build & Run
```bash
make build    # Build CLI tool
make run      # Run server
make test     # Run tests
make clean    # Clean artifacts
```

### CLI Tool
```bash
dmmvc make:crud Product      # Create full CRUD
dmmvc make:controller About  # Create controller
dmmvc make:model Category    # Create model
dmmvc list                   # List resources
```

### Database Testing
```bash
go run scripts/test-db-connection.go
```

## 🐳 Docker

```bash
# Start with PostgreSQL
docker-compose -f docker/docker-compose.postgres.yml up -d

# View logs
docker-compose -f docker/docker-compose.postgres.yml logs -f

# Stop
docker-compose -f docker/docker-compose.postgres.yml down
```

See [docker/README.md](docker/README.md) for more.

## 📁 Project Structure

```
dmmvc/
├── cmd/          # Executables (cli, server)
├── internal/     # Application code
├── static/       # CSS, JS
├── templates/    # HTML templates
├── docs/         # 📚 Documentation
├── docker/       # 🐳 Docker files
└── scripts/      # 🔧 Utilities
```

Full structure: [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md)

## 🌐 Language

- 🇬🇧 English - Most files have `.md` extension
- 🇷🇺 Russian - Files with `.ru.md` extension

Example:
- `README.md` - English
- `README.ru.md` - Russian

## 🔍 Find Something?

### I want to...

**Learn the basics**
→ [docs/QUICKSTART.md](docs/QUICKSTART.md)

**Generate code**
→ [docs/CLI.md](docs/CLI.md)

**Setup PostgreSQL**
→ [docs/POSTGRESQL.md](docs/POSTGRESQL.md)

**Deploy to production**
→ [docs/DEPLOYMENT.md](docs/DEPLOYMENT.md)

**Use Docker**
→ [docs/DOCKER.md](docs/DOCKER.md)

**See examples**
→ [docs/EXAMPLES.md](docs/EXAMPLES.md)

**Understand architecture**
→ [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md)

**Migrate database**
→ [docs/MIGRATION_GUIDE.md](docs/MIGRATION_GUIDE.md)

## 📝 Configuration

- `.env.example` - Configuration template
- Copy to `.env` and customize
- See [docs/QUICKSTART.md](docs/QUICKSTART.md) for details

## 🆘 Need Help?

1. Check [docs/](docs/) for documentation
2. Read [CHANGELOG.md](CHANGELOG.md) for recent changes
3. Open an issue on GitHub

## 🎯 Quick Commands

```bash
# Development
go run cmd/server/main.go              # Run server
go run scripts/test-db-connection.go   # Test DB

# CLI
dmmvc make:crud Product                # Full CRUD
dmmvc list                             # List resources

# Docker
docker-compose -f docker/docker-compose.postgres.yml up -d

# Build
make build                             # Build CLI
make clean                             # Clean up
```

---

**Happy coding!** 🚀
