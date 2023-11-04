package api

import (
	"fmt"
	"go-fun/database"

	"github.com/gofiber/fiber/v2"

	_ "go-fun/docs/go-fiber-fun" // This line is necessary for go-swagger to find your docs!

	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger" // Import Fiber-Swagger middleware
)

// cmd: swag init -g pkgs/api/api.go --output docs/go-fiber-fun

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
// StartServer starts the server
func StartServer(dbConnStr string) {
	// Connect to the database
	db, err := database.ConnectToDatabase(dbConnStr)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close() // Don't forget to close the connection when done

	// Create a new Fiber instance
	app := fiber.New()
	app.Use(cors.New())

	// Register routes
	app.Get("/health", healthCheck)
	registerSwaggerRoutes(app)
	RegisterArtistRoutes(app, db)
	RegisterAlbumRoutes(app, db)

	// Start server on http://localhost:8080
	fmt.Println("Server started on port 8080")
	app.Listen(":8080")
}

func registerSwaggerRoutes(app *fiber.App) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func healthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
