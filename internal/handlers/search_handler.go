package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/services"
	"github.com/labstack/echo/v4"
)

type SearchHandler struct {
	Service *services.SearchService
}

func NewSearchHandler(
	service *services.SearchService,
) *SearchHandler {
	return &SearchHandler{
		Service: service,
	}
}

// SearchRestaurants searches restaurants by name.
func (h *SearchHandler) SearchRestaurants(c echo.Context) error {
	name := strings.TrimSpace(c.QueryParam("name"))

	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "restaurant search name is required",
		})
	}

	restaurants, err := h.Service.SearchRestaurants(
		c.Request().Context(),
		name,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to search restaurants",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": restaurants,
	})
}

// SearchMenuItems searches food using optional filters.
func (h *SearchHandler) SearchMenuItems(c echo.Context) error {
	name := strings.TrimSpace(c.QueryParam("name"))

	var categoryID int64
	if value := c.QueryParam("category_id"); value != "" {
		parsedID, err := strconv.ParseInt(value, 10, 64)
		if err != nil || parsedID <= 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid category ID",
			})
		}
		categoryID = parsedID
	}

	var restaurantID int64
	if value := c.QueryParam("restaurant_id"); value != "" {
		parsedID, err := strconv.ParseInt(value, 10, 64)
		if err != nil || parsedID <= 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid restaurant ID",
			})
		}
		restaurantID = parsedID
	}

	// Available food is returned by default.
	available := true

	if value := c.QueryParam("available"); value != "" {
		parsedValue, err := strconv.ParseBool(value)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "available must be true or false",
			})
		}
		available = parsedValue
	}

	menuItems, err := h.Service.SearchMenuItems(
		c.Request().Context(),
		name,
		categoryID,
		restaurantID,
		available,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to search food items",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuItems,
	})
}
