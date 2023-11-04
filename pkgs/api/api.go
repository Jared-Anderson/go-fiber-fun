package api

import (
	// Import go fiber
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	// Import go fiber middleware
	"fmt"
	"go-fun/database"
	album "go-fun/pkgs/album"
	artist "go-fun/pkgs/artist"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

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

	app.Get("/artists", func(c *fiber.Ctx) error {
		return c.SendString(getAllArtists(db))
	})

	app.Get("/albums", func(c *fiber.Ctx) error {
		data := struct {
			Name string
		}{
			Name: "John",
		}
		return c.Render("views/artist.html", data)
	})

	app.Get("/albums/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return c.SendString("Invalid id")
		}
		return c.SendString(getAlbumsByArtistId(db, idInt))
	})

	// Start server on http://localhost:8080
	fmt.Println("Server started on port 8080")
	app.Listen(":8080")
}

func getAllArtists(db *sql.DB) string {
	fmt.Println("getAllArtists")
	// Get all artists
	artists, err := artist.GetArtists(db)
	if err != nil {
		fmt.Println("Error getting artists:", err)
		return "Nothing to return."
	}

	res := ""
	for _, artist := range artists {
		res += strconv.Itoa(artist.ArtistId) + ". " + artist.Name + "\n"
	}
	return res
}

func getAlbumsByArtistId(db *sql.DB, id int64) string {
	fmt.Println("getAlbumsByArtistId with id: " + strconv.Itoa(int(id)))
	// Get all albums by artist id
	albums, err := album.GetAlbumsByArtistId(db, id)
	if err != nil {
		fmt.Println("Error getting artists:", err)
		return "Nothing to return."
	}
	res := ""
	for _, album := range albums {
		res += strconv.Itoa(album.AlbumId) + ". " + album.Title + "\n"
	}
	return res
}
