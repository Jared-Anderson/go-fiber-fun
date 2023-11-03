package structs

type Album struct {
	AlbumId  int    `db:"AlbumId"`
	Title    string `db:"Title"`
	ArtistId int    `db:"ArtistId"`
}
