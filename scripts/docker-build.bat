@echo off
REM DMMVC Docker Build Script

set IMAGE_NAME=%1
set IMAGE_TAG=%2

if "%IMAGE_NAME%"=="" set IMAGE_NAME=dmmvc
if "%IMAGE_TAG%"=="" set IMAGE_TAG=latest

set FULL_IMAGE=%IMAGE_NAME%:%IMAGE_TAG%

echo ========================================
echo Building Docker Image
echo ========================================
echo.
echo Image: %FULL_IMAGE%
echo.

REM Check if Dockerfile exists
if not exist "Dockerfile" (
    echo [ERROR] Dockerfile not found
    exit /b 1
)

echo [1/3] Building Docker image...
docker build -t %FULL_IMAGE% .
echo.

echo [2/3] Checking image...
docker images | findstr %IMAGE_NAME%
echo.

echo [3/3] Image built successfully!
echo.
echo ========================================
echo Next Steps
echo ========================================
echo.
echo Run container:
echo   docker run -p 8080:8080 %FULL_IMAGE%
echo.
echo Run with environment file:
echo   docker run -p 8080:8080 --env-file .env %FULL_IMAGE%
echo.
echo Run in background:
echo   docker run -d -p 8080:8080 --name dmmvc-app %FULL_IMAGE%
echo.
echo View logs:
echo   docker logs dmmvc-app
echo.
echo Stop container:
echo   docker stop dmmvc-app
echo.
