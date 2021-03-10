package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/config"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/database"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
)

func main() {

	// Setting up environment
	config.SetupEnvVars()

	// Database connection
	database.ConnectDatabase()

	migrations := config.Migrate{DB: database.Instance}
	err := migrations.MigrateAll()

	if err != nil {
		log.Fatalf("cannot migrate database, stack: %s", err.Error())
	}

	app := fiber.New()

	//Helmet security
	app.Use(helmet.New())

	//Handle Cors
	app.Use(cors.New())

	//Rate limiting
	app.Use(limiter.New())

	//Handle panics
	app.Use(recover.New())

	//Handle logs
	app.Use(logger.New())

	//Request ID
	app.Use(requestid.New())

	//Handle routes
	router.SetupRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	log.Fatal(app.Listen(port))

}
