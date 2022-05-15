package routes

import (
	"api/endpoints"

	"github.com/gin-gonic/gin"
)

func DirectorRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("directors", endpoints.GetAllDirectories())
	incomingRoutes.GET("director/:id", endpoints.GetDirector())
}
