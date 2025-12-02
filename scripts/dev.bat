@echo off
REM DMMVC Development Helper Script

setlocal enabledelayedexpansion

if "%~1"=="" goto :help

if "%~1"=="install" goto :install
if "%~1"=="build" goto :build
if "%~1"=="cli" goto :cli
if "%~1"=="server" goto :server
if "%~1"=="run" goto :run
if "%~1"=="test" goto :test
if "%~1"=="fmt" goto :fmt
if "%~1"=="format" goto :fmt
if "%~1"=="lint" goto :lint
if "%~1"=="swagger" goto :swagger
if "%~1"=="clean" goto :clean
if "%~1"=="help" goto :help
goto :help

:check_env
if not exist ".env" (
    echo [ERROR] .env file not found
    if exist ".env.example" (
        echo Creating .env from .env.example...
        copy .env.example .env >nul
        echo [OK] .env created
    ) else (
        echo [ERROR] .env.example not found
        exit /b 1
    )
)
exit /b 0

:install
echo [INFO] Installing dependencies...
go mod download
go mod tidy
echo [OK] Dependencies installed
exit /b 0

:cli
echo [INFO] Building CLI...
go build -o dmmvc.exe cmd\cli\main.go
echo [OK] CLI built: dmmvc.exe
exit /b 0

:server
echo [INFO] Building server...
go build -o server.exe cmd\server\main.go
echo [OK] Server built: server.exe
exit /b 0

:build
call :cli
call :server
exit /b 0

:run
call :check_env
echo [INFO] Starting server...
go run cmd\server\main.go
exit /b 0

:test
echo [INFO] Running tests...
go test ./... -v
exit /b 0

:fmt
echo [INFO] Formatting code...
go fmt ./...
echo [OK] Code formatted
exit /b 0

:lint
echo [INFO] Linting code...
where golangci-lint >nul 2>&1
if %errorlevel% equ 0 (
    golangci-lint run
    echo [OK] Linting complete
) else (
    echo [ERROR] golangci-lint not installed
    echo Install: https://golangci-lint.run/usage/install/
)
exit /b 0

:swagger
echo [INFO] Generating Swagger documentation...
where swag >nul 2>&1
if %errorlevel% equ 0 (
    swag init -g cmd\server\main.go -o docs\swagger --parseDependency --parseInternal
    echo [OK] Swagger docs generated
) else (
    echo [ERROR] swag not installed
    echo Install: go install github.com/swaggo/swag/cmd/swag@latest
)
exit /b 0

:clean
echo [INFO] Cleaning build artifacts...
if exist dmmvc.exe del dmmvc.exe
if exist server.exe del server.exe
if exist *.db del *.db
if exist *.log del *.log
echo [OK] Clean complete
exit /b 0

:help
echo DMMVC Development Helper
echo.
echo Usage: dev.bat [command]
echo.
echo Commands:
echo   install     - Install dependencies
echo   build       - Build CLI and server
echo   cli         - Build CLI only
echo   server      - Build server only
echo   run         - Run development server
echo   test        - Run tests
echo   fmt         - Format code
echo   lint        - Lint code
echo   swagger     - Generate Swagger docs
echo   clean       - Clean build artifacts
echo   help        - Show this help
echo.
exit /b 0
