package repos

import (
	"api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type database struct {
	session *gorm.DB
}

func conn() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect db")
	}
	return db
}

func MakeMigrations(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Comment{},
		&models.Director{},
		&models.Genre{},
		&models.Movie{},
	)
}

func init() {
	MakeMigrations(conn())
}
