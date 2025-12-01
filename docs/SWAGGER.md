**English** | [Русский](SWAGGER.ru.md)

# Swagger API Documentation

DMMVC includes built-in support for Swagger/OpenAPI documentation, allowing you to automatically generate interactive API documentation.

## Quick Start

### 1. Generate Documentation

```bash
make swagger
```

Or manually:

```bash
swag init -g cmd/server/main.go -o docs/swagger --parseDependency --parseInternal
```

### 2. Start Server

```bash
make run
```

### 3. Access Swagger UI

Open in browser: **http://localhost:8080/swagger/index.html**

## Documenting Your API

### Basic Controller Example

```go
// GetUser godoc
// @Summary Get user by ID
// @Description Get user details by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} APIResponse{data=models.User}
// @Failure 404 {object} APIResponse
// @Router /api/users/{id} [get]
// @Security SessionAuth
func GetUser(c *gin.Context) {
    // Your implementation
}
```

### Annotation Tags

- `@Summary` - Short description
- `@Description` - Detailed description
- `@Tags` - Group endpoints by tags
- `@Accept` - Request content type (json, xml, etc.)
- `@Produce` - Response content type
- `@Param` - Parameter definition
- `@Success` - Success response
- `@Failure` - Error response
- `@Router` - Route path and HTTP method
- `@Security` - Security scheme

### Parameter Types

```go
// Path parameter
// @Param id path int true "User ID"

// Query parameter
// @Param page query int false "Page number"

// Body parameter
// @Param user body UserCreateRequest true "User data"

// Header parameter
// @Param Authorization header string true "Bearer token"
```

### Response Examples

```go
// Simple response
// @Success 200 {object} models.User

// Response with nested data
// @Success 200 {object} APIResponse{data=models.User}

// Array response
// @Success 200 {object} APIResponse{data=[]models.User}

// Multiple status codes
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 404 {object} APIResponse
// @Failure 500 {object} APIResponse
```

## API Response Structure

Standard API response format:

```go
type APIResponse struct {
    Success bool        `json:"success" example:"true"`
    Message string      `json:"message,omitempty" example:"Operation successful"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty" example:"Error message"`
}
```

## Model Documentation

Add Swagger annotations to your models:

```go
// User model
// @Description User account information
type User struct {
    ID       uint   `json:"id" example:"1"`
    Username string `json:"username" example:"john_doe"`
    Email    string `json:"email" example:"john@example.com"`
    Role     string `json:"role" example:"user"`
}
```

## Request Models

Document request structures:

```go
type UserCreateRequest struct {
    Username string `json:"username" binding:"required" example:"john_doe"`
    Email    string `json:"email" binding:"required,email" example:"john@example.com"`
    Password string `json:"password" binding:"required,min=6" example:"password123"`
}
```

## Security Schemes

DMMVC uses session-based authentication by default:

```go
// @securityDefinitions.apikey SessionAuth
// @in cookie
// @name session_token
```

To require authentication for an endpoint:

```go
// @Security SessionAuth
```

## General API Information

Configure in `cmd/server/main.go`:

```go
// @title DMMVC API
// @version 1.0
// @description Lightweight MVC Web Framework API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@dmmvc.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
```

## Example API Endpoints

DMMVC includes example API endpoints in `internal/controllers/api_example.go`:

- `GET /api/users` - List all users
- `GET /api/users/:id` - Get user by ID
- `POST /api/users` - Create new user
- `DELETE /api/users/:id` - Delete user

## Testing API

Use Swagger UI to test your API:

1. Open http://localhost:8080/swagger/index.html
2. Click on an endpoint
3. Click "Try it out"
4. Fill in parameters
5. Click "Execute"

## Regenerating Documentation

After making changes to your API:

```bash
make swagger
```

The documentation will be automatically updated.

## Customization

### Custom Theme

Edit `internal/routes/routes.go`:

```go
url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
```

### Hide Swagger in Production

```go
if os.Getenv("GIN_MODE") != "release" {
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
```

## Best Practices

1. **Document all public APIs** - Every API endpoint should have Swagger annotations
2. **Use examples** - Add `example` tags to help users understand expected values
3. **Group by tags** - Use `@Tags` to organize endpoints logically
4. **Document errors** - Include all possible error responses
5. **Keep it updated** - Run `make swagger` after API changes
6. **Use standard responses** - Stick to the APIResponse structure for consistency

## Troubleshooting

### Documentation not updating

```bash
# Clean and regenerate
rm -rf docs/swagger
make swagger
```

### Swagger UI not loading

Check that the import is present in `internal/routes/routes.go`:

```go
_ "dmmvc/docs/swagger"
```

### Type definition errors

Use `--parseDependency --parseInternal` flags:

```bash
swag init -g cmd/server/main.go -o docs/swagger --parseDependency --parseInternal
```

## Resources

- [Swaggo Documentation](https://github.com/swaggo/swag)
- [OpenAPI Specification](https://swagger.io/specification/)
- [Swagger UI](https://swagger.io/tools/swagger-ui/)

## Next Steps

- Add more API endpoints
- Document existing controllers
- Customize response formats
- Add API versioning
- Export to Postman/Insomnia
