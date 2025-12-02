@echo off
REM DMMVC Project Creation Script

if "%~1"=="" (
    echo Usage: create-project.bat [project-name]
    echo Example: create-project.bat my-app
    exit /b 1
)

set PROJECT_NAME=%~1
set PROJECT_DIR=%CD%\%PROJECT_NAME%

echo ========================================
echo Creating new DMMVC project: %PROJECT_NAME%
echo ========================================
echo.

REM Check if directory exists
if exist "%PROJECT_DIR%" (
    echo [ERROR] Directory '%PROJECT_NAME%' already exists
    exit /b 1
)

echo [1/5] Creating project directory...
mkdir "%PROJECT_DIR%"
cd "%PROJECT_DIR%"
echo.

echo [2/5] Initializing Go module...
go mod init %PROJECT_NAME%
echo.

echo [3/5] Adding DMMVC framework...
go get github.com/dedomorozoff/dmmvc@latest
echo.

echo [4/5] Creating project structure...
mkdir cmd\server
mkdir internal\controllers
mkdir internal\models
mkdir static\css
mkdir static\js
mkdir templates\layouts
mkdir templates\pages
mkdir templates\partials
echo.

echo [5/5] Creating initial files...

REM Create main.go
(
echo package main
echo.
echo import ^(
echo     "%PROJECT_NAME%/internal/controllers"
echo     "github.com/dedomorozoff/dmmvc/pkg/database"
echo     "github.com/dedomorozoff/dmmvc/pkg/logger"
echo     "github.com/gin-gonic/gin"
echo     "github.com/joho/godotenv"
echo     "log"
echo     "os"
echo ^)
echo.
echo func main^(^) {
echo     // Load .env
echo     godotenv.Load^(^)
echo.
echo     // Initialize logger
echo     logger.Init^(^)
echo.
echo     // Initialize database
echo     database.Init^(^)
echo.
echo     // Setup Gin
echo     r := gin.Default^(^)
echo.
echo     // Load templates
echo     r.LoadHTMLGlob^("templates/**/*"^)
echo     r.Static^("/static", "./static"^)
echo.
echo     // Setup routes
echo     r.GET^("/", controllers.HomeHandler^)
echo.
echo     // Start server
echo     port := os.Getenv^("PORT"^)
echo     if port == "" {
echo         port = "8080"
echo     }
echo     log.Printf^("Server starting on port %%s", port^)
echo     log.Fatal^(r.Run^(":" + port^)^)
echo }
) > cmd\server\main.go

REM Create .env
(
echo PORT=8080
echo GIN_MODE=debug
echo DB_TYPE=sqlite
echo DB_DSN=%PROJECT_NAME%.db
echo SESSION_SECRET=change-this-secret-key
echo LOG_LEVEL=info
echo LOG_FILE=%PROJECT_NAME%.log
echo DEBUG=true
) > .env

REM Create .gitignore
(
echo # Binaries
echo *.exe
echo *.dll
echo *.so
echo *.dylib
echo.
echo # Test files
echo *.test
echo.
echo # Output
echo *.out
echo.
echo # Database
echo *.db
echo.
echo # Logs
echo *.log
echo.
echo # Environment
echo .env
echo.
echo # IDE
echo .vscode/
echo .idea/
echo.
echo # Uploads
echo uploads/
echo.
echo # Temp
echo tmp/
) > .gitignore

REM Create home controller
(
echo package controllers
echo.
echo import ^(
echo     "github.com/gin-gonic/gin"
echo     "net/http"
echo ^)
echo.
echo func HomeHandler^(c *gin.Context^) {
echo     c.HTML^(http.StatusOK, "home.html", gin.H{
echo         "title": "Welcome",
echo     }^)
echo }
) > internal\controllers\home.go

REM Create base layout
(
echo ^<!DOCTYPE html^>
echo ^<html lang="en"^>
echo ^<head^>
echo     ^<meta charset="UTF-8"^>
echo     ^<meta name="viewport" content="width=device-width, initial-scale=1.0"^>
echo     ^<title^>{{ .title }}^</title^>
echo     ^<link rel="stylesheet" href="/static/css/style.css"^>
echo ^</head^>
echo ^<body^>
echo     {{ template "content" . }}
echo ^</body^>
echo ^</html^>
) > templates\layouts\base.html

REM Create home page
(
echo {{ define "content" }}
echo ^<div class="container"^>
echo     ^<h1^>Welcome to %PROJECT_NAME%^</h1^>
echo     ^<p^>Your DMMVC application is running!^</p^>
echo ^</div^>
echo {{ end }}
) > templates\pages\home.html

REM Create basic CSS
(
echo body {
echo     font-family: Arial, sans-serif;
echo     margin: 0;
echo     padding: 20px;
echo     background-color: #f5f5f5;
echo }
echo.
echo .container {
echo     max-width: 800px;
echo     margin: 0 auto;
echo     background: white;
echo     padding: 40px;
echo     border-radius: 8px;
echo     box-shadow: 0 2px 4px rgba^(0,0,0,0.1^);
echo }
echo.
echo h1 {
echo     color: #333;
echo }
) > static\css\style.css

REM Create README
(
echo # %PROJECT_NAME%
echo.
echo DMMVC-based web application
echo.
echo ## Quick Start
echo.
echo ```bash
echo # Install dependencies
echo go mod tidy
echo.
echo # Run server
echo go run cmd/server/main.go
echo ```
echo.
echo Open http://localhost:8080
) > README.md

echo.
echo ========================================
echo Project created successfully!
echo ========================================
echo.
echo Next steps:
echo   cd %PROJECT_NAME%
echo   go mod tidy
echo   go run cmd/server/main.go
echo.
echo Happy coding!
