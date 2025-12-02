@echo off
REM DMMVC Quick Start Script

echo ========================================
echo DMMVC Quick Start
echo ========================================
echo.

REM Check if .env exists
if not exist ".env" (
    echo Creating .env file...
    if exist ".env.example" (
        copy .env.example .env >nul
        echo [OK] .env created from .env.example
    ) else (
        echo Creating default .env...
        (
            echo PORT=8080
            echo GIN_MODE=debug
            echo DB_TYPE=sqlite
            echo DB_DSN=dmmvc.db
            echo SESSION_SECRET=change-this-in-production
            echo LOG_LEVEL=info
            echo LOG_FILE=dmmvc.log
            echo DEBUG=true
        ) > .env
        echo [OK] Default .env created
    )
    echo.
)

REM Install dependencies if needed
if not exist "go.sum" (
    echo Installing dependencies...
    go mod download
    go mod tidy
    echo [OK] Dependencies installed
    echo.
)

REM Build CLI if not exists
if not exist "dmmvc.exe" (
    echo Building CLI tool...
    go build -o dmmvc.exe cmd\cli\main.go
    echo [OK] CLI built
    echo.
)

echo ========================================
echo Setup Complete!
echo ========================================
echo.
echo Available commands:
echo   dmmvc.exe --help           - Show CLI help
echo   dmmvc.exe make:crud User   - Generate CRUD
echo   go run cmd\server\main.go  - Start server
echo.
echo Starting development server...
echo.

go run cmd\server\main.go
