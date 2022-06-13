package routes

import (
	"api/endpoints"

	"github.com/gin-gonic/gin"
)

func MovieRouterPOST(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("movie/add", endpoints.CreateMovie())
	incomingRoutes.POST("movie/update/:id", endpoints.UpdateMovie())
}

func MovieRouterGET(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("movies/", endpoints.GetAllMovies())
	incomingRoutes.GET("movie/:id", endpoints.GetMovie())
	incomingRoutes.GET("movie/genre/:id", endpoints.GetMoviesByGenre())
	incomingRoutes.GET("movies/:page/:page_size", endpoints.GetOffsetMoviesCollection())
	incomingRoutes.GET("movies/search/:title", endpoints.GetMovieByName())
}
