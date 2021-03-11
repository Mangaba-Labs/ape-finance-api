package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/config"

	"github.com/Mangaba-Labs/ape-finance-api/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
)

func main() {

	// Setting up environment
	config.SetupEnvVars()

	// Database connection
	db, err := database.NewDatabase()

	if err != nil {
		log.Fatalf("cannot migrate database, err: %s", err.Error())
	}

	migrations := config.Migrate{DB: db}
	err = migrations.MigrateAll()

	if err != nil {
		log.Fatalf("cannot migrate database, err: %s", err.Error())
	}

	app := fiber.New()

	//Helmet security
	app.Use(helmet.New())

	//Handle Cors
	app.Use(cors.New())

	//Handle panics
	app.Use(recover.New())

	//Handle logs
	app.Use(logger.New())

	//Request ID
	app.Use(requestid.New())

	// Router

	server, err := initializeServer()

	if err != nil {
		log.Fatalf("cannot instantiate server, err: %s", err.Error())
	}

	server.SetupRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	log.Fatal(app.Listen(port))

}
