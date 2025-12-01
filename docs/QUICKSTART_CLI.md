# DMMVC CLI Quick Start Guide

Get started with DMMVC CLI in 5 minutes!

## Step 1: Build the CLI

```bash
make build
```

This creates `dmmvc.exe` in your project root.

## Step 2: Create Your First CRUD

Let's create a simple blog post management system:

```bash
dmmvc make:crud Post
```

This generates:
- Model: `internal/models/post.go`
- Controller: `internal/controllers/post_controller.go` (with 7 CRUD methods)
- Templates: `templates/pages/post/*.html` (index, show, create, edit)

## Step 3: Customize the Model

Edit `internal/models/post.go`:

```go
type Post struct {
    gorm.Model
    Title   string `json:"title" gorm:"not null"`
    Content string `json:"content" gorm:"type:text"`
    Author  string `json:"author"`
}
```

## Step 4: Add Migration

Edit `internal/database/db.go`, add to `InitDB()`:

```go
db.AutoMigrate(&models.Post{})
```

## Step 5: Add Routes

Edit `internal/routes/routes.go`, add these routes:

```go
// Post routes
authorized.GET("/post", controllers.PostControllerIndex)
authorized.GET("/post/:id", controllers.PostControllerShow)
authorized.GET("/post/create", controllers.PostControllerCreate)
authorized.POST("/post", controllers.PostControllerStore)
authorized.GET("/post/:id/edit", controllers.PostControllerEdit)
authorized.POST("/post/:id", controllers.PostControllerUpdate)
authorized.POST("/post/:id/delete", controllers.PostControllerDelete)
```

## Step 6: Customize Templates

Edit templates in `templates/pages/post/`:

### `create.html` - Add form fields:

```html
<form method="POST" action="/post">
    <div class="form-group">
        <label>Title</label>
        <input type="text" name="title" required>
    </div>
    <div class="form-group">
        <label>Content</label>
        <textarea name="content" rows="10" required></textarea>
    </div>
    <div class="form-group">
        <label>Author</label>
        <input type="text" name="author" required>
    </div>
    <button type="submit">Create Post</button>
</form>
```

### `index.html` - Update table columns:

```html
<thead>
    <tr>
        <th>ID</th>
        <th>Title</th>
        <th>Author</th>
        <th>Created</th>
        <th>Actions</th>
    </tr>
</thead>
<tbody>
    {{range .items}}
    <tr>
        <td>{{.ID}}</td>
        <td>{{.Title}}</td>
        <td>{{.Author}}</td>
        <td>{{.CreatedAt}}</td>
        <td>
            <a href="/post/{{.ID}}">View</a>
            <a href="/post/{{.ID}}/edit">Edit</a>
            <form method="POST" action="/post/{{.ID}}/delete" style="display:inline;">
                <button type="submit">Delete</button>
            </form>
        </td>
    </tr>
    {{end}}
</tbody>
```

## Step 7: Run and Test

```bash
go run cmd/server/main.go
```

Visit: `http://localhost:8080/post`

## More CLI Commands

```bash
# Create a simple controller
dmmvc make:controller About

# Create a model with migration hint
dmmvc make:model Category --migration

# Create middleware
dmmvc make:middleware RateLimit

# Create a page template
dmmvc make:page contact

# List all resources
dmmvc list

# Show help
dmmvc --help
```

## Tips

- **Use `make:crud`** for quick scaffolding, then customize
- **Use `--resource`** flag for controllers with CRUD methods
- **Use `list`** command to see all your resources
- **Always review** generated code and adapt to your needs

## Next Steps

- Read full documentation: [CLI.md](CLI.md)
- Check examples: [EXAMPLES.md](EXAMPLES.md)
- Learn about architecture: [ARCHITECTURE.md](ARCHITECTURE.md)
- Back to [Documentation Index](README.md)

Happy coding!
