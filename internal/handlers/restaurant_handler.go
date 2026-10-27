package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/services"
	"github.com/labstack/echo/v4"
)

type RestaurantHandler struct {
	Service *services.RestaurantService
}

func NewRestaurantHandler(
	service *services.RestaurantService,
) *RestaurantHandler {
	return &RestaurantHandler{Service: service}
}

func (h *RestaurantHandler) Create(c echo.Context) error {
	var restaurant models.Restaurant

	if err := c.Bind(&restaurant); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	if err := h.Service.Create(c.Request().Context(), &restaurant); err != nil {
		if errors.Is(err, context.Canceled) {
			return err
		}

		if err.Error() == "restaurant name is required" ||
			err.Error() == "restaurant address is required" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to create restaurant",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "restaurant created successfully",
		"data":    restaurant,
	})
}

func (h *RestaurantHandler) GetAll(c echo.Context) error {
	restaurants, err := h.Service.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to retrieve restaurants",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": restaurants,
	})
}

func (h *RestaurantHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid restaurant ID",
		})
	}

	restaurant, err := h.Service.GetByID(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrRestaurantNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "restaurant not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to retrieve restaurant",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": restaurant,
	})
}

func (h *RestaurantHandler) Update(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid restaurant ID",
		})
	}

	var restaurant models.Restaurant

	if err := c.Bind(&restaurant); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	restaurant.ID = id

	if err := h.Service.Update(c.Request().Context(), &restaurant); err != nil {
		if errors.Is(err, services.ErrRestaurantNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "restaurant not found",
			})
		}

		if err.Error() == "restaurant name is required" ||
			err.Error() == "restaurant address is required" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to update restaurant",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "restaurant updated successfully",
		"data":    restaurant,
	})
}

func (h *RestaurantHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid restaurant ID",
		})
	}

	err = h.Service.Delete(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrRestaurantNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "restaurant not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to delete restaurant",
		})
	}

	return c.NoContent(http.StatusNoContent)
}
