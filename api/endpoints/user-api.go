package endpoints

import (
	"api/dto"
	"api/mappers"
	"api/models"
	"api/repos"
	"api/utilities"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()
var userRepo = repos.InitUserRepo()

func SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userDto dto.CreateUserDto

		if err := ctx.BindJSON(&userDto); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(userDto)
		if validationErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
			return
		}
		user := mappers.CreateUserDtoToUserModel(userDto)
		token, refreshToken, _ := utilities.GenerateAllTokens(user.UserName, user.Email, user.FirstName, user.LastName, user.ID)
		user.Token = token
		user.RefreshToken = refreshToken
		user.Password = utilities.HashPassword(user.Password)

		if err := userRepo.CreateUser(user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, 1)
	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var login dto.LoginDto

		if err := ctx.BindJSON(&login); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		foundUser, err := userRepo.GetUserByEmail(login.Email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong email or password"})
			return
		}
		passIsValid, msg := utilities.VerifyPassword(foundUser.Password, login.Password)
		if passIsValid == false {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		token, refreshToken, err := utilities.GenerateAllTokens(foundUser.UserName, foundUser.Email, foundUser.FirstName, foundUser.LastName, foundUser.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		utilities.UpdateAllTokens(token, refreshToken, foundUser.ID)
		userDto := mappers.UserModelToUserDto(foundUser)
		ctx.JSON(http.StatusOK, userDto)
	}
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, _ := strconv.ParseUint(ctx.Param("id"), 0, 0)

		if err := utilities.MatchUserTypeToUid(ctx, fmt.Sprintf("%d", userId)); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var user models.User
		user, err := userRepo.GetUserById(uint(userId))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userDto := mappers.UserModelToUserDto(user)
		ctx.JSON(http.StatusOK, userDto)
	}
}
