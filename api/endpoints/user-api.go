package endpoints

import (
	"api/consts"
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
		user.UserType = consts.ROLE_USER
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
		token, refreshToken, err := utilities.GenerateAllTokens(
			foundUser.ID,
			foundUser.UserName,
			foundUser.Email,
			foundUser.FirstName,
			foundUser.LastName,
			foundUser.UserType)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		tokens := map[string]string{
			"token":        token,
			"refreshToken": refreshToken,
		}
		ctx.JSON(http.StatusOK, tokens)
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

func RegenerateAccessToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var refreshToken dto.RefreshTokenDto
		if err := ctx.BindJSON(&refreshToken); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		claims, errMsg := utilities.ValidateToken(refreshToken.RefreshToken)
		if errMsg != "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
			ctx.Abort()
			return
		}
		if claims.TokenType != consts.TOKEN_REFRESH {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong token type"})
			return
		}

		user, err := userRepo.GetUserById(claims.UserId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := utilities.GenerateAccessToken(user.ID, user.UserName, user.Email, user.FirstName, user.LastName, user.UserType)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, map[string]string{"token": token})
	}
}
