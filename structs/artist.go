package structs

type Artist struct {
	ArtistId int    `db:"ArtistId"`
	Name     string `db:"Name"`
}
