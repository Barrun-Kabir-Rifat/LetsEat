package main

import (
	"log"
	"net/http"
	//"crypto/subtle"
    "os"

    
	"github.com/Barrun-Kabir-Rifat/letseat/internal/database"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/handlers"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/repositories"
	"github.com/Barrun-Kabir-Rifat/letseat/internal/services"
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
	
  "github.com/Barrun-Kabir-Rifat/letseat/internal/middleware"

)


func main() {



	db, err := database.Connect()

    	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not configured")
	}

	adminRepo := repositories.NewAdminRepository(db)

	adminService := services.NewAdminService(
    	adminRepo,
    	os.Getenv("JWT_SECRET"),
	)

	adminHandler := handlers.NewAdminHandler(adminService)



	// Public login


	e := echo.New()

	e.POST("/api/v1/admin/login", adminHandler.Login)

	// Protected registration
    e.POST(
		"/api/v1/admin/register",
		adminHandler.Register,
		middleware.JWTAuth(os.Getenv("JWT_SECRET")),
    )

	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer db.Close()
	//basicAuthMiddleware(e)


	restaurantRepository := repositories.NewRestaurantRepository(db)
	restaurantService := services.NewRestaurantService(restaurantRepository)
	restaurantHandler := handlers.NewRestaurantHandler(restaurantService)

	e.POST("/api/v1/restaurants", restaurantHandler.Create, middleware.JWTAuth(jwtSecret),)
	e.GET("/api/v1/restaurants", restaurantHandler.GetAll)
	e.GET("/api/v1/restaurants/:id", restaurantHandler.GetByID)
	e.PUT("/api/v1/restaurants/:id", restaurantHandler.Update, middleware.JWTAuth(jwtSecret),)
	e.DELETE("/api/v1/restaurants/:id", restaurantHandler.Delete, middleware.JWTAuth(jwtSecret),)

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	e.POST("/api/v1/categories", categoryHandler.Create, middleware.JWTAuth(jwtSecret),)
	e.GET("/api/v1/categories", categoryHandler.GetAll)
	e.GET("/api/v1/categories/:id", categoryHandler.GetByID)
	e.PUT("/api/v1/categories/:id", categoryHandler.Update, middleware.JWTAuth(jwtSecret),)
	e.DELETE("/api/v1/categories/:id", categoryHandler.Delete, middleware.JWTAuth(jwtSecret),)

	menuItemRepository := repositories.NewMenuItemRepository(db)
	menuItemService := services.NewMenuItemService(menuItemRepository)
	menuItemHandler := handlers.NewMenuItemHandler(menuItemService)

	e.POST("/api/v1/restaurants/:id/menu-items", menuItemHandler.Create, middleware.JWTAuth(jwtSecret),)
	e.GET("/api/v1/restaurants/:id/menu-items", menuItemHandler.GetAllByRestaurantID)

	e.GET("/api/v1/menu-items/:id", menuItemHandler.GetByID)

	e.PUT("/api/v1/menu-items/:id", menuItemHandler.Update, middleware.JWTAuth(jwtSecret),)

	e.PATCH("/api/v1/menu-items/:id/price", menuItemHandler.UpdatePrice, middleware.JWTAuth(jwtSecret),)

	e.PATCH("/api/v1/menu-items/:id/availability", menuItemHandler.UpdateAvailability, middleware.JWTAuth(jwtSecret),)

	e.DELETE("/api/v1/menu-items/:id", menuItemHandler.Delete, middleware.JWTAuth(jwtSecret),)

	searchRepository := repositories.NewSearchRepository(db)
	searchService := services.NewSearchService(searchRepository)
	searchHandler := handlers.NewSearchHandler(searchService)

	e.GET("/api/v1/restaurants/search", searchHandler.SearchRestaurants)

	e.GET("/api/v1/menu-items/search", searchHandler.SearchMenuItems)

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
