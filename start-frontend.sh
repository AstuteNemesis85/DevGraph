#!/bin/bash

echo "========================================"
echo "  DevGraph - Starting Frontend"
echo "========================================"
echo ""

cd "$(dirname "$0")/frontend"

if [ ! -d "node_modules" ]; then
    echo "Installing npm dependencies..."
    npm install
    echo ""
fi

echo "Starting React development server..."
echo ""
echo "Frontend App: http://localhost:3000"
echo ""
echo "Press Ctrl+C to stop the server"
echo "========================================"
echo ""

npm run dev
