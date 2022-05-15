package endpoints

import (
	"api/mappers"
	"api/repos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var genreRepo = repos.InitGenreRepo()

func GetAllGenres() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		genres, err := genreRepo.GetAllGenres()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		genreDtos := mappers.CreateGenreDtoList(genres)
		ctx.JSON(http.StatusOK, genreDtos)
	}
}

func GetGenre() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		genreId, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)
		genre, err := genreRepo.GetGenreById(uint(genreId))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		genreDto := mappers.GenreModelToGenreDto(&genre)
		ctx.JSON(http.StatusOK, genreDto)
	}
}
