[English](CLI.md) | [Русский](CLI.ru.md) | [Docs](README.md)

# DMMVC CLI Tool

Command-line tool for code generation in DMMVC framework.

## Installation

### Build CLI

```bash
make build
```

This will create an executable file `dmmvc.exe` in the project root.

### Global Installation

```bash
make install
```

After this, the `dmmvc` command will be available globally.

## Usage

### Main Commands

#### Create Controller

```bash
# Simple controller
dmmvc make:controller Product

# Resource controller with CRUD methods
dmmvc make:controller Product --resource
```

#### Create Model

```bash
# Simple model
dmmvc make:model Product

# Model with migration hint
dmmvc make:model Product --migration
```

#### Create Middleware

```bash
dmmvc make:middleware RateLimit
```

#### Create Page

```bash
dmmvc make:page about
```

#### Create Full CRUD

```bash
# Creates model, controller, and all pages for CRUD
dmmvc make:crud Product
```

This command will create:
- Model: `internal/models/product.go`
- Resource controller: `internal/controllers/product_controller.go`
- Pages:
  - `templates/pages/product/index.html` - list
  - `templates/pages/product/show.html` - view
  - `templates/pages/product/create.html` - create
  - `templates/pages/product/edit.html` - edit

#### List Resources

```bash
# Show all controllers, models, middleware, and pages
dmmvc list
```

### Help

```bash
# Show help
dmmvc --help
dmmvc -h

# Show version
dmmvc --version
dmmvc -v
```

## Usage Examples

### Example 1: Create Simple Controller

```bash
dmmvc make:controller About
```

Will create file `internal/controllers/about_controller.go`:

```go
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func AboutController(c *gin.Context) {
    c.HTML(http.StatusOK, "pages/about.html", gin.H{
        "title": "About",
    })
}
```

### Example 2: Create CRUD for Products

```bash
# 1. Create full CRUD
dmmvc make:crud Product

# 2. Edit model, add fields
# internal/models/product.go
type Product struct {
    gorm.Model
    Name        string  `json:"name" gorm:"not null"`
    Description string  `json:"description"`
    Price       float64 `json:"price" gorm:"not null"`
}

# 3. Add migration to internal/database/db.go
db.AutoMigrate(&models.Product{})

# 4. Add routes to internal/routes/routes.go
authorized.GET("/product", controllers.ProductControllerIndex)
authorized.GET("/product/:id", controllers.ProductControllerShow)
authorized.GET("/product/create", controllers.ProductControllerCreate)
authorized.POST("/product", controllers.ProductControllerStore)
authorized.GET("/product/:id/edit", controllers.ProductControllerEdit)
authorized.POST("/product/:id", controllers.ProductControllerUpdate)
authorized.POST("/product/:id/delete", controllers.ProductControllerDelete)

# 5. Customize templates in templates/pages/product/
```

### Example 3: Create Middleware

```bash
dmmvc make:middleware RateLimit
```

Will create file `internal/middleware/rate_limit.go`:

```go
package middleware

import "github.com/gin-gonic/gin"

func RateLimit() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Before request
        c.Next()
        // After request
    }
}
```

## Naming Conventions

### Controllers
- Input name: `Product` or `ProductController`
- File name: `product_controller.go`
- Function name: `ProductController` (simple) or `ProductControllerIndex`, `ProductControllerShow`, etc. (resource)

### Models
- Input name: `Product`
- File name: `product.go`
- Struct name: `Product`

### Middleware
- Input name: `RateLimit`
- File name: `rate_limit.go`
- Function name: `RateLimit`

### Pages
- Input name: `about`
- File name: `about.html`
- Path: `templates/pages/about.html`

## File Structure After Generation

```
dmmvc/
├── internal/
│   ├── controllers/
│   │   ├── product_controller.go    # make:controller Product --resource
│   │   └── about_controller.go      # make:controller About
│   ├── models/
│   │   └── product.go               # make:model Product
│   └── middleware/
│       └── rate_limit.go            # make:middleware RateLimit
└── templates/
    └── pages/
        ├── about.html               # make:page about
        └── product/                 # make:crud Product
            ├── index.html
            ├── show.html
            ├── create.html
            └── edit.html
```

## Tips

1. **Use make:crud** for quick creation of full CRUD functionality
2. **Use --resource** when creating a controller if you plan CRUD operations
3. **Use --migration** when creating a model to remember to add migration
4. **Use dmmvc list** to see all created resources
5. Always review and customize generated code for your needs

## CLI Development

If you want to add new commands to CLI:

1. Open `cmd/cli/main.go`
2. Add new case to switch statement
3. Create handler function
4. Update `printUsage()`
5. Rebuild: `make build`

## Troubleshooting

### CLI not found after installation

Make sure `%GOPATH%\bin` is added to PATH:

```bash
echo %PATH%
```

### File already exists

CLI doesn't overwrite existing files. If you need to recreate a file, delete it manually.

### Errors when creating files

Make sure you're running CLI from the DMMVC project root.
