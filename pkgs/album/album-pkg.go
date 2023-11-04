package pkgs

import (
	"database/sql"
	"fmt"
	"go-fun/database"
	"go-fun/structs"
)

func GetAlbumsByArtistId(db *sql.DB, AlbumId int64) ([]structs.Album, error) {
	albums, er := database.GetAlbumsByArtistId(db, AlbumId)
	if er != nil {
		fmt.Println("Error getting artists")
		fmt.Println(er)
		return nil, er
	}
	return albums, nil
}

func GetAlbumById(db *sql.DB, AlbumId int64) ([]structs.Album, error) {
	albums, er := database.GetAlbumById(db, AlbumId)
	if er != nil {
		fmt.Println("Error getting artists")
		fmt.Println(er)
		return nil, er
	}
	return albums, nil

}
