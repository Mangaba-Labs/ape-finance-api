package router

import (
	"time"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/api/handler"
	middleware "github.com/Mangaba-Labs/ape-finance-api/pkg/api/middlewares"
	auth "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/auth/handler"
	category "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/handler"
	stock "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/handler"
	userHandler "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// Server structure
type Server struct {
	userHandler     userHandler.Handler
	authHandler     auth.AuthHandler
	stockHandler    stock.Handler
	categoryHandler category.CategoryHandler
}

// NewServer instance
func NewServer(userHandler userHandler.Handler, authHandler auth.AuthHandler, stockHandler stock.Handler, categoryHandler category.CategoryHandler) *Server {
	return &Server{userHandler: userHandler, authHandler: authHandler, stockHandler: stockHandler, categoryHandler: categoryHandler}
}

// SetupRoutes setup router pkg
func (s *Server) SetupRoutes(app *fiber.App) {
	// Api base
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Health
	health := v1.Group("/health")
	health.Get("/", handler.HealthCheck)

	// Auth
	auth := v1.Group("/auth")
	auth.Post("/login", s.authHandler.Login)

	// User
	user := v1.Group("/users", limiter.New(limiter.Config{
		Max:        50,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		}}))

	user.Post("/", s.userHandler.CreateUser)
	user.Get("/:id", s.userHandler.GetUser)
	user.Delete("/:id", s.userHandler.DeleteUser)
	user.Put("/:id", s.userHandler.EditUser)

	// Stock
	stock := v1.Group("/stock")
	stock.Post("/", middleware.Protected(), s.stockHandler.CreateStock)
	stock.Get("/", middleware.Protected(), s.stockHandler.GetStocks)

	// Category
	category := v1.Group("/categories", middleware.Protected())
	category.Delete("/:id", middleware.Protected(), s.categoryHandler.DeleteCategory)
	category.Get("/", middleware.Protected(), s.categoryHandler.GetCategories)
	category.Post("/", middleware.Protected(), s.categoryHandler.CreateCategory)
	category.Put("/:id", middleware.Protected(), s.categoryHandler.EditCategory)
}
