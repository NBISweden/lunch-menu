# Lunch Menu API - Go + PostgreSQL

A clean, modern REST API built with Go and PostgreSQL for restaurant and menu management.

## Features

- **REST API** for restaurant and menu item management
- **PostgreSQL database** with GORM ORM
- **CRUD operations** for restaurants and menu items
- **JSON responses** with proper error handling
- **Database migrations** handled automatically
- **Docker support** with docker-compose
- **Environment-based configuration**

## Architecture

The application follows a clean layered architecture:
- **Handlers**: HTTP request/response handling (Gin framework)
- **Models**: Database models with GORM tags and validation
- **Database**: PostgreSQL connection and CRUD operations
- **Utils**: Business logic services
- **Main**: Application entry point and initialization

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

#### `POST /api/restaurants`
Creates a new restaurant.

**Request Body:**
```json
{
  "name": "New Restaurant",
  "address": "456 Oak St",
  "phone": "+1987654321",
  "email": "info@newrestaurant.com",
  "website": "https://newrestaurant.com",
  "cuisine": "Mexican",
  "rating": 4.0,
  "price_range": "$",
  "opening_time": "10:00",
  "closing_time": "21:00",
  "coordinates": [59.3293, 18.0686]
}
```

#### `PUT /api/restaurants/{id}`
Updates an existing restaurant.

#### `DELETE /api/restaurants/{id}`
Soft deletes a restaurant (sets is_active to false).

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

#### `POST /api/menu-items`
Creates a new menu item.

#### `PUT /api/menu-items/{id}`
Updates an existing menu item.

#### `DELETE /api/menu-items/{id}`
Soft deletes a menu item.

## Project Structure

```
lunch-menu-api/
├── main.go                 # Application entry point
├── go.mod                  # Go module definition and dependencies
├── go.sum                  # Go module checksums
├── docker-compose.yml      # Docker setup with PostgreSQL
├── Dockerfile             # Docker image definition
├── internal/
│   ├── handlers/          # HTTP request handlers (Gin)
│   │   └── handlers.go
│   ├── models/            # Database models (GORM)
│   │   └── models.go
│   ├── database/          # Database connection and CRUD
│   │   └── database.go
│   └── utils/             # Business logic services
│       └── utils.go
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

### Local Development

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd lunch-menu
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Set up PostgreSQL:**
   ```bash
   # Option 1: Using Docker
   docker-compose up -d postgres
   
   # Option 2: Local PostgreSQL
   createdb lunch_menu
   ```

4. **Set environment variables:**
   ```bash
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_PASSWORD=postgres
   export DB_NAME=lunch_menu
   export DB_SSLMODE=disable
   ```

5. **Run the application:**
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:8080`.

### Building

```bash
# Build binary
go build -o lunch-menu-api .

# Run binary
./lunch-menu-api
```

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

### Option 0: Docker Compose (Automatic)

**Recommended for development and testing:**

```bash
docker-compose up -d
```

This automatically:
- Creates PostgreSQL database
- Loads all 18 original restaurants
- Starts the API with data ready
- No manual steps required!

### Option 1: SQL Seed Script

Use the provided SQL script with original restaurant data:

```bash
# Using the provided script (recommended)
./seed_database.sh

# Or manually with psql
psql -d lunch_menu -f seed_data.sql
```

This loads 18 restaurants from the original lunch-menu dataset (Uppsala and Solna locations) plus sample menu items.

### Option 2: REST API

Create restaurants and menu items through the API endpoints:

```bash
# Create a restaurant
curl -X POST http://localhost:8080/api/restaurants \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My Restaurant",
    "description": "Great food",
    "address": "123 Main St",
    "coordinate": [59.3293, 18.0686],
    "homepage": "https://myrestaurant.com",
    "region": "Stockholm",
    "phone": "+46 8 123 456",
    "email": "info@myrestaurant.com"
  }'

# Create menu items
curl -X POST http://localhost:8080/api/menu-items \
  -H "Content-Type: application/json" \
  -d '{
    "restaurant_id": 1,
    "name": "Today Special",
    "description": "Delicious daily special",
    "price": 125.00,
    "category": "Main Course"
  }'
```

### Option 3: CSV Import

Use PostgreSQL's COPY command with the provided CSV files:

```bash
# Copy CSV files to database
psql -d lunch_menu -c "\COPY restaurants(name,description,address,coordinate,homepage,region,phone,email,is_active) FROM 'restaurants.csv' WITH CSV HEADER;"
psql -d lunch_menu -c "\COPY menu_items(restaurant_id,name,description,price,category,is_available) FROM 'menu_items.csv' WITH CSV HEADER;"
```

## Database Migrations

Database migrations are handled automatically on application startup using GORM's AutoMigrate feature. This will:

