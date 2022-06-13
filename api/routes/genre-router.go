package routes

import (
	"api/endpoints"

	"github.com/gin-gonic/gin"
)

func GenreRouterGET(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("genres/", endpoints.GetAllGenres())
	incomingRoutes.GET("genre/:id", endpoints.GetGenre())
}
