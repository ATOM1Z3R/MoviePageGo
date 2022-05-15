package dto

type SimpleCommentDto struct {
	Id          uint       `json:"id"`
	Name        string     `json:"name"`
	Author      *AuthorDto `json:"author,omnitempty"`
	Rating      uint16     `json:"rating"`
	CreatedDate string     `json:"createdDate"`
}

type CreateCommentDto struct {
	Name   string          `json:"name" binding:"required"`
	Author *AuthorDto      `json:"author" binding:"required"`
	Movie  *SimpleMovieDto `json:"movie" binding:"required"`
	Rating uint16          `json:"rating" binding:"required"`
}
