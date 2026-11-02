package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/services"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"

	"github.com/shopspring/decimal"
)

type MenuItemHandler struct {
	Service *services.MenuItemService
}

func NewMenuItemHandler(
	service *services.MenuItemService,
) *MenuItemHandler {
	return &MenuItemHandler{
		Service: service,
	}
}

func (h *MenuItemHandler) Create(c echo.Context) error {
	restaurantID, err := strconv.ParseInt(
		c.Param("id"), 10, 64,
	)

	if err != nil || restaurantID <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid restaurant ID",
		})
	}

	var menuItem models.MenuItem

	if err := c.Bind(&menuItem); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	menuItem.RestaurantID = restaurantID

	if err := h.Service.Create(
		c.Request().Context(),
		&menuItem,
	); err != nil {
		if err.Error() == "invalid restaurant ID" ||
			err.Error() == "invalid category ID" ||
			err.Error() == "food item name is required" ||
			err.Error() == "food item price must be greater than zero" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to create food item",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "food item created successfully",
		"data":    menuItem,
	})
}

func (h *MenuItemHandler) GetAllByRestaurantID(
	c echo.Context,
) error {
	restaurantID, err := strconv.ParseInt(
		c.Param("id"), 10, 64,
	)

	if err != nil || restaurantID <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid restaurant ID",
		})
	}

	menuItems, err := h.Service.GetAllByRestaurantID(
		c.Request().Context(),
		restaurantID,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to retrieve menu items",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuItems,
	})
}

func menuItemError(c echo.Context, err error) error {
	if errors.Is(err, services.ErrMenuItemNotFound) {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "food item not found",
		})
	}

	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		if pgErr.Code == "23503" {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "referenced restaurant or category not found",
			})
		}

		if pgErr.Code == "23514" ||
			pgErr.Code == "22003" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid food item data",
			})
		}
	}

	switch err.Error() {
	case "invalid food item ID",
		"invalid category ID",
		"food item name is required",
		"food item price must be greater than zero":
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(
		http.StatusInternalServerError,
		map[string]string{
			"error": "food item operation failed",
		},
	)
}

func (h *MenuItemHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid food item ID",
		})
	}

	item, err := h.Service.GetByID(
		c.Request().Context(), id,
	)
	if err != nil {
		return menuItemError(c, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": item,
	})
}

func (h *MenuItemHandler) Update(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid food item ID",
		})
	}

	var request struct {
		CategoryID  int64           `json:"category_id"`
		Name        string          `json:"name"`
		Description *string         `json:"description"`
		Price       decimal.Decimal `json:"price"`
		IsAvailable bool            `json:"is_available"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	item := &models.MenuItem{
		ID:          id,
		CategoryID:  request.CategoryID,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		IsAvailable: request.IsAvailable,
	}

	if err := h.Service.Update(
		c.Request().Context(), item,
	); err != nil {
		return menuItemError(c, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "food item updated successfully",
		"data":    item,
	})
}

func (h *MenuItemHandler) UpdatePrice(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid food item ID",
		})
	}

	var request struct {
		Price *decimal.Decimal `json:"price"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	if request.Price == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "price is required",
		})
	}

	if err := h.Service.UpdatePrice(
		c.Request().Context(),
		id,
		*request.Price,
	); err != nil {
		return menuItemError(c, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "food price updated successfully",
		"price":   request.Price,
	})
}

func (h *MenuItemHandler) UpdateAvailability(
	c echo.Context,
) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid food item ID",
		})
	}

	var request struct {
		IsAvailable *bool `json:"is_available"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	if request.IsAvailable == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "is_available is required",
		})
	}

	if err := h.Service.UpdateAvailability(
		c.Request().Context(),
		id,
		*request.IsAvailable,
	); err != nil {
		return menuItemError(c, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "food availability updated successfully",
		"is_available": *request.IsAvailable,
	})
}

func (h *MenuItemHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid food item ID",
		})
	}

	if err := h.Service.Delete(
		c.Request().Context(), id,
	); err != nil {
		return menuItemError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}
