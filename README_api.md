# Lunch Menu API - Go + PostgreSQL

A clean, modern REST API built with Go and PostgreSQL for restaurant and menu management.

## Database Schema

- **restaurants**: Store restaurant information (name, address, contact, cuisine, etc.)
- **menu_items**: Store menu items linked to restaurants with dietary information

## API Endpoints

### General

- **`GET /api`**: Returns basic API information, including the current version.

### Restaurants

- **`GET /api/restaurants`**: Returns a paginated list of active restaurants.
- **`GET /api/restaurants/{id}`**: Returns a specific restaurant by ID.
- **`GET /api/restaurants/{id}/menu`**: Returns menu items for a specific restaurant.

### Menu Items

- **`GET /api/menu-items/{id}`**: Returns a specific menu item by ID.

### Statistics

- **`GET /api/stats`**: Returns business analytics and statistics.


## Deployment 

### Docker Compose

Docker Compose automatically sets up PostgreSQL and loads the restaurant data:

```bash
# Start all services (PostgreSQL + API with data)
docker compose up -d

# View logs to see initialization
docker compose logs -f

# Stop services
docker compose down
```

### Manual Setup

1. **Start PostgreSQL**
   ```bash
   # Using Docker
   docker run -d --name postgres \
     -e POSTGRES_DB=lunch_menu \
     -e POSTGRES_USER=postgres \
     -p 5432:5432 \
     postgres:15-alpine
   ```

2. **Configure Environment Variables**
   ```bash
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_NAME=lunch_menu
   export DB_SSLMODE=disable
   ```

3. **Run the Application**
   ```bash
   # Install dependencies
   go mod tidy

   # Run the application
   go run .
   ```

## Loading Data

The database starts empty after initial setup. You have several options to populate it with restaurant data:

### Option 1: Docker Compose (Automatic)

**Recommended for development and testing:**

```bash
docker-compose up -d
```

This automatically:
- Creates PostgreSQL database
- Loads all 18 original restaurants
- Starts the API with data ready
- No manual steps required!

### Option 2: SQL Seed Script

Use the provided SQL script with original restaurant data:

```bash
# Using the provided script (recommended)
./seed_database.sh

# Or manually with psql
psql -d lunch_menu -f seed_data.sql
```

This loads 18 restaurants from the original lunch-menu dataset (Uppsala and Solna locations) plus sample menu items.

## Testing the API

You can test the API endpoints using curl:

```bash
# Get API info
curl http://localhost:8000/api

# List all restaurants (supports pagination with limit and offset)
curl "http://localhost:8000/api/restaurants?limit=10&offset=0"

# Get a specific restaurant info
curl http://localhost:8000/api/restaurants/<restaurant-ID>

# Get menu items for a specific restaurant (supports pagination with limit and offset)
curl "http://localhost:8000/api/restaurants/<restaurant-ID>/menu?limit=10&offset=0"

# Get a specific menu item
curl http://localhost:8000/api/menu-items/<menu-item-ID>

# Get business statistics
curl http://localhost:8000/api/stats
```
