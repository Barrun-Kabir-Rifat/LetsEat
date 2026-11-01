package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/services"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	Service *services.CategoryService
}

func NewCategoryHandler(
	service *services.CategoryService,
) *CategoryHandler {
	return &CategoryHandler{Service: service}
}

func (h *CategoryHandler) Create(c echo.Context) error {
	var category struct {
		Name string `json:"name"`
	}

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	err := h.Service.Create(c.Request().Context(), &models.Category{
		Name: category.Name,
	})
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return err
		}

		if err.Error() == "category name is required" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to create category",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "category created successfully",
		"data":    category,
	})
}
