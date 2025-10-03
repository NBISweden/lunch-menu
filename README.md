# Lunch Menu API - Go + PostgreSQL

A clean, modern REST API built with Go and PostgreSQL for restaurant and menu management.

## Database Schema

- **restaurants**: Store restaurant information (name, address, contact, cuisine, etc.)
- **menu_items**: Store menu items linked to restaurants with dietary information

## API Endpoints

### Restaurants

#### `GET /api/restaurants`
Returns a paginated list of active restaurants.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "name": "Restaurant Name",
      "address": "123 Main St",
      "phone": "+1234567890",
      "email": "contact@restaurant.com",
      "website": "https://restaurant.com",
      "cuisine": "Italian",
      "rating": 4.5,
      "price_range": "$$",
      "opening_time": "11:00",
      "closing_time": "22:00",
      "coordinates": [59.3293, 18.0686],
      "is_active": true,
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "total": 25,
  "page": 1,
  "limit": 10,
  "total_pages": 3
}
{
  "restaurants": [
    {
      "coordinate": [59.35125, 18.02999],
      "homepage": "http://www.dengladarestaurangen.se/",
      "identifier": "glada",
      "menuUrl": "http://www.dengladarestaurangen.se/#!meny/c30g",
      "name": "Den Glada Restaurangen",
      "region": "Solna"
    }
    // ... more restaurants
  ]
}
```

#### `GET /api/restaurants/{id}`
Returns a specific restaurant by ID.

**Response:**
```json
{
  "id": 1,
  "name": "Restaurant Name",
  "address": "123 Main St",
  "phone": "+1234567890",
  "email": "contact@restaurant.com",
  "website": "https://restaurant.com",
  "cuisine": "Italian",
  "rating": 4.5,
  "price_range": "$$",
  "opening_time": "11:00",
  "closing_time": "22:00",
  "coordinates": [59.3293, 18.0686],
  "is_active": true,
  "created_at": "2024-01-01T10:00:00Z",
  "updated_at": "2024-01-01T10:00:00Z"
}
```

### Menu Items

#### `GET /api/restaurants/{restaurant_id}/menu-items`
Returns menu items for a specific restaurant.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "restaurant_id": 1,
      "name": "Spaghetti Carbonara",
      "description": "Classic Italian pasta dish",
      "price": 125.00,
      "category": "Main Course",
      "is_vegetarian": false,
      "is_vegan": false,
      "is_gluten_free": false,
      "calories": 650,
      "is_available": true,
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "total": 15,
  "page": 1,
  "limit": 10,
  "total_pages": 2
}
```

## Environment Configuration

The application uses environment variables for configuration:

```bash
# Database Configuration
DB_HOST=localhost          # PostgreSQL host
DB_PORT=5432              # PostgreSQL port
DB_USER=postgres          # Database username
DB_PASSWORD=postgres      # Database password
DB_NAME=lunch_menu        # Database name
DB_SSLMODE=disable        # SSL mode (disable for local dev)

# Server Configuration  
PORT=8080                 # HTTP server port
GIN_MODE=release          # Gin mode (debug/release)
```

## Development

### Prerequisites

- Go 1.23 or later
- PostgreSQL 13 or later
- Docker and Docker Compose (optional)

## Deployment

### Docker Compose (Recommended)

Docker Compose automatically sets up PostgreSQL and loads the restaurant data:

```bash
# Start all services (PostgreSQL + API with data)
docker-compose up -d

# View logs to see initialization
docker-compose logs -f

# Stop services
docker-compose down
```

**What happens automatically:**
1. PostgreSQL container starts and creates the database
2. Restaurant data is automatically loaded from `init_db.sql`
3. Go API starts and connects to the pre-populated database
4. API is ready at `http://localhost:8000` with 18 restaurants loaded

**Database persistence:**
- Data persists between container restarts
- To reset data: `docker-compose down -v` (removes volumes)

### Manual Deployment

1. **Build the application:**
   ```bash
   CGO_ENABLED=0 GOOS=linux go build -o lunch-menu-api .
   ```

2. **Set up PostgreSQL database**

3. **Configure environment variables**

5. **Run the application:**
   ```bash
   ./lunch-menu-api
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


## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

### Manual Setup

1. **Start PostgreSQL**
   ```bash
   # Using Docker
   docker run -d --name postgres \
     -e POSTGRES_DB=lunch_menu \
     -e POSTGRES_USER=postgres \
     -e POSTGRES_PASSWORD=postgres \
     -p 5432:5432 \
     postgres:15-alpine
   ```

2. **Configure Environment Variables**
   ```bash
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_PASSWORD=postgres
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

## Testing the API

You can test the API endpoints using curl:

```bash
# Get API info
curl http://localhost:8000/api

# List all restaurants
curl http://localhost:8000/api/restaurants

# Get a specific restaurant info
curl http://localhost:8000/api/restaurants/<restaurant-ID>

# Get dish description
curl http://localhost:8000/api/menu-items/<menu-item-ID>
```

