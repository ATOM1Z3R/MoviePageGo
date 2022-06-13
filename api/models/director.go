package models

import "gorm.io/gorm"

type Director struct {
	gorm.Model
	FirstName    string  `binding:"required" gorm:"type:varchar(100)"`
	LastName     string  `binding:"required" gorm:"type:varchar(100)"`
	BirthDate    string  `binding:"required" gorm:"type:varchar(175)"`
	BirthCountry string  `binding:"required" gorm:"type:varchar(155)"`
	Movies       []Movie `binding:"required" gorm:"many2many:movies_directors;"`
}
