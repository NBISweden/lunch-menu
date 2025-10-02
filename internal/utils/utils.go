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

// CreateRestaurant creates a new restaurant
func (s *RestaurantService) CreateRestaurant(req *models.CreateRestaurantRequest) (*models.Restaurant, error) {
	restaurant := &models.Restaurant{
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		Coordinate:  req.Coordinate,
		Homepage:    req.Homepage,
		Region:      req.Region,
		Phone:       req.Phone,
		Email:       req.Email,
		IsActive:    true,
	}

	if err := database.CreateRestaurant(restaurant); err != nil {
		return nil, err
	}

	return restaurant, nil
}

// UpdateRestaurant updates an existing restaurant
func (s *RestaurantService) UpdateRestaurant(id uint, req *models.UpdateRestaurantRequest) (*models.Restaurant, error) {
	restaurant, err := database.GetRestaurantByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Name != nil {
		restaurant.Name = *req.Name
	}
	if req.Description != nil {
		restaurant.Description = *req.Description
	}
	if req.Address != nil {
		restaurant.Address = *req.Address
	}
	if req.Coordinate != nil {
		restaurant.Coordinate = *req.Coordinate
	}
	if req.Homepage != nil {
		restaurant.Homepage = *req.Homepage
	}
	if req.Region != nil {
		restaurant.Region = *req.Region
	}
	if req.Phone != nil {
		restaurant.Phone = *req.Phone
	}
	if req.Email != nil {
		restaurant.Email = *req.Email
	}
	if req.IsActive != nil {
		restaurant.IsActive = *req.IsActive
	}

	if err := database.UpdateRestaurant(restaurant); err != nil {
		return nil, err
	}

	return restaurant, nil
}

// DeleteRestaurant deletes a restaurant
func (s *RestaurantService) DeleteRestaurant(id uint) error {
	return database.DeleteRestaurant(id)
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

// CreateMenuItem creates a new menu item
func (s *MenuItemService) CreateMenuItem(req *models.CreateMenuItemRequest) (*models.MenuItem, error) {
	// Verify restaurant exists
	_, err := database.GetRestaurantByID(req.RestaurantID)
	if err != nil {
		return nil, err
	}

	menuItem := &models.MenuItem{
		RestaurantID: req.RestaurantID,
		Name:         req.Name,
		Description:  req.Description,
		Price:        req.Price,
		Category:     req.Category,
		IsAvailable:  true,
	}

	if err := database.CreateMenuItem(menuItem); err != nil {
		return nil, err
	}

	return menuItem, nil
}

// UpdateMenuItem updates an existing menu item
func (s *MenuItemService) UpdateMenuItem(id uint, req *models.UpdateMenuItemRequest) (*models.MenuItem, error) {
	menuItem, err := database.GetMenuItemByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Name != nil {
		menuItem.Name = *req.Name
	}
	if req.Description != nil {
		menuItem.Description = *req.Description
	}
	if req.Price != nil {
		menuItem.Price = *req.Price
	}
	if req.Category != nil {
		menuItem.Category = *req.Category
	}
	if req.IsAvailable != nil {
		menuItem.IsAvailable = *req.IsAvailable
	}

	if err := database.UpdateMenuItem(menuItem); err != nil {
		return nil, err
	}

	return menuItem, nil
}

// DeleteMenuItem deletes a menu item
func (s *MenuItemService) DeleteMenuItem(id uint) error {
	return database.DeleteMenuItem(id)
}
