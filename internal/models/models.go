package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// FloatArray represents a PostgreSQL array of floats for coordinates
type FloatArray []float64

// Value implements the driver.Valuer interface
func (fa FloatArray) Value() (driver.Value, error) {
	return json.Marshal(fa)
}

// Scan implements the sql.Scanner interface
func (fa *FloatArray) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, fa)
	case string:
		return json.Unmarshal([]byte(v), fa)
	default:
		return errors.New("cannot scan into FloatArray")
	}
}

// Restaurant represents a restaurant
type Restaurant struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Address     string     `json:"address"`
	Coordinate  FloatArray `json:"coordinate"`
	Homepage    string     `json:"homepage"`
	Region      string     `json:"region"`
	Phone       string     `json:"phone"`
	Email       string     `json:"email"`
	IsActive    bool       `json:"is_active"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// MenuItem represents a menu item
type MenuItem struct {
	ID           uint      `json:"id"`
	RestaurantID uint      `json:"restaurant_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	Category     string    `json:"category"`
	IsAvailable  bool      `json:"is_available"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Response Models

// RestaurantsResponse represents the response for listing restaurants
type RestaurantsResponse struct {
	Restaurants []Restaurant `json:"restaurants"`
	Total       int64        `json:"total"`
}

// RestaurantResponse represents the response for a single restaurant
type RestaurantResponse struct {
	Restaurant Restaurant `json:"restaurant"`
}

// MenuItemsResponse represents the response for restaurant menu items
type MenuItemsResponse struct {
	MenuItems []MenuItem `json:"menu_items"`
	Total     int64      `json:"total"`
}

// APIResponse represents basic API information
type APIResponse struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// BusinessStatistics represents business analytics data
type BusinessStatistics struct {
	TotalRestaurants    int64                    `json:"total_restaurants"`
	ActiveRestaurants   int64                    `json:"active_restaurants"`
	InactiveRestaurants int64                    `json:"inactive_restaurants"`
	TotalMenuItems      int64                    `json:"total_menu_items"`
	AveragePrice        float64                  `json:"average_price"`
	RevenueByCateory    map[string]float64       `json:"revenue_by_category"`
	RestaurantDetails   []RestaurantBusinessData `json:"restaurant_details"`
}

// RestaurantBusinessData represents per-restaurant business data
type RestaurantBusinessData struct {
	RestaurantID   uint    `json:"restaurant_id"`
	RestaurantName string  `json:"restaurant_name"`
	MenuItemCount  int64   `json:"menu_item_count"`
	AveragePrice   float64 `json:"average_price"`
	TotalRevenue   float64 `json:"total_revenue"`
}
