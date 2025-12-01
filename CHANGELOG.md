# Changelog

All notable changes to DMMVC will be documented in this file.

## [1.2.0] - 2024-12-02

### Added
- ✨ **PostgreSQL Support** - Full support for PostgreSQL database
  - PostgreSQL driver integration (gorm.io/driver/postgres)
  - Connection configuration with DSN and URL formats
  - SSL mode support (disable, require, verify-ca, verify-full)
  - Connection pool configuration
  - PostgreSQL-specific features (JSONB, arrays, full-text search)
- 📖 PostgreSQL Documentation (POSTGRESQL.md, POSTGRESQL.ru.md)
  - Installation guides for Windows, Linux, macOS
  - Configuration examples
  - Docker setup with docker-compose
  - Migration guides from SQLite/MySQL
  - Performance tips and troubleshooting
- Database type detection and logging
- Database connection testing function

### Changed
- Updated database connection logic to support postgres/postgresql
- Enhanced .env.example with PostgreSQL configuration examples
- Updated README.md and README.ru.md with PostgreSQL information
- Marked "PostgreSQL support" as completed in roadmap
- **Reorganized project structure**:
  - Moved all documentation to `docs/` folder
  - Moved Docker files to `docker/` folder
  - Cleaned up root directory for better organization

## [1.1.0] - 2024-12-02

### Added
- ✨ **CLI Tool** - Complete command-line tool for code generation
  - `make:controller` - Create controllers (simple or resource with CRUD)
  - `make:model` - Create models with optional migration hints
  - `make:middleware` - Create middleware
  - `make:page` - Create page templates
  - `make:crud` - Generate complete CRUD scaffolding (model + controller + pages)
  - `list` - List all project resources
- 📖 CLI Documentation (CLI.md, CLI.ru.md)
- 🔨 Makefile for easy building and installation
- Resource controller templates with full CRUD operations
- CRUD page templates (index, show, create, edit)

### Changed
- Updated README.md and README.ru.md with CLI information
- Marked "CLI tool for code generation" as completed in roadmap

## [1.0.0] - 2024-11-XX

### Added
- Initial release of DMMVC framework
- MVC architecture pattern
- Authentication system with bcrypt
- Session management
- GORM ORM integration
- SQLite and MySQL support
- Go Templates with layouts and partials
- Logrus structured logging
- Middleware system
- Static file serving
- User management
- Dashboard
- Example controllers and models
