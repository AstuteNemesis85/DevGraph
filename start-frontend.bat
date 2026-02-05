@echo off
echo ========================================
echo   DevGraph - Starting Frontend
echo ========================================
echo.

cd /d "%~dp0\frontend"

if not exist "node_modules" (
    echo Installing npm dependencies...
    call npm install
    echo.
)

echo Starting React development server...
echo.
echo Frontend App: http://localhost:3000
echo.
echo Press Ctrl+C to stop the server
echo ========================================
echo.

call npm run dev
