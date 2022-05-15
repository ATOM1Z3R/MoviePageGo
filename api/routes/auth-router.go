package routes

import (
	"api/endpoints"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", endpoints.SignUp())
	incomingRoutes.POST("users/login", endpoints.Login())
}
