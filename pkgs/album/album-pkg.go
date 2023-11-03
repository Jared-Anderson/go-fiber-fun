package pkgs

import (
	"database/sql"
	"fmt"
	"go-fun/database"
	"go-fun/structs"
)

func GetAlbumsByArtistId(db *sql.DB, AlbumId int) ([]structs.Album, error) {
	albums, er := database.GetAlbumsByArtistId(db, AlbumId)
	if er != nil {
		fmt.Println("Error getting artists")
		fmt.Println(er)
		return nil, er
	}
	return albums, nil
}
