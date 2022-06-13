package main

import (
	"api/middlewares"
	"api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(middlewares.CORSMiddleware())

	routes.AuthRoutes(router)
	routes.GenreRouterGET(router)
	routes.MovieRouterGET(router)
	routes.DirectorRouterGET(router)
	//Auth required
	router.Use(middlewares.Auth())
	routes.UserRoutes(router)
	routes.CommentRouterPOST(router)
	routes.MovieRouterPOST(router)

	router.Run(":8080")

}
