package routes

import (
	"api/endpoints"

	"github.com/gin-gonic/gin"
)

func CommentRouterPOST(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("comment/add", endpoints.CreateComment())
}
