# Changelog

All notable changes to DMMVC will be documented in this file.

## [1.5.0] - 2025-12-04

### Added
- **Feature Toggles System** - Modular feature management
  - Enable/disable WebSocket, Redis, Swagger, File Upload, i18n, Queue, and Email features
  - Environment variable configuration (WEBSOCKET_ENABLE, REDIS_ENABLE, etc.)
  - All features disabled by default - enable only what you need
  - Conditional initialization of features
  - Conditional route registration
  - Reduced dependencies for disabled features
  - Feature status logging on startup
  - Dashboard UI showing enabled/disabled features with instructions
  - Navigation menu items conditionally displayed based on enabled features
- New config package (internal/config/features.go)
- New middleware for injecting feature status (internal/middleware/features.go)
- Feature toggles documentation (docs/FEATURES.md, docs/FEATURES.ru.md)
- Quick start guides (docs/QUICKSTART_FEATURES.md, docs/QUICKSTART_FEATURES.ru.md)
- Feature cheatsheet (docs/FEATURES_CHEATSHEET.md)
- Feature configuration in .env.example and .env
- Dependency update documentation (docs/UPDATING.md, docs/UPDATING.ru.md)

### Changed
- Updated cmd/server/main.go to conditionally initialize features
- Updated internal/routes/routes.go to conditionally register routes
- Updated README.md and README.ru.md with feature toggles information
- All templates now use static English text (home, login, dashboard, profile, user management)
- Language switcher hidden when i18n is disabled
- Controllers updated to support both i18n-enabled and disabled modes
- User management pages (list, create, edit) converted to English
- **Updated all dependencies to latest versions**:
  - gorilla/websocket: v1.5.1 → v1.5.3
  - go-sql-driver/mysql: v1.8.1 → v1.9.3
  - jackc/pgx/v5: v5.6.0 → v5.7.6
  - glebarez/go-sqlite: v1.21.2 → v1.22.0
  - modernc.org/sqlite: v1.23.1 → v1.40.1
  - spf13/cast: v1.7.0 → v1.10.0
  - golang.org/x/time: v0.12.0 → v0.14.0
  - And other indirect dependencies

## [1.4.0] - 2025-12-02

### Added
- **Internationalization (i18n) Support** - Multi-language support system
  - Automatic locale detection from query params, cookies, and Accept-Language header
  - Translation management with JSON files
  - Template integration with `t` and `locale` functions
  - Middleware for locale detection and context injection
  - API endpoints for locale switching and listing available locales
  - Language switcher UI component with dropdown selector
  - English and Russian translations included
  - Thread-safe translation loading and caching
  - Fallback mechanism to default locale
  - Configurable default locale via DEFAULT_LOCALE environment variable
- i18n Documentation (I18N.md, I18N.ru.md)
  - Configuration and usage guide
  - Template and handler integration examples
  - API endpoint documentation
  - Guide for adding new languages
  - Best practices and troubleshooting
- Translation files (locales/en.json, locales/ru.json)
- Language switcher JavaScript component (static/js/i18n.js)
- Language switcher CSS styles (static/css/i18n.css)

### Changed
- Updated routes to include i18n middleware and API endpoints
- Updated routes to use filepath.Glob for better Windows compatibility with nested templates
- Updated main.go to initialize i18n on startup
- Updated base templates to include i18n scripts and styles
- Updated README.md and README.ru.md with i18n information
- Updated .env.example with DEFAULT_LOCALE configuration
- Marked "Localization (i18n)" as completed in roadmap

### Fixed
- Fixed template loading on Windows for nested directories (templates/pages/users/)

## [1.3.0] - 2025-12-02

### Added
- **WebSocket Support** - Real-time bidirectional communication
  - Hub system for managing multiple connections
  - Client connection management with auto-reconnection
  - Broadcast messaging to all connected clients
  - Ping/Pong health checks
  - Read/Write pumps for message handling
  - WebSocket demo page with interactive chat interface
- WebSocket Documentation (WEBSOCKET.md, WEBSOCKET.ru.md)
  - Architecture overview (Hub and Client)
  - Usage examples (chat, notifications, live updates)
  - Client-side JavaScript integration
  - Security and authentication guidelines
  - Performance optimization tips
  - Testing and troubleshooting guides
- gorilla/websocket dependency

### Changed
- Updated routes to include WebSocket endpoint (/ws)
- Added WebSocket demo page to authorized routes
- Updated README.md and README.ru.md with WebSocket information
- Marked "WebSocket support" as completed in roadmap

## [1.2.0] - 2025-12-02

### Added
- **PostgreSQL Support** - Full support for PostgreSQL database
  - PostgreSQL driver integration (gorm.io/driver/postgres)
  - Connection configuration with DSN and URL formats
  - SSL mode support (disable, require, verify-ca, verify-full)
  - Connection pool configuration
  - PostgreSQL-specific features (JSONB, arrays, full-text search)
- PostgreSQL Documentation (POSTGRESQL.md, POSTGRESQL.ru.md)
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

## [1.1.0] - 2025-12-02

### Added
- **CLI Tool** - Complete command-line tool for code generation
  - `make:controller` - Create controllers (simple or resource with CRUD)
  - `make:model` - Create models with optional migration hints
  - `make:middleware` - Create middleware
  - `make:page` - Create page templates
  - `make:crud` - Generate complete CRUD scaffolding (model + controller + pages)
  - `list` - List all project resources
- CLI Documentation (CLI.md, CLI.ru.md)
- Makefile for easy building and installation
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
