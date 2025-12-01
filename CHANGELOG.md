# Changelog

All notable changes to DMMVC will be documented in this file.

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
