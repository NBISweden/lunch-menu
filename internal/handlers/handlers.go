package handlers

import (
	"lunch-menu-api/internal/models"
	"lunch-menu-api/internal/utils"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	restaurantService = utils.NewRestaurantService()
	menuItemService   = utils.NewMenuItemService()
)

// SetupRoutes sets up all API routes
func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Basic API info
		api.GET("", GetAPIInfo)

		// Restaurant routes
		restaurants := api.Group("/restaurants")
		{
			restaurants.GET("", GetRestaurants)
			restaurants.POST("", CreateRestaurant)
			restaurants.GET("/:id", GetRestaurant)
			restaurants.PUT("/:id", UpdateRestaurant)
			restaurants.DELETE("/:id", DeleteRestaurant)

			// Menu items for a restaurant
			restaurants.GET("/:id/menu", GetRestaurantMenu)
		}

		// Menu item routes
		menuItems := api.Group("/menu-items")
		{
			menuItems.POST("", CreateMenuItem)
			menuItems.GET("/:id", GetMenuItem)
			menuItems.PUT("/:id", UpdateMenuItem)
			menuItems.DELETE("/:id", DeleteMenuItem)
		}
	}
}

// GetAPIInfo returns basic API information
func GetAPIInfo(c *gin.Context) {
	version := os.Getenv("VERSION")
	if version == "" {
		version = "development"
	}

	response := models.APIResponse{
		Message: "Restaurant Management API",
		Version: version,
	}
	c.JSON(http.StatusOK, response)
}

// Restaurant handlers

// GetRestaurants returns all restaurants with pagination
func GetRestaurants(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	restaurants, total, err := restaurantService.GetRestaurants(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "DATABASE_ERROR",
			Message: "Failed to retrieve restaurants",
		})
		return
	}

	response := models.RestaurantsResponse{
		Restaurants: restaurants,
		Total:       total,
	}
	c.JSON(http.StatusOK, response)
}

// GetRestaurant returns a specific restaurant
func GetRestaurant(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "Invalid restaurant ID",
		})
		return
	}

	restaurant, err := restaurantService.GetRestaurant(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "NOT_FOUND",
			Message: "Restaurant not found",
		})
		return
	}

	response := models.RestaurantResponse{
		Restaurant: *restaurant,
	}
	c.JSON(http.StatusOK, response)
}

// CreateRestaurant creates a new restaurant
func CreateRestaurant(c *gin.Context) {
	var req models.CreateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
		})
		return
	}

	restaurant, err := restaurantService.CreateRestaurant(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "DATABASE_ERROR",
			Message: "Failed to create restaurant",
		})
		return
	}

	response := models.RestaurantResponse{
		Restaurant: *restaurant,
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateRestaurant updates an existing restaurant
func UpdateRestaurant(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "Invalid restaurant ID",
		})
		return
	}

	var req models.UpdateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
		})
		return
	}

	restaurant, err := restaurantService.UpdateRestaurant(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "NOT_FOUND",
			Message: "Restaurant not found",
		})
		return
	}

	response := models.RestaurantResponse{
		Restaurant: *restaurant,
	}
	c.JSON(http.StatusOK, response)
}

// DeleteRestaurant deletes a restaurant
func DeleteRestaurant(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "Invalid restaurant ID",
		})
		return
	}

	if err := restaurantService.DeleteRestaurant(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "NOT_FOUND",
			Message: "Restaurant not found",
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetRestaurantMenu returns menu items for a restaurant
func GetRestaurantMenu(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "Invalid restaurant ID",
		})
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	menuItems, total, err := menuItemService.GetMenuItems(uint(id), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "DATABASE_ERROR",
			Message: "Failed to retrieve menu items",
		})
		return
	}

	response := models.MenuItemsResponse{
		MenuItems: menuItems,
		Total:     total,
	}
	c.JSON(http.StatusOK, response)
}

// Menu item handlers

// GetMenuItem returns a specific menu item
func GetMenuItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "Invalid menu item ID",
		})
		return
	}

	menuItem, err := menuItemService.GetMenuItem(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "NOT_FOUND",
			Message: "Menu item not found",
		})
		return
	}

	c.JSON(http.StatusOK, menuItem)
}

// CreateMenuItem creates a new menu item
func CreateMenuItem(c *gin.Context) {
	var req models.CreateMenuItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
		})
		return
	}

	menuItem, err := menuItemService.CreateMenuItem(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "DATABASE_ERROR",
			Message: "Failed to create menu item",
		})
		return
	}

	c.JSON(http.StatusCreated, menuItem)
}

// UpdateMenuItem updates an existing menu item
func UpdateMenuItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "Invalid menu item ID",
		})
		return
	}

	var req models.UpdateMenuItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "VALIDATION_ERROR",
			Message: err.Error(),
		})
		return
	}

	menuItem, err := menuItemService.UpdateMenuItem(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "NOT_FOUND",
			Message: "Menu item not found",
		})
		return
	}

	c.JSON(http.StatusOK, menuItem)
}

// DeleteMenuItem deletes a menu item
func DeleteMenuItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "INVALID_ID",
			Message: "Invalid menu item ID",
		})
		return
	}

	if err := menuItemService.DeleteMenuItem(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error:   "NOT_FOUND",
			Message: "Menu item not found",
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
