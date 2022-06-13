package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName  string `gorm:"type:varchar(255)" validate:"required,min=2,max=100"`
	FirstName string `gorm:"type:varchar(255)" validate:"required,min=2,max=100"`
	LastName  string `gorm:"type:varchar(255)" validate:"required,min=2,max=100"`
	Email     string `gorm:"type:varchar(100)"`
	Password  string
	UserType  string
}
