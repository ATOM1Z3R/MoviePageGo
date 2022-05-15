package routes

import (
	"api/endpoints"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/user/:id", endpoints.GetUser())
}
