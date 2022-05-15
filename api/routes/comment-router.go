package routes

import (
	"api/endpoints"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("comment/add", endpoints.CreateComment())
}
