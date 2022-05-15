package dto

type MovieDto struct {
	Id           uint                 `json:"id"`
	Title        string               `json:"title"`
	PremiereDate string               `json:"premiereDate"`
	Description  string               `json:"description"`
	Location     string               `json:"location"`
	Poster       string               `json:"poster"`
	Teaser       string               `json:"teaser"`
	UserId       uint64               `json:"-"`
	User         *AuthorDto           `json:"author,omnitempty"`
	GenreId      uint64               `json:"-"`
	Genre        *SimpleGenreDto      `json:"genre,omnitempty"`
	Comments     []*SimpleCommentDto  `json:"comments,omnitempty"`
	Directors    []*SimpleDirectorDto `json:"directors,omnitempty"`
}

type SimpleMovieDto struct {
	Id           uint            `json:"id"`
	Title        string          `json:"title"`
	PremiereDate string          `json:"premiereDate"`
	Description  string          `json:"description"`
	Location     string          `json:"location"`
	Genre        *SimpleGenreDto `json:"genre,omnitempty"`
}

type CreateMovieDto struct {
	Title        string               `json:"title" binding:"required"`
	PremiereDate string               `json:"premiereDate" binding:"required"`
	Description  string               `json:"description" binding:"required"`
	Location     string               `json:"location" binding:"required"`
	User         *AuthorDto           `json:"user" binding:"required"`
	Genre        *SimpleGenreDto      `json:"genre" binding:"required"`
	Directors    []*SimpleDirectorDto `json:"directors" binding:"required"`
	Poster       string               `json:"poster" binding:"required"`
	Teaser       string               `json:"teaser"`
}
