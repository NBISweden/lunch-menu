#!/bin/bash

# Test script for the Go lunch menu API with PostgreSQL

echo "🚀 Starting PostgreSQL database..."
docker run -d --name lunch-menu-postgres \
  -e POSTGRES_DB=lunch_menu \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -p 5432:5432 \
  postgres:15-alpine

echo "⏳ Waiting for PostgreSQL to be ready..."
sleep 10

echo "🔧 Setting environment variables..."
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=lunch_menu
export DB_SSLMODE=disable

echo "🏗️  Building the application..."
go build -o lunch-menu-api .

echo "🌟 Starting the lunch menu API..."
echo "The API will:"
echo "  1. Connect to PostgreSQL"
echo "  2. Run database migrations" 
echo "  3. Load restaurant data from JSON and seed the database"
echo "  4. Start the HTTP server on port 8000"
echo ""
echo "📡 API Endpoints will be available at:"
echo "  GET http://localhost:8000/api"
echo "  GET http://localhost:8000/api/restaurant"
echo "  GET http://localhost:8000/api/restaurant/{name}"
echo "  GET http://localhost:8000/api/version"
echo ""
echo "Press Ctrl+C to stop..."

./lunch-menu-api