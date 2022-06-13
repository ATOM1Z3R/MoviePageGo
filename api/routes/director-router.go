package routes

import (
	"api/endpoints"

	"github.com/gin-gonic/gin"
)

func DirectorRouterGET(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("directors", endpoints.GetAllDirectories())
	incomingRoutes.GET("director/:id", endpoints.GetDirector())
}
