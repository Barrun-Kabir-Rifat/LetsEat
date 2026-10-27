package main

import (
	"log"
	"net/http"

	"github.com/Barrun-Kabir-Rifat/letseat/internal/database"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/handlers"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/repositories"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/services"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer db.Close()

	e := echo.New()

	restaurantRepository := repositories.NewRestaurantRepository(db)
	restaurantService := services.NewRestaurantService(restaurantRepository)
	restaurantHandler := handlers.NewRestaurantHandler(restaurantService)

	e.POST("/api/v1/restaurants", restaurantHandler.Create)
	e.GET("/api/v1/restaurants", restaurantHandler.GetAll)
	e.GET("/api/v1/restaurants/:id", restaurantHandler.GetByID)
	e.PUT("/api/v1/restaurants/:id", restaurantHandler.Update)
	e.DELETE("/api/v1/restaurants/:id", restaurantHandler.Delete)

	e.GET("/api/v1/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "success",
			"message": "LetsEat API is running.Great",
		})
	})

	if err := e.Start(":8080"); err != nil &&
		err != http.ErrServerClosed {
		log.Fatal("API server failed: ", err)
	}
}
