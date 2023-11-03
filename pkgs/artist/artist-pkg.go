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
