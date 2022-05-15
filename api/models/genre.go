package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name   string  `json:"name" binding:"required" gorm:"type:varchar(255)"`
	Movies []Movie `json:"movies,omnitempty"`
}
