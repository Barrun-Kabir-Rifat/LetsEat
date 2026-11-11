package handlers

import (
	"net/http"
	"time"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/services"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	Service *services.AdminService
}

func NewAdminHandler(
	service *services.AdminService,
) *AdminHandler {
	return &AdminHandler{Service: service}
}

type AdminRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AdminHandler) Register(c echo.Context) error {

	var req AdminRegisterRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"Invalid request body",
		)
	}

	admin, err := h.Service.Register(
		c.Request().Context(),
		req.Username,
		req.Password,
	)

	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Admin registered successfully",
		"admin":   admin,
	})
}

func (h *AdminHandler) Login(c echo.Context) error {

	var req AdminLoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"Invalid request body",
		)
	}

	token, err := h.Service.Login(
		c.Request().Context(),
		req.Username,
		req.Password,
	)

	if err != nil {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			"Invalid username or password",
		)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	})
}

func (h *AdminHandler) Logout(c echo.Context) error {

	jti, ok := c.Get("jwt_id").(string)

	if !ok || jti == "" {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			"Invalid authentication",
		)
	}

	expiresAt, ok := c.Get("jwt_expires_at").(time.Time)

	if !ok {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			"Invalid token expiration",
		)
	}

	err := h.Service.Logout(
		c.Request().Context(),
		jti,
		expiresAt,
	)

	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Could not logout",
		)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Logout successful",
	})
}
