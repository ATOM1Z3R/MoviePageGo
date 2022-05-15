package dto

type GenreDto struct {
	Id     uint              `json:"id"`
	Name   string            `json:"name"`
	Movies []*SimpleMovieDto `json:"movies,omnitempty"`
}

type SimpleGenreDto struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
