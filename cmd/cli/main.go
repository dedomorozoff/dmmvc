package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	// Handle version and help flags
	if command == "--version" || command == "-v" {
		fmt.Printf("DMMVC CLI v%s\n", version)
		return
	}

	if command == "--help" || command == "-h" {
		printUsage()
		return
	}

	args := os.Args[2:]

	switch command {
	case "make:controller":
		if len(args) < 1 {
			fmt.Println("Usage: dmmvc make:controller <ControllerName> [--resource]")
			return
		}
		resourceFlag := false
		if len(args) > 1 && args[1] == "--resource" {
			resourceFlag = true
		}
		makeController(args[0], resourceFlag)
	case "make:model":
		if len(args) < 1 {
			fmt.Println("Usage: dmmvc make:model <ModelName> [--migration]")
			return
		}
		migrationFlag := false
		if len(args) > 1 && args[1] == "--migration" {
			migrationFlag = true
		}
		makeModel(args[0], migrationFlag)
	case "make:middleware":
		if len(args) < 1 {
			fmt.Println("Usage: dmmvc make:middleware <MiddlewareName>")
			return
		}
		makeMiddleware(args[0])
	case "make:page":
		if len(args) < 1 {
			fmt.Println("Usage: dmmvc make:page <PageName>")
			return
		}
		makePage(args[0])
	case "make:crud":
		if len(args) < 1 {
			fmt.Println("Usage: dmmvc make:crud <Name>")
			return
		}
		makeCRUD(args[0])
	case "list":
		listResources()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println("╔═══════════════════════════════════════╗")
	fmt.Println("║       DMMVC CLI Tool v" + version + "        ║")
	fmt.Println("╚═══════════════════════════════════════╝")
	fmt.Println("\nUsage:")
	fmt.Println("  dmmvc <command> [arguments] [flags]")
	fmt.Println("\nCommands:")
	fmt.Println("  make:controller <Name>      Create a new controller")
	fmt.Println("    --resource                Create a resource controller with CRUD methods")
	fmt.Println("  make:model <Name>           Create a new model")
	fmt.Println("    --migration               Also create a migration file")
	fmt.Println("  make:middleware <Name>      Create a new middleware")
	fmt.Println("  make:page <Name>            Create a new page template")
	fmt.Println("  make:crud <Name>            Create model, controller, and pages for CRUD")
	fmt.Println("  list                        List all controllers, models, and middleware")
	fmt.Println("\nFlags:")
	fmt.Println("  -h, --help                  Show this help message")
	fmt.Println("  -v, --version               Show version")
	fmt.Println("\nExamples:")
	fmt.Println("  dmmvc make:controller Product")
	fmt.Println("  dmmvc make:controller Product --resource")
	fmt.Println("  dmmvc make:model Product --migration")
	fmt.Println("  dmmvc make:crud Product")
}

func toSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

func createFile(path string, tmpl string, data interface{}) {
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("File already exists: %s\n", path)
		return
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("Error creating directory %s: %v\n", dir, err)
		return
	}

	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", path, err)
		return
	}
	defer f.Close()

	t, err := template.New("file").Parse(tmpl)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	if err := t.Execute(f, data); err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return
	}

	fmt.Printf("Created: %s\n", path)
}

