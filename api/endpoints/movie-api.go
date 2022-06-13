package endpoints

import (
	"api/consts"
	"api/dto"
	"api/mappers"
	"api/models"
	"api/repos"
	"net/http"
	"strconv"
	"strings"

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
		userRole := ctx.GetString("userRole")
		userId := ctx.GetUint("userId")

		movie, err := movieRepo.GetMovieById(uint(movieId))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if userRole != consts.ROLE_MOD {
			if movie.UserId != userId {
				ctx.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access to resource"})
				return
			}
		}

		var updatedMovie models.Movie
		if err := ctx.BindJSON(&updatedMovie); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		movie = updatedMovie
		if err := movieRepo.UpdateMovie(&movie); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		movieDto := mappers.MovieModelToMovieDto(&movie)
		ctx.JSON(http.StatusOK, movieDto)
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

func GetOffsetMoviesCollection() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, err := strconv.Atoi(ctx.Param("page"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		pageSize, err := strconv.Atoi(ctx.Param("page_size"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 20:
			pageSize = 20
		case pageSize <= 0:
			pageSize = 5
		}
		offset := (page - 1) * pageSize

		movies, err := movieRepo.GetMovieCollection(offset, pageSize)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		moviesDtos := mappers.CreateMovieDtoList(movies)
		ctx.JSON(http.StatusOK, moviesDtos)
	}
}

func GetMovieByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		title := strings.Trim(ctx.Param("title"), " ")

		movie, err := movieRepo.GetMovieByName(title)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		movieDto := mappers.MovieModelToMovieDto(&movie)
		ctx.JSON(http.StatusOK, movieDto)
	}
}
