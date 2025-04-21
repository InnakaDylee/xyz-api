package main

import (
	"xyz/app/router"
	"xyz/middlewares"
	configs "xyz/packages/databases/config"
	"xyz/packages/databases/database"
	"xyz/storages/seed"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load() // Load environment variables from .env file
	
	config, err := configs.LoadConfig()
	if err != nil {
		panic("failed to load configuration")
	}
	// Initialize Echo
	e := echo.New()

	// Set up static file serving
	e.Static("/storages/images", "storages/images")

	// Initialize MySQL connection
	db := database.ConnectMySQL(e)
	seed.MigrateAndSeedProduct(db)

	// Setup middlewares
	middlewares.CORS(e)
	middlewares.Recover(e)
	middlewares.Logger(e)
	middlewares.RateLimiter(e)
	middlewares.RemoveTrailingSlash(e)

	e.Debug = true

	// Setup Router
	router.SetupRouter(e, db)

	// Define a simple route
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	host := config.SERVER.SERVER_HOST
	port := config.SERVER.SERVER_PORT

	address := host + ":" + port
	// Check if the address is empty
	if address == "" {
		address = "localhost:8080" // Default to localhost::8080 if not set
	}
	
	// Print the server address
	e.Logger.Printf("Starting server at %s:%s\n", host, port)
	// Start the server
	if err := e.Start(address); err != nil {
		e.Logger.Fatalf("Failed to start server: %v", err)
	}
}