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
			restaurants.GET("/:id", GetRestaurant)

			// Menu items for a restaurant
			restaurants.GET("/:id/menu", GetRestaurantMenu)
		}

		// Menu item routes
		menuItems := api.Group("/menu-items")
		{
			menuItems.GET("/:id", GetMenuItem)
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
