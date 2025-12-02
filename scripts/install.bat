@echo off
REM DMMVC Installation Script for Windows

echo ========================================
echo DMMVC Framework Installation
echo ========================================
echo.

REM Check if Go is installed
where go >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Go is not installed or not in PATH
    echo Please install Go from https://golang.org/dl/
    exit /b 1
)

echo [1/4] Checking Go installation...
go version
echo.

echo [2/5] Installing dependencies...
go mod download
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Failed to download dependencies
    exit /b 1
)
echo.

echo [3/5] Installing Swagger tool...
go install github.com/swaggo/swag/cmd/swag@latest
if %ERRORLEVEL% NEQ 0 (
    echo [WARNING] Failed to install swag, skipping...
) else (
    echo Generating Swagger documentation...
    swag init -g cmd/server/main.go -o docs/swagger --parseDependency --parseInternal
    if %ERRORLEVEL% NEQ 0 (
        echo [WARNING] Failed to generate Swagger docs
    ) else (
        echo [OK] Swagger documentation generated
    )
)
echo.

echo [4/5] Building CLI tool...
go build -o dmmvc.exe cmd/cli/main.go
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Failed to build CLI
    exit /b 1
)
echo.

echo [5/5] Installing CLI globally...
if not defined GOPATH (
    echo [WARNING] GOPATH not set, using default Go bin directory
    go install ./cmd/cli
) else (
    if not exist "%GOPATH%\bin" mkdir "%GOPATH%\bin"
    copy /Y dmmvc.exe "%GOPATH%\bin\dmmvc.exe" >nul
    echo CLI installed to: %GOPATH%\bin\dmmvc.exe
)
echo.

echo ========================================
echo Installation Complete!
echo ========================================
echo.
echo To use DMMVC CLI globally, make sure your Go bin directory is in PATH:
echo   - %GOPATH%\bin (if GOPATH is set)
echo   - %USERPROFILE%\go\bin (default)
echo.
echo Quick start:
echo   1. dmmvc --help              - Show available commands
echo   2. dmmvc make:crud Product   - Generate CRUD for Product
echo   3. go run cmd/server/main.go - Start the server
echo.
echo Documentation: https://github.com/dedomorozoff/dmmvc
echo.
pause
