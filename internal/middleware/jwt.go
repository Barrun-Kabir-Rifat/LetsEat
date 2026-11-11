package middleware

import (
	"net/http"
	"strings"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/repositories"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/services"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTAuth(
	secret string,
	adminRepo *repositories.AdminRepository,
) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			authHeader := c.Request().Header.Get("Authorization")
			parts := strings.SplitN(authHeader, " ", 2)

			if len(parts) != 2 ||
				!strings.EqualFold(parts[0], "Bearer") {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Missing or invalid authorization header",
				)
			}

			token, err := jwt.ParseWithClaims(
				parts[1],
				&services.AdminClaims{},
				func(token *jwt.Token) (interface{}, error) {

					if token.Method != jwt.SigningMethodHS256 {
						return nil, jwt.ErrSignatureInvalid
					}

					return []byte(secret), nil
				},
				jwt.WithIssuer("letseat-api"),
			)

			if err != nil || token == nil || !token.Valid {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Invalid or expired token",
				)
			}

			claims, ok := token.Claims.(*services.AdminClaims)

			if !ok ||
				claims.AdminID <= 0 ||
				claims.ID == "" ||
				claims.ExpiresAt == nil {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Invalid token claims",
				)
			}

			revoked, err := adminRepo.IsTokenRevoked(
				c.Request().Context(),
				claims.ID,
			)

			if err != nil {
				return echo.NewHTTPError(
					http.StatusInternalServerError,
					"Could not verify token status",
				)
			}

			if revoked {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Token has been revoked. Please login again.",
				)
			}

			c.Set("admin_id", claims.AdminID)
			c.Set("admin_username", claims.Username)
			c.Set("jwt_id", claims.ID)
			c.Set("jwt_expires_at", claims.ExpiresAt.Time)

			return next(c)
		}
	}
}