- Create tables if they don't exist
- Add new columns when models are updated
- Create indexes defined in model tags

## API Testing

### Using curl

```bash
# Get all restaurants
curl http://localhost:8080/api/restaurants

# Get specific restaurant
curl http://localhost:8080/api/restaurants/1

# Create new restaurant
curl -X POST http://localhost:8080/api/restaurants \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Restaurant",
    "address": "123 Test St",
    "cuisine": "Italian",
    "rating": 4.5
  }'

# Get menu items for restaurant
curl http://localhost:8080/api/restaurants/1/menu-items
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
│       └── utils.go
└── README.md              # This file
```

## Dependencies

- **gin-gonic/gin** - HTTP web framework
- **gin-contrib/cors** - CORS middleware
- **PuerkitoBio/goquery** - HTML parsing (for menu scraping)
- **gorm.io/gorm** - ORM for database operations
- **gorm.io/driver/postgres** - PostgreSQL driver for GORM

## Prerequisites

- Go 1.23 or higher
- PostgreSQL 12+ (or use Docker Compose)

## Running the Application

### Quick Start with Docker Compose

The easiest way to run the application is using Docker Compose:

```bash
# Start PostgreSQL and the API
docker-compose up -d

# Check logs
docker-compose logs -f api
```

This will:
1. Start a PostgreSQL database
2. Run database migrations
3. Seed the database with restaurant data
4. Start the API server on port 8000

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

### Development
```bash
# Clone the repository and navigate to the Go backend
cd backend-go

# Install dependencies
go mod tidy

# Set up environment variables (see above)
# Start PostgreSQL database

# Run the application
go run .
```

The server will start on port 8000 by default. The application will:
1. Connect to PostgreSQL
2. Run database migrations automatically
3. Load restaurant data from restaurants.json and seed the database
4. Start the HTTP server

### Production Build
```bash
# Build the binary
go build -o lunch-menu-api .

# Run the binary
./lunch-menu-api
```

### Using Docker
```bash
# Build the Docker image
docker build -t lunch-menu-api .

# Run the container
docker run -p 8000:8000 lunch-menu-api
```

## Configuration

The application can be configured using environment variables:

**Server Configuration:**
- `PORT` - Server port (default: 8000)
- `VERSION` - Application version (default: "development")
- `GIN_MODE` - Gin mode: debug, release, or test (default: debug)

**Database Configuration:**
- `DB_HOST` - PostgreSQL host (default: localhost)
- `DB_PORT` - PostgreSQL port (default: 5432)
- `DB_USER` - Database username (default: postgres)
- `DB_PASSWORD` - Database password (default: postgres)
- `DB_NAME` - Database name (default: lunch_menu)
- `DB_SSLMODE` - SSL mode (default: disable)

## Testing the API

You can test the API endpoints using curl:

```bash
# Get API info
curl http://localhost:8000/api

# List all restaurants
curl http://localhost:8000/api/restaurant

# Get a specific restaurant
curl http://localhost:8000/api/restaurant/glada

# Get version
curl http://localhost:8000/api/version
```

## How It Works

1. **Startup**: The application connects to PostgreSQL and runs migrations
2. **Data Loading**: Restaurant data is loaded from restaurants.json and stored in the database
3. **Menu Parsing**: When a restaurant is requested:
   - Check if fresh menu exists in database
   - If not, parse the restaurant's website for current menu
   - Store the parsed menu in the database for future requests
4. **API Responses**: Serve data from the database with real-time menu parsing as needed

## Menu Parsing Implementation Status

Currently, the menu parsing logic contains stub implementations for all restaurant parsers. Each parser returns a placeholder message indicating that menu parsing is not yet implemented for that specific restaurant.

The database is ready to store parsed menus, and the infrastructure supports:
- Automatic menu refresh based on age
- Caching of parsed results
- Fallback to database when parsing fails

To fully implement menu parsing, you would need to:

1. Examine the original Python parsers in `parser.py`
2. Implement the specific parsing logic for each restaurant in the parsers package
3. Handle HTML parsing, date filtering, and text cleanup
4. Add error handling for network requests and parsing failures

## Differences from Original

This Go implementation maintains API compatibility with the original Python backend but has some architectural differences:

1. **No FastAPI cache decorator** - Instead uses cache headers
2. **Structured project layout** - Separates concerns into packages
3. **Stub menu parsers** - Parsing logic needs to be implemented
4. **Static binary** - Can be deployed without runtime dependencies

## Development Notes

- The application loads restaurant data from `restaurants.json` on startup
- CORS is configured to allow all origins for development
- The parser map is initialized with stub functions that can be replaced with actual implementations
- Error handling follows Go idioms with explicit error returns
- JSON responses match the original API structure exactly