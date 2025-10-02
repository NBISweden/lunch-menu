package database

import (
	"database/sql"
	"fmt"
	"log"
	"lunch-menu-api/internal/models"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDatabase() error {
	var err error

	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "lunch_menu")
	sslmode := getEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)

	log.Println("Database connection established successfully")
	return nil
}

func MigrateDatabase() error {
	// Since we're using SQL initialization scripts, no migration needed
	log.Println("Database migration skipped - using SQL initialization")
	return nil
}

func CloseDatabase() error {
	if DB == nil {
		return nil
	}
	return DB.Close()
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func GetRestaurants(limit, offset int) ([]models.Restaurant, int64, error) {
	var restaurants []models.Restaurant
	var total int64

	// Get total count
	err := DB.QueryRow("SELECT COUNT(*) FROM restaurants WHERE is_active = true").Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count restaurants: %w", err)
	}

	// Get restaurants with pagination
	query := `
		SELECT id, name, description, address, coordinate, homepage, region, phone, email, is_active, created_at, updated_at 
		FROM restaurants 
		WHERE is_active = true 
		ORDER BY id 
		LIMIT $1 OFFSET $2`

	rows, err := DB.Query(query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query restaurants: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var r models.Restaurant
		err := rows.Scan(&r.ID, &r.Name, &r.Description, &r.Address, &r.Coordinate,
			&r.Homepage, &r.Region, &r.Phone, &r.Email, &r.IsActive, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan restaurant: %w", err)
		}
		restaurants = append(restaurants, r)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating restaurants: %w", err)
	}

	return restaurants, total, nil
}

func GetRestaurantByID(id uint) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	query := `
		SELECT id, name, description, address, coordinate, homepage, region, phone, email, is_active, created_at, updated_at 
		FROM restaurants 
		WHERE id = $1 AND is_active = true`

	err := DB.QueryRow(query, id).Scan(&restaurant.ID, &restaurant.Name, &restaurant.Description,
		&restaurant.Address, &restaurant.Coordinate, &restaurant.Homepage, &restaurant.Region,
		&restaurant.Phone, &restaurant.Email, &restaurant.IsActive, &restaurant.CreatedAt, &restaurant.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("restaurant not found")
		}
		return nil, fmt.Errorf("failed to get restaurant: %w", err)
	}
	return &restaurant, nil
}

func CreateRestaurant(restaurant *models.Restaurant) error {
	query := `
		INSERT INTO restaurants (name, description, address, coordinate, homepage, region, phone, email, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, true, NOW(), NOW())
		RETURNING id, created_at, updated_at`

	err := DB.QueryRow(query, restaurant.Name, restaurant.Description, restaurant.Address,
		restaurant.Coordinate, restaurant.Homepage, restaurant.Region, restaurant.Phone, restaurant.Email).
		Scan(&restaurant.ID, &restaurant.CreatedAt, &restaurant.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create restaurant: %w", err)
	}
	restaurant.IsActive = true
	return nil
}

func UpdateRestaurant(restaurant *models.Restaurant) error {
	query := `
		UPDATE restaurants 
		SET name = $1, description = $2, address = $3, coordinate = $4, homepage = $5, 
		    region = $6, phone = $7, email = $8, updated_at = NOW()
		WHERE id = $9 AND is_active = true
		RETURNING updated_at`

	err := DB.QueryRow(query, restaurant.Name, restaurant.Description, restaurant.Address,
		restaurant.Coordinate, restaurant.Homepage, restaurant.Region, restaurant.Phone,
		restaurant.Email, restaurant.ID).Scan(&restaurant.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("restaurant not found or inactive")
		}
		return fmt.Errorf("failed to update restaurant: %w", err)
	}
	return nil
}

func DeleteRestaurant(id uint) error {
	query := "UPDATE restaurants SET is_active = false, updated_at = NOW() WHERE id = $1"
	result, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete restaurant: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("restaurant not found")
	}
	return nil
}

func GetMenuItems(restaurantID uint, limit, offset int) ([]models.MenuItem, int64, error) {
	var menuItems []models.MenuItem
	var total int64

	// Get total count
	err := DB.QueryRow("SELECT COUNT(*) FROM menu_items WHERE restaurant_id = $1 AND is_available = true",
		restaurantID).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count menu items: %w", err)
	}

	// Get menu items with pagination
	query := `
		SELECT id, restaurant_id, name, description, price, category, is_available, created_at, updated_at 
		FROM menu_items 
		WHERE restaurant_id = $1 AND is_available = true 
		ORDER BY id 
		LIMIT $2 OFFSET $3`

	rows, err := DB.Query(query, restaurantID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query menu items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.MenuItem
		err := rows.Scan(&item.ID, &item.RestaurantID, &item.Name, &item.Description,
			&item.Price, &item.Category, &item.IsAvailable, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan menu item: %w", err)
		}
		menuItems = append(menuItems, item)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating menu items: %w", err)
	}

	return menuItems, total, nil
}

func GetMenuItemByID(id uint) (*models.MenuItem, error) {
	var menuItem models.MenuItem
	query := `
		SELECT id, restaurant_id, name, description, price, category, is_available, created_at, updated_at 
		FROM menu_items 
		WHERE id = $1 AND is_available = true`

	err := DB.QueryRow(query, id).Scan(&menuItem.ID, &menuItem.RestaurantID, &menuItem.Name,
		&menuItem.Description, &menuItem.Price, &menuItem.Category, &menuItem.IsAvailable,
		&menuItem.CreatedAt, &menuItem.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("menu item not found")
		}
		return nil, fmt.Errorf("failed to get menu item: %w", err)
	}
	return &menuItem, nil
}

func CreateMenuItem(menuItem *models.MenuItem) error {
	query := `
		INSERT INTO menu_items (restaurant_id, name, description, price, category, is_available, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, true, NOW(), NOW())
		RETURNING id, created_at, updated_at`

	err := DB.QueryRow(query, menuItem.RestaurantID, menuItem.Name, menuItem.Description,
		menuItem.Price, menuItem.Category).Scan(&menuItem.ID, &menuItem.CreatedAt, &menuItem.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create menu item: %w", err)
	}
	menuItem.IsAvailable = true
	return nil
}

func UpdateMenuItem(menuItem *models.MenuItem) error {
	query := `
		UPDATE menu_items 
		SET name = $1, description = $2, price = $3, category = $4, updated_at = NOW()
		WHERE id = $5 AND is_available = true
		RETURNING updated_at`

	err := DB.QueryRow(query, menuItem.Name, menuItem.Description, menuItem.Price,
		menuItem.Category, menuItem.ID).Scan(&menuItem.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("menu item not found or inactive")
		}
		return fmt.Errorf("failed to update menu item: %w", err)
	}
	return nil
}

func DeleteMenuItem(id uint) error {
	query := "UPDATE menu_items SET is_available = false, updated_at = NOW() WHERE id = $1"
	result, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete menu item: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("menu item not found")
	}
	return nil
}
