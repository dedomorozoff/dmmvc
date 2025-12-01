# DMMVC Project Structure

## Directory Layout

```
dmmvc/
├── cmd/                     # Application entry points
│   ├── cli/                 # CLI tool for code generation
│   └── server/              # Web server
│
├── internal/                # Private application code
│   ├── controllers/         # HTTP request handlers
│   ├── database/            # Database connection and migrations
│   ├── logger/              # Logging configuration
│   ├── middleware/          # HTTP middleware
│   ├── models/              # Data models (GORM)
│   └── routes/              # Route definitions
│
├── static/                  # Static assets
│   ├── css/                 # Stylesheets
│   └── js/                  # JavaScript files
│
├── templates/               # HTML templates
│   ├── layouts/             # Base layouts
│   ├── partials/            # Reusable components
│   └── pages/               # Page templates
│
├── docs/                    # Documentation
│   ├── README.md            # Documentation index
│   ├── QUICKSTART.md        # Quick start guide
│   ├── CLI.md               # CLI tool documentation
│   ├── POSTGRESQL.md        # PostgreSQL setup
│   ├── DOCKER.md            # Docker deployment
│   └── ...                  # Other documentation
│
├── docker/                  # Docker configuration
│   ├── docker-compose.postgres.yml
│   ├── init-db.sql
│   └── README.md
│
├── scripts/                 # Utility scripts
│   ├── test-db-connection.go
│   └── README.md
│
├── .env.example             # Environment configuration template
├── .gitignore               # Git ignore rules
├── CHANGELOG.md             # Version history
├── Dockerfile               # Docker image definition
├── go.mod                   # Go module definition
├── LICENSE                  # MIT License
├── Makefile                 # Build automation
├── README.md                # Main documentation (English)
└── README.ru.md             # Main documentation (Russian)
```

## Key Directories

### `/cmd`
Application entry points. Each subdirectory is a separate executable:
- `cli/` - Code generation tool
- `server/` - Web application server

### `/internal`
Private application code that cannot be imported by other projects:
- `controllers/` - Handle HTTP requests
- `database/` - Database setup and seeding
- `logger/` - Centralized logging
- `middleware/` - Request/response processing
- `models/` - Database models
- `routes/` - URL routing

### `/static`
Public assets served directly:
- `css/` - Stylesheets
- `js/` - Client-side JavaScript

### `/templates`
Go HTML templates:
- `layouts/` - Base page structures
- `partials/` - Reusable components (header, footer, etc.)
- `pages/` - Individual page templates

### `/docs`
All project documentation:
- Getting started guides
- API documentation
- Deployment guides
- Architecture documentation

### `/docker`
Docker-related files:
- Docker Compose configurations
- Database initialization scripts
- Container setup documentation

### `/scripts`
Utility scripts for development and testing

## Configuration Files

- `.env` - Environment variables (not in git)
- `.env.example` - Template for environment configuration
- `go.mod` / `go.sum` - Go dependencies
- `Makefile` - Build and development commands
- `Dockerfile` - Container image definition
- `.gitignore` - Files to exclude from git
- `.dockerignore` - Files to exclude from Docker build

## Generated Files

These files are generated and should not be committed:
- `*.exe` - Compiled binaries
- `*.db` - SQLite database files
- `*.log` - Log files
- `tmp/` - Temporary files

## Documentation

For detailed documentation, see:
- [docs/README.md](docs/README.md) - Documentation index
- [README.md](README.md) - Main project README

## Development

```bash
# Build CLI tool
make build

# Run server
make run

# Run tests
make test

# Clean build artifacts
make clean
```

## Docker

```bash
# Start with Docker Compose
docker-compose -f docker/docker-compose.postgres.yml up -d

# Build Docker image
docker build -t dmmvc:latest .
```

## Learn More

- [Quick Start Guide](docs/QUICKSTART.md)
- [CLI Documentation](docs/CLI.md)
- [Architecture Guide](docs/ARCHITECTURE.md)
- [Deployment Guide](docs/DEPLOYMENT.md)
