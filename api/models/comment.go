package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Name    string `json:"name" binding:"required" gorm:"type:varchar(255)"`
	Author  User   `json:"author" binding:"required" gorm:"foreignKey:UserId"`
	UserId  uint64 `json:"-"`
	Movie   Movie  `json:"movie" binding:"required" gorm:"foreignKey:MovieId"`
	MovieId uint64 `json:"-"`
	Rating  uint16 `json:"rating" binding:"required"`
}
