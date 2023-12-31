package pkgs

import (
	"database/sql"
	"fmt"
	"go-fun/database"
	"go-fun/structs"
)

func GetArtists(db *sql.DB) ([]structs.Artist, error) {
	artists, er := database.GetArtists(db)
	if er != nil {
		fmt.Println("Error getting artists")
		fmt.Println(er)
		return nil, er
	}
	return artists, nil
}

func GetArtistById(db *sql.DB, artistId int64) ([]structs.Artist, error) {
	artists, er := database.GetArtistById(db, artistId)
	if er != nil {
		fmt.Println("Error getting artists")
		fmt.Println(er)
		return nil, er
	}
	return artists, nil
}

func GetAlbumsByArtistId(db *sql.DB, artistId int64) ([]structs.Album, error) {
	albums, er := database.GetAlbumsByArtistId(db, artistId)
	if er != nil {
		fmt.Println("Error getting albums")
		fmt.Println(er)
		return nil, er
	}
	return albums, nil
}
