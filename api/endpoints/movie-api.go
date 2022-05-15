package endpoints

import (
	"api/dto"
	"api/mappers"
	"api/models"
	"api/repos"
	"net/http"
	"strconv"

	gin "github.com/gin-gonic/gin"
)

var movieRepo = repos.InitMovieRepo()

func CreateMovie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var createMovieDto dto.CreateMovieDto

		if err := ctx.BindJSON(&createMovieDto); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		movie := mappers.CreateMovieDtoToMovieModel(createMovieDto)
		validationErr := validate.Struct(movie)
		if validationErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
			return
		}

		err := movieRepo.CreateMovie(movie)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, createMovieDto)
	}
}

func UpdateMovie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		movieId, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)
		var updatedMovie models.Movie

		if err := ctx.BindJSON(&updatedMovie); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		movie, err := movieRepo.GetMovieById(uint(movieId))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		movie = updatedMovie
		if err := movieRepo.UpdateMovie(&movie); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, movie)
	}
}

func GetAllMovies() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		movies, err := movieRepo.GetAllMovies()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		movieDtos := mappers.CreateSimpleMovieDtoList(movies)
		ctx.JSON(http.StatusOK, movieDtos)
	}
}

func GetMovie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		movieId, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)
		movie, err := movieRepo.GetMovieById(uint(movieId))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		movieDto := mappers.MovieModelToMovieDto(&movie)
		ctx.JSON(http.StatusOK, movieDto)
	}
}

func GetMoviesByGenre() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		genreId, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)
		movies, err := movieRepo.FilterMoviesByGenreId(uint(genreId))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		movieDtos := mappers.CreateSimpleMovieDtoList(movies)
		ctx.JSON(http.StatusOK, movieDtos)
	}
}
