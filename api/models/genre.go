package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name   string `binding:"required" gorm:"type:varchar(255)"`
	Movies []Movie
}
