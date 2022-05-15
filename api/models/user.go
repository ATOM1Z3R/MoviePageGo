package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string `json:"userName" validate:"required,min=2,max=100"`
	FirstName    string `json:"firstName" validate:"required,min=2,max=100"`
	LastName     string `json:"lastName" validate:"required,min=2,max=100"`
	Email        string `json:"email" validate:"required"`
	Password     string `json:"password" validate:"required,min=6"`
	Token        string `json:"token"`
	UserType     string
	RefreshToken string `json:"refreshToken"`
}