func makeController(name string, resource bool) {
	baseName := name
	if strings.HasSuffix(name, "Controller") {
		baseName = strings.TrimSuffix(name, "Controller")
	}
	controllerName := baseName + "Controller"
	fileName := toSnakeCase(baseName) + "_controller.go"
	path := filepath.Join("internal", "controllers", fileName)

	var tmpl string
	if resource {
		tmpl = `package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"dmmvc/internal/database"
	"dmmvc/internal/models"
)

// Index displays list of {{.BaseName}}
func {{.Name}}Index(c *gin.Context) {
	var items []models.{{.BaseName}}
	database.DB.Find(&items)
	
	c.HTML(http.StatusOK, "pages/{{.PageName}}/index.html", gin.H{
		"title": "{{.Title}} List",
		"items": items,
	})
}

// Show displays a single {{.BaseName}}
func {{.Name}}Show(c *gin.Context) {
	id := c.Param("id")
	var item models.{{.BaseName}}
	
	if err := database.DB.First(&item, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "pages/error.html", gin.H{
			"error": "{{.BaseName}} not found",
		})
		return
	}
	
	c.HTML(http.StatusOK, "pages/{{.PageName}}/show.html", gin.H{
		"title": "{{.Title}} Details",
		"item":  item,
	})
}

// Create displays form for creating {{.BaseName}}
func {{.Name}}Create(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/{{.PageName}}/create.html", gin.H{
		"title": "Create {{.Title}}",
	})
}

// Store saves new {{.BaseName}}
func {{.Name}}Store(c *gin.Context) {
	var item models.{{.BaseName}}
	
	if err := c.ShouldBind(&item); err != nil {
		c.HTML(http.StatusBadRequest, "pages/{{.PageName}}/create.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	
	database.DB.Create(&item)
	c.Redirect(http.StatusFound, "/{{.PageName}}")
}

// Edit displays form for editing {{.BaseName}}
func {{.Name}}Edit(c *gin.Context) {
	id := c.Param("id")
	var item models.{{.BaseName}}
	
	if err := database.DB.First(&item, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "pages/error.html", gin.H{
			"error": "{{.BaseName}} not found",
		})
		return
	}
	
	c.HTML(http.StatusOK, "pages/{{.PageName}}/edit.html", gin.H{
		"title": "Edit {{.Title}}",
		"item":  item,
	})
}

// Update saves changes to {{.BaseName}}
func {{.Name}}Update(c *gin.Context) {
	id := c.Param("id")
	var item models.{{.BaseName}}
	
	if err := database.DB.First(&item, id).Error; err != nil {
		c.HTML(http.StatusNotFound, "pages/error.html", gin.H{
			"error": "{{.BaseName}} not found",
		})
		return
	}
	
	if err := c.ShouldBind(&item); err != nil {
		c.HTML(http.StatusBadRequest, "pages/{{.PageName}}/edit.html", gin.H{
			"error": err.Error(),
			"item":  item,
		})
		return
	}
	
	database.DB.Save(&item)
	c.Redirect(http.StatusFound, "/{{.PageName}}")
}

// Delete removes {{.BaseName}}
func {{.Name}}Delete(c *gin.Context) {
	id := c.Param("id")
	var item models.{{.BaseName}}
	
	if err := database.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "{{.BaseName}} not found"})
		return
	}
	
	database.DB.Delete(&item)
	c.Redirect(http.StatusFound, "/{{.PageName}}")
}
`
	} else {
		tmpl = `package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func {{.Name}}(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/{{.PageName}}.html", gin.H{
		"title": "{{.Title}}",
	})
}
`
	}

	data := struct {
		Name     string
		BaseName string
		PageName string
		Title    string
	}{
		Name:     controllerName,
		BaseName: baseName,
		PageName: toSnakeCase(baseName),
		Title:    baseName,
	}

	createFile(path, tmpl, data)
	
	if resource {
		fmt.Println("\n✓ Resource controller created!")
		fmt.Println("\nAdd these routes to internal/routes/routes.go:")
		fmt.Printf("  authorized.GET(\"/%s\", controllers.%sIndex)\n", toSnakeCase(baseName), controllerName)
		fmt.Printf("  authorized.GET(\"/%s/:id\", controllers.%sShow)\n", toSnakeCase(baseName), controllerName)
		fmt.Printf("  authorized.GET(\"/%s/create\", controllers.%sCreate)\n", toSnakeCase(baseName), controllerName)
		fmt.Printf("  authorized.POST(\"/%s\", controllers.%sStore)\n", toSnakeCase(baseName), controllerName)
		fmt.Printf("  authorized.GET(\"/%s/:id/edit\", controllers.%sEdit)\n", toSnakeCase(baseName), controllerName)
		fmt.Printf("  authorized.POST(\"/%s/:id\", controllers.%sUpdate)\n", toSnakeCase(baseName), controllerName)
		fmt.Printf("  authorized.POST(\"/%s/:id/delete\", controllers.%sDelete)\n", toSnakeCase(baseName), controllerName)
	}
}

func makeModel(name string, migration bool) {
	fileName := toSnakeCase(name) + ".go"
	path := filepath.Join("internal", "models", fileName)

	tmpl := `package models

import "gorm.io/gorm"

type {{.Name}} struct {
	gorm.Model
	// Add your fields here
	// Example:
	// Name        string ` + "`json:\"name\" gorm:\"not null\"`" + `
	// Description string ` + "`json:\"description\"`" + `
}
`
	data := struct {
		Name string
	}{
		Name: name,
	}

	createFile(path, tmpl, data)
	
	if migration {
		fmt.Println("\n✓ Model created!")
		fmt.Println("\nTo run migration, add this to internal/database/db.go in InitDB():")
		fmt.Printf("  db.AutoMigrate(&models.%s{})\n", name)
	}
}

func makeMiddleware(name string) {
	fileName := toSnakeCase(name) + ".go"
	path := filepath.Join("internal", "middleware", fileName)

	tmpl := `package middleware

import "github.com/gin-gonic/gin"

func {{.Name}}() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Before request
		c.Next()
		// After request
	}
}
`
	data := struct {
		Name string
	}{
		Name: name,
	}

	createFile(path, tmpl, data)
}

func makePage(name string) {
	fileName := toSnakeCase(name) + ".html"
	path := filepath.Join("templates", "pages", fileName)

	tmpl := `{{define "pages/{{.Name}}.html"}}
{{template "layouts/base.html" .}}

{{define "content"}}
<div class="container">
    <h1>{{.Title}}</h1>
    <p>Welcome to {{.Title}}</p>
</div>
{{end}}
{{end}}
`
	data := struct {
		Name  string
		Title string
	}{
		Name:  toSnakeCase(name),
		Title: name,
	}

	createFile(path, tmpl, data)
}

