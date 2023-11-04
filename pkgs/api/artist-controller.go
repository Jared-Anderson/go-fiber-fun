package api

import (
	"database/sql"
	"fmt"
	"strconv"

	artist "go-fun/pkgs/artist"

	"github.com/gofiber/fiber/v2"
)

func RegisterArtistRoutes(app *fiber.App, db *sql.DB) {
	app.Get("/artists", func(c *fiber.Ctx) error {
		artists, err := artist.GetArtists(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error getting artists")
		}
		return c.JSON(artists)
	})

	app.Get("/artists/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid id")
		}
		artists, err := artist.GetArtistById(db, idInt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error getting artist")
		}
		if len(artists) == 0 {
			return c.Status(fiber.StatusNotFound).SendString("Artist not found")
		}
		return c.JSON(artists[0])
	})

	app.Get("/artists/:id/albums", func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid id")
		}
		albums, err := artist.GetAlbumsByArtistId(db, idInt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error getting albums")
		}
		return c.JSON(albums)
	})
}

// @Summary Get a list of all artists
// @Description Get a list of all artists in the database
// @ID get-artists
// @Produce json
// @Success 200 {object} []structs.Artist
// @Router /artists [get]
func getAllArtists(db *sql.DB) string {
	fmt.Println("getAllArtists")
	// Get all artists
	artists, err := artist.GetArtists(db)
	if err != nil {
		return "Nothing to return."
	}

	res := ""
	for _, artist := range artists {
		res += strconv.Itoa(artist.ArtistId) + ". " + artist.Name + "\n"
	}
	return res
}

// @Summary Get a specific artist by ID
// @Description Get a specific artist by its ID
// @ID get-artist-by-id
// @Produce json
// @Param id path int true "Artist ID"
// @Success 200 {object} structs.Artist
// @Router /artists/{id} [get]
func getArtistById(db *sql.DB, id int64) string {
	fmt.Println("getArtistById with id: " + strconv.Itoa(int(id)))
	// Get artist by id
	artists, err := artist.GetArtistById(db, id)
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

// @Summary Get a specific artists albums by artist ID
// @Description Get all albums for a specific artist ID
// @ID get-artist-albums-by-id
// @Produce json
// @Param id path int true "Artist ID"
// @Success 200 {object} []structs.Album
// @Router /artists/{id}/albums [get]
func getAlbumsByArtistId(db *sql.DB, id int64) string {
	fmt.Println("getAlbumsByArtistId with id: " + strconv.Itoa(int(id)))
	// Get all albums by artist id
	albums, err := artist.GetAlbumsByArtistId(db, id)
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
