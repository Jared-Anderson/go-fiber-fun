package api

import (
	"database/sql"
	"fmt"
	"strconv"

	album "go-fun/pkgs/album"

	"github.com/gofiber/fiber/v2"
)

func RegisterAlbumRoutes(app *fiber.App, db *sql.DB) {
	app.Get("/albums", func(c *fiber.Ctx) error {
		data := struct {
			Name string
		}{
			Name: "Jared Anderson",
		}
		return c.Render("views/artist.html", data)
	})

	app.Get("/albums/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return c.SendString("Invalid id")
		}
		albums, err := album.GetAlbumById(db, idInt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error getting albums",
			})
		}
		return c.JSON(albums)
	})
}

// @Summary Get a specific Album by ID
// @Description Get a specific Album by its ID
// @ID get-album
// @Produce json
// @Param id path int true "Album ID"
// @Success 200 {object} structs.Album
// @Router /albums/{id} [get]
func getAlbumById(db *sql.DB, id int64) string {
	fmt.Println("getAlbumById with id: " + strconv.Itoa(int(id)))
	// Get all albums by artist id
	albums, err := album.GetAlbumById(db, id)
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
