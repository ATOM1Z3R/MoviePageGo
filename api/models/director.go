package models

import "gorm.io/gorm"

type Director struct {
	gorm.Model
	FirstName    string  `json:"firstName" binding:"required" gorm:"type:varchar(100)"`
	LastName     string  `json:"lastName" binding:"required" gorm:"type:varchar(100)"`
	BirthDate    string  `json:"birthDate" binding:"required" gorm:"type:varchar(175)"`
	BirthCountry string  `json:"birthCountry" binding:"required" gorm:"type:varchar(155)"`
	Movies       []Movie `json:"movies" binding:"required" gorm:"many2many:movies_directors;"`
}
