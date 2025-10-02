package utils

import (
	"lunch-menu-api/internal/database"
	"lunch-menu-api/internal/models"
)

// RestaurantService provides business logic for restaurant operations
type RestaurantService struct{}

// NewRestaurantService creates a new restaurant service
func NewRestaurantService() *RestaurantService {
	return &RestaurantService{}
}

// GetRestaurants retrieves all restaurants with pagination
func (s *RestaurantService) GetRestaurants(limit, offset int) ([]models.Restaurant, int64, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	return database.GetRestaurants(limit, offset)
}

// GetRestaurant retrieves a restaurant by ID
func (s *RestaurantService) GetRestaurant(id uint) (*models.Restaurant, error) {
	return database.GetRestaurantByID(id)
}

// MenuItemService provides business logic for menu item operations
type MenuItemService struct{}

// NewMenuItemService creates a new menu item service
func NewMenuItemService() *MenuItemService {
	return &MenuItemService{}
}

// GetMenuItems retrieves menu items for a restaurant with pagination
func (s *MenuItemService) GetMenuItems(restaurantID uint, limit, offset int) ([]models.MenuItem, int64, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	return database.GetMenuItems(restaurantID, limit, offset)
}

// GetMenuItem retrieves a menu item by ID
func (s *MenuItemService) GetMenuItem(id uint) (*models.MenuItem, error) {
	return database.GetMenuItemByID(id)
}
