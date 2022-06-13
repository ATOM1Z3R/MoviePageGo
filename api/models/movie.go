package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title        string `gorm:"type:varchar(100)"`
	PremiereDate string `gorm:"type:varchar(100)"`
	Description  string `gorm:"type:varchar(355)"`
	Location     string `gorm:"type:varchar(75)"`
	Poster       []byte
	Teaser       string `grom:"type:varchar(512)"`
	UserId       uint
	User         User `gorm:"foreignKey:UserId"`
	GenreId      uint
	Genre        Genre `gorm:"foreignKey:GenreId"`
	Comments     []Comment
	Directors    []Director `gorm:"many2many:movies_directors;"`
}
