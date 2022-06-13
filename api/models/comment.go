package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Name    string `binding:"required" gorm:"type:varchar(255)"`
	Author  User   `binding:"required" gorm:"foreignKey:UserId"`
	UserId  uint64
	Movie   Movie `binding:"required" gorm:"foreignKey:MovieId"`
	MovieId uint64
	Rating  uint16 `binding:"required"`
}
