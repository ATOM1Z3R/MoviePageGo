package mappers

import (
	"api/dto"
	"api/models"
)

func CreateCommentDtoToCommentModel(commentDto *dto.CreateCommentDto) models.Comment {
	return models.Comment{
		Name:   commentDto.Name,
		Author: AuthorDtoToUserModel(commentDto.Author),
		Movie:  SimpleMovieDtoToMovieModel(commentDto.Movie),
	}
}

func CommentModelToSimpleCommentDto(comment *models.Comment) *dto.SimpleCommentDto {
	return &dto.SimpleCommentDto{
		Name:        comment.Name,
		Author:      UserModelToAuthorDto(comment.Author),
		Rating:      comment.Rating,
		CreatedDate: comment.CreatedAt.String(),
	}
}

func CreateSimpleCommentDtoList(comments []models.Comment) (commentDtos []*dto.SimpleCommentDto) {
	for _, item := range comments {
		commentDtos = append(commentDtos, CommentModelToSimpleCommentDto(&item))
	}
	return commentDtos
}
