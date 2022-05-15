package endpoints

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"api/dto"
	"api/mappers"
	"api/repos"
)

var commentRepo = repos.InitCommentRepo()

func CreateComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var commentDto dto.CreateCommentDto
		movieId, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)

		if err := ctx.BindJSON(&commentDto); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		comment := mappers.CreateCommentDtoToCommentModel(&commentDto)
		validationErr := validate.Struct(comment)
		if validationErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
			return
		}

		err := commentRepo.CreateComment(comment)
		println(comment.Movie.Title)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		comments, err := commentRepo.GetCommentsByMovieId(uint(movieId))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, comments)
	}
}
