package database

import (
	"database/sql"
	"fmt"
	"go-fun/structs"

	_ "github.com/mattn/go-sqlite3" // Import and register
)

func ConnectToDatabase(connString string) (*sql.DB, error) {
	// Open a database connection
	fmt.Println("Open a database connection: " + connString)
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetArtists(db *sql.DB) ([]structs.Artist, error) {
	query := "SELECT ArtistId, Name FROM Artist"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	// Map the results to artist struct slice
	var artists []structs.Artist
	for rows.Next() {
		var artist structs.Artist
		err := rows.Scan(&artist.ArtistId, &artist.Name)
		if err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}
	return artists, nil
}

func GetAlbumsByArtistId(db *sql.DB, artistId int64) ([]structs.Album, error) {
	query := "SELECT AlbumId, Title FROM Album WHERE ArtistId = ?"
	rows, err := db.Query(query, artistId)
	if err != nil {
		return nil, err
	}

	var albums []structs.Album
	for rows.Next() {
		var album structs.Album
		err := rows.Scan(&album.AlbumId, &album.Title)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}
