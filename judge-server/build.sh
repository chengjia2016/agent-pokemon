#!/bin/bash

echo "Building Judge Server..."

cd "$(dirname "$0")"

go mod tidy

go build -o judge-server cmd/main.go

echo "Build complete!"
echo "Run with: ./judge-server"
