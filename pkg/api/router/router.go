package router

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/api/handler"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/services"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router pkg
func SetupRoutes(app *fiber.App) {

	userService := services.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	// Api base
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Health
	health := v1.Group("/health")
	health.Get("/", handler.HealthCheck)

	// Auth
	auth := v1.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	user := v1.Group("/users")
	user.Post("/", userHandler.CreateUser)
	user.Get("/:id", userHandler.GetUser)
	user.Delete("/:id", userHandler.DeleteUser)
	user.Put("/:id", userHandler.EditUser)
}