func makeCRUD(name string) {
	fmt.Printf("Creating CRUD for %s...\n\n", name)
	
	// Create model
	makeModel(name, true)
	
	// Create resource controller
	makeController(name, true)
	
	// Create pages
	baseName := toSnakeCase(name)
	pagesDir := filepath.Join("templates", "pages", baseName)
	
	// Index page
	indexTmpl := `{{define "pages/` + baseName + `/index.html"}}
{{template "layouts/base.html" .}}

{{define "content"}}
<div class="container">
    <div class="header">
        <h1>{{.title}}</h1>
        <a href="/` + baseName + `/create" class="btn btn-primary">Create New</a>
    </div>
    
    <table class="table">
        <thead>
            <tr>
                <th>ID</th>
                <th>Created At</th>
                <th>Actions</th>
            </tr>
        </thead>
        <tbody>
            {{range .items}}
            <tr>
                <td>{{.ID}}</td>
                <td>{{.CreatedAt}}</td>
                <td>
                    <a href="/` + baseName + `/{{.ID}}">View</a>
                    <a href="/` + baseName + `/{{.ID}}/edit">Edit</a>
                    <form method="POST" action="/` + baseName + `/{{.ID}}/delete" style="display:inline;">
                        <button type="submit" onclick="return confirm('Are you sure?')">Delete</button>
                    </form>
                </td>
            </tr>
            {{end}}
        </tbody>
    </table>
</div>
{{end}}
{{end}}
`
	createFile(filepath.Join(pagesDir, "index.html"), indexTmpl, nil)
	
	// Create page
	createTmpl := `{{define "pages/` + baseName + `/create.html"}}
{{template "layouts/base.html" .}}

{{define "content"}}
<div class="container">
    <h1>{{.title}}</h1>
    
    <form method="POST" action="/` + baseName + `">
        <!-- Add your form fields here -->
        
        <button type="submit" class="btn btn-primary">Create</button>
        <a href="/` + baseName + `" class="btn btn-secondary">Cancel</a>
    </form>
</div>
{{end}}
{{end}}
`
	createFile(filepath.Join(pagesDir, "create.html"), createTmpl, nil)
	
	// Edit page
	editTmpl := `{{define "pages/` + baseName + `/edit.html"}}
{{template "layouts/base.html" .}}

{{define "content"}}
<div class="container">
    <h1>{{.title}}</h1>
    
    <form method="POST" action="/` + baseName + `/{{.item.ID}}">
        <!-- Add your form fields here -->
        
        <button type="submit" class="btn btn-primary">Update</button>
        <a href="/` + baseName + `" class="btn btn-secondary">Cancel</a>
    </form>
</div>
{{end}}
{{end}}
`
	createFile(filepath.Join(pagesDir, "edit.html"), editTmpl, nil)
	
	// Show page
	showTmpl := `{{define "pages/` + baseName + `/show.html"}}
{{template "layouts/base.html" .}}

{{define "content"}}
<div class="container">
    <h1>{{.title}}</h1>
    
    <div class="details">
        <p><strong>ID:</strong> {{.item.ID}}</p>
        <p><strong>Created At:</strong> {{.item.CreatedAt}}</p>
        <!-- Add more fields here -->
    </div>
    
    <a href="/` + baseName + `/{{.item.ID}}/edit" class="btn btn-primary">Edit</a>
    <a href="/` + baseName + `" class="btn btn-secondary">Back to List</a>
</div>
{{end}}
{{end}}
`
	createFile(filepath.Join(pagesDir, "show.html"), showTmpl, nil)
	
	fmt.Println("\n✓ CRUD scaffolding complete!")
	fmt.Println("\nNext steps:")
	fmt.Println("1. Add fields to the model in internal/models/" + toSnakeCase(name) + ".go")
	fmt.Println("2. Add migration to internal/database/db.go")
	fmt.Println("3. Add routes to internal/routes/routes.go (see controller output above)")
	fmt.Println("4. Customize the templates in templates/pages/" + baseName + "/")
}

func listResources() {
	fmt.Println("╔═══════════════════════════════════════╗")
	fmt.Println("║         Project Resources            ║")
	fmt.Println("╚═══════════════════════════════════════╝")
	
	// List controllers
	fmt.Println("\n📁 Controllers:")
	listFiles("internal/controllers", "_controller.go")
	
	// List models
	fmt.Println("\n📁 Models:")
	listFiles("internal/models", ".go")
	
	// List middleware
	fmt.Println("\n📁 Middleware:")
	listFiles("internal/middleware", ".go")
	
	// List pages
	fmt.Println("\n📁 Pages:")
	listFiles("templates/pages", ".html")
}

func listFiles(dir string, suffix string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("  (directory not found or empty)\n")
		return
	}
	
	count := 0
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.HasSuffix(entry.Name(), suffix) {
			fmt.Printf("  - %s\n", entry.Name())
			count++
		}
	}
	
	if count == 0 {
		fmt.Println("  (no files found)")
	}
}
