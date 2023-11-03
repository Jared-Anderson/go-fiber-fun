package main

import (
	"fmt"
	"go-fun/database"
	albumpkg "go-fun/pkgs/album"
	artistpkg "go-fun/pkgs/artist"
)

func main() {
	dbConnStr := "file:///C:/Users/jande/AppData/Roaming/DBeaverData/workspace6/.metadata/sample-database-sqlite-1/Chinook.db"
	// Connect to the database
	db, err := database.ConnectToDatabase(dbConnStr)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close() // Don't forget to close the connection when done

	// Get all artists
	artists, err := artistpkg.GetArtists(db)
	if err != nil {
		fmt.Println("Error getting artists:", err)
		return
	}
	fmt.Println(artists)

	// Get all artists
	albums, err := albumpkg.GetAlbumsByArtistId(db, artists[0].ArtistId)
	if err != nil {
		fmt.Println("Error getting artists:", err)
		return
	}
	fmt.Println()
	fmt.Println(albums)
}