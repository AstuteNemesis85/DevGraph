#!/bin/bash

echo "========================================"
echo "  DevGraph - Starting Backend"
echo "========================================"
echo ""

cd "$(dirname "$0")"

echo "Installing Go dependencies..."
go mod tidy
echo ""

echo "Starting Go backend on port 8080..."
echo ""
echo "Backend API: http://localhost:8080"
echo ""
echo "Press Ctrl+C to stop the server"
echo "========================================"
echo ""

go run cmd/server/main.go
