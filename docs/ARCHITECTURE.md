**English** | [Русский](ARCHITECTURE.ru.md)

# DMMVC Framework Architecture

## Architecture Overview

DMMVC follows the classic MVC (Model-View-Controller) pattern with additional layers for middleware and services.

```
┌─────────────────────────────────────────────────────────────┐
│                         Browser                              │
└────────────────────────┬────────────────────────────────────┘
                         │ HTTP Request
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                      Gin Router                              │
│                    (routes/routes.go)                        │
└────────────────────────┬────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                     Middleware Layer                         │
│  ┌──────────────┬──────────────┬──────────────────────┐    │
│  │   Logger     │     Auth     │   Custom Middleware  │    │
│  └──────────────┴──────────────┴──────────────────────┘    │
└────────────────────────┬────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                     Controllers                              │
│  ┌──────────────┬──────────────┬──────────────────────┐    │
│  │     Auth     │     Home     │        User          │    │
│  └──────────────┴──────────────┴──────────────────────┘    │
└────────────────────────┬────────────────────────────────────┘
                         │
                         ▼
         ┌───────────────┴───────────────┐
         ▼                               ▼
┌─────────────────────┐         ┌─────────────────────┐
│      Models         │         │       Views         │
│  ┌──────────────┐   │         │  ┌──────────────┐   │
│  │     User     │   │         │  │   Templates  │   │
│  └──────────────┘   │         │  └──────────────┘   │
│         │           │         └─────────────────────┘
│         ▼           │
│  ┌──────────────┐   │
│  │   Database   │   │
│  │   (GORM)     │   │
│  └──────────────┘   │
└─────────────────────┘
```

## Directory Structure

```
dmmvc/
│
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
│
├── internal/                       # Internal application code
│   │
│   ├── controllers/                # HTTP handlers (Controller)
│   │   ├── auth_controller.go      # Authentication
│   │   ├── home_controller.go      # Home page
│   │   └── user_controller.go      # User CRUD
│   │
│   ├── models/                     # Data models (Model)
│   │   └── user.go                 # User model
│   │
│   ├── database/                   # Database operations
│   │   ├── db.go                   # Database connection
│   │   └── seeder.go               # Initial data
│   │
│   ├── middleware/                 # Middleware layer
│   │   ├── auth.go                 # Authorization check
│   │   └── logger.go               # Request logging
│   │
│   ├── routes/                     # Routing
│   │   └── routes.go               # Route definitions
│   │
│   └── logger/                     # Logging
│       └── logger.go               # Logger configuration
│
├── templates/                      # HTML templates (View)
│   ├── layouts/
│   │   └── base.html               # Base layout
│   ├── partials/
│   │   ├── header.html             # Site header
│   │   └── footer.html             # Site footer
│   └── pages/
│       ├── home.html               # Home page
│       ├── login.html              # Login page
│       ├── dashboard.html          # Dashboard
│       └── users/
│           ├── list.html           # User list
│           ├── create.html         # Create user
│           └── edit.html           # Edit user
│
├── static/                         # Static files
│   ├── css/
│   │   └── style.css               # Styles
│   └── js/
│       └── app.js                  # JavaScript
│
├── .env                            # Configuration
├── .env.example                    # Configuration example
├── go.mod                          # Go modules
├── README.md                       # Documentation
└── QUICKSTART.md                   # Quick start
```

## Data Flow

### 1. Public Request (e.g., Home Page)

```
Browser → Router → Middleware (Logger) → Controller (HomePage) → View (home.html) → Browser
```

### 2. Protected Request (e.g., Dashboard)

```
Browser → Router → Middleware (Logger, Auth) → Controller (DashboardPage) → View (dashboard.html) → Browser
```

### 3. CRUD Operation (e.g., Create User)

```
Browser → Router → Middleware (Logger, Auth) → Controller (UserStore) → Model (User) → Database → Redirect → Browser
```

## Components

### Router (routes/routes.go)
- Defines all application routes
- Groups routes (public, protected, admin)
- Applies middleware to route groups

### Middleware
- **Logger**: Logs all HTTP requests
- **Auth**: Checks user authorization
- **Custom**: You can add your own middleware

### Controllers
- Handle HTTP requests
- Interact with models
- Return HTML or JSON responses

### Models
- Define data structure
- Use GORM for database interaction
- Contain business logic

### Views (Templates)
- **Layouts**: Base page structure
- **Partials**: Reusable components
- **Pages**: Specific pages

### Database
- Connection to SQLite/MySQL
- Automatic migrations
- Initial data seeding

## Development Principles

### 1. Separation of Concerns
Each component is responsible for its own task:
- Controllers - request handling
- Models - data operations
- Views - display

### 2. DRY (Don't Repeat Yourself)
- Reuse layouts and partials
- Common middleware for all routes
- Base models with GORM

### 3. Convention over Configuration
- Standard directory structure
- Naming by convention
- Minimal configuration

### 4. Security First
- Password hashing (bcrypt)
- Session protection
- Authorization middleware

## Extending the Framework

### Adding New Functionality

1. **Create Model** in `internal/models/`
2. **Create Controller** in `internal/controllers/`
3. **Add Routes** in `internal/routes/routes.go`
4. **Create Templates** in `templates/pages/`
5. **Add Styles** in `static/css/style.css`

### Adding Middleware

1. Create file in `internal/middleware/`
2. Implement `gin.HandlerFunc`
3. Apply in `routes.go`

### Adding Service

1. Create directory `internal/services/`
2. Implement business logic
3. Use in controllers

---

This architecture ensures:
- Scalability
- Maintainability
- Testability
- Security
- Performance
