package repos

import "api/models"

func InitGenreRepo() IGenreRepo {
	db := conn()
	MakeMigrations(db)
	return &database{
		session: db,
	}
}

func (db *database) GetAllGenres() ([]models.Genre, error) {
	var genres []models.Genre
	res := db.session.Preload("Movies").Find(&genres)
	return genres, res.Error
}

func (db *database) GetGenreById(id uint) (models.Genre, error) {
	var genre models.Genre
	res := db.session.Preload("Movies").First(&genre, id)
	return genre, res.Error
}
