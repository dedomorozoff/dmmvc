@echo off
REM DMMVC Test Runner Script

setlocal enabledelayedexpansion

set COVERAGE=false
set VERBOSE=false
set BENCH=false

:parse_args
if "%~1"=="" goto run_tests
if /i "%~1"=="-c" set COVERAGE=true
if /i "%~1"=="--coverage" set COVERAGE=true
if /i "%~1"=="-v" set VERBOSE=true
if /i "%~1"=="--verbose" set VERBOSE=true
if /i "%~1"=="-b" set BENCH=true
if /i "%~1"=="--bench" set BENCH=true
if /i "%~1"=="-h" goto show_help
if /i "%~1"=="--help" goto show_help
shift
goto parse_args

:show_help
echo Usage: test.bat [options]
echo.
echo Options:
echo   -c, --coverage    Run tests with coverage
echo   -v, --verbose     Verbose output
echo   -b, --bench       Run benchmarks
echo   -h, --help        Show this help
exit /b 0

:run_tests
echo ========================================
echo DMMVC Test Runner
echo ========================================
echo.

if "%BENCH%"=="true" (
    echo Running benchmarks...
    go test -bench=. -benchmem ./...
    goto end
)

if "%COVERAGE%"=="true" (
    echo Running tests with coverage...
    if "%VERBOSE%"=="true" (
        go test -v -cover -coverprofile=coverage.out ./...
    ) else (
        go test -cover -coverprofile=coverage.out ./...
    )
    
    echo.
    echo Coverage report:
    go tool cover -func=coverage.out
    
    echo.
    echo Generate HTML coverage report:
    echo   go tool cover -html=coverage.out -o coverage.html
    
) else (
    echo Running tests...
    if "%VERBOSE%"=="true" (
        go test -v ./...
    ) else (
        go test ./...
    )
)

:end
echo.
echo [OK] Tests completed
echo.
