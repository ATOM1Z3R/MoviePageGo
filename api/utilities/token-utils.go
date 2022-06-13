package utilities

import (
	"api/consts"
	"api/repos"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	UserId    uint
	UserName  string
	Email     string
	FirstName string
	LastName  string
	UserRole  string
	TokenType string
	jwt.StandardClaims
}

var SECRET_KEY = "asd89u9328yryw7fhwe8f73489f743ty87hg837h438f7he4387f"

var userRepo = repos.InitUserRepo()

func GenerateAllTokens(
	userId uint,
	userName string,
	email string,
	firstName string,
	lastName string,
	userRole string) (signedToken string, signedRefreshToken string, err error) {

	token, err := GenerateAccessToken(userId, userName, email, firstName, lastName, userRole)
	refreshToken, err := GenerateRefreshToken(userId)

	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshToken, err
}

func GenerateRefreshToken(userId uint) (signedRefreshToken string, err error) {
	refreshClaims := &SignedDetails{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(336)).Unix(),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}
	return refreshToken, err
}

func GenerateAccessToken(userId uint, userName string, email string, firstName string, lastName string, userRole string) (
	signedToken string, err error) {
	claims := &SignedDetails{
		UserName:  userName,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UserId:    userId,
		TokenType: consts.TOKEN_ACCESS,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(12)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}
	return token, err
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}
	return claims, msg
}
