package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/models"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/services"
	"github.com/labstack/echo/v4"

	"github.com/jackc/pgx/v5/pgconn"
)

func isCategoryInUse(err error) bool {
    var pgErr *pgconn.PgError

    return errors.As(err, &pgErr) &&
        pgErr.Code == "23503"
}

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

func (h *CategoryHandler) GetAll(c echo.Context) error {
	categories, err := h.Service.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to retrieve categories",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": categories,
	})
}

func (h *CategoryHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid category ID",
		})
	}

	category, err := h.Service.GetByID(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrCategoryNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "category not found",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to retrieve category",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": category,
	})
}



func (h *CategoryHandler) Update(c echo.Context) error {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil || id <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "invalid category ID",
        })
    }

    var request struct {
        Name string `json:"name"`
    }

    if err := c.Bind(&request); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "invalid request body",
        })
    }

    category := &models.Category{
        ID:   id,
        Name: request.Name,
    }

    err = h.Service.Update(c.Request().Context(), category)
    if err != nil {
        if errors.Is(err, services.ErrCategoryNotFound) {
            return c.JSON(http.StatusNotFound, map[string]string{
                "error": "category not found",
            })
        }

        if err.Error() == "category name is required" {
            return c.JSON(http.StatusBadRequest, map[string]string{
                "error": err.Error(),
            })
        }

        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "failed to update category",
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "category updated successfully",
        "data":    category,
    })
}



func (h *CategoryHandler) Delete(c echo.Context) error {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil || id <= 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "invalid category ID",
        })
    }

    err = h.Service.Delete(c.Request().Context(), id)
    if err != nil {
        if errors.Is(err, services.ErrCategoryNotFound) {
            return c.JSON(http.StatusNotFound, map[string]string{
                "error": "category not found",
            })
        }

        if isCategoryInUse(err) {
            return c.JSON(http.StatusConflict, map[string]string{
                "error": "category cannot be deleted because menu items are using it",
            })
        }

        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "failed to delete category",
        })
    }

    return c.NoContent(http.StatusNoContent)
}

