package utilities

import (
	"api/repos"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	UserName  string
	Email     string
	FirstName string
	LastName  string
	UserId    uint
	jwt.StandardClaims
}

var SECRET_KEY = "asd89u9328yryw7fhwe8f73489f743ty87hg837h438f7he4387f"

var userRepo = repos.InitUserRepo()

func GenerateAllTokens(userName string, email string, firstName string, lastName string, userId uint) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		UserName:  userName,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UserId:    userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshToken, err
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

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId uint) {
	user, _ := userRepo.GetUserById(userId)
	user.Token = signedToken
	user.RefreshToken = signedRefreshToken
	if err := userRepo.UpdateUser(&user); err != nil {
		log.Panic(err)
		return
	}
}
