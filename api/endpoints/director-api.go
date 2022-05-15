package endpoints

import (
	"api/mappers"
	"api/repos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var directorRepo = repos.InitDirectorRepo()

func GetAllDirectories() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		directors, err := directorRepo.GetAllDirectors()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, mappers.CreateSimpleDiectorDtoList(directors))
	}
}

func GetDirector() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		directorId, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)
		director, err := directorRepo.GetDirectorById(uint(directorId))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		directorDto := mappers.DirectorModelToDiretorDto(&director)
		ctx.JSON(http.StatusOK, directorDto)
	}
}
