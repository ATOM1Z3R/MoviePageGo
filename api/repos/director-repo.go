package repos

import "api/models"

func InitDirectorRepo() IDirectorRepo {
	db := conn()
	return &database{
		session: db,
	}
}

func (db *database) GetAllDirectors() ([]models.Director, error) {
	var directors []models.Director
	res := db.session.Find(&directors)
	return directors, res.Error
}

func (db *database) GetDirectorById(id uint) (models.Director, error) {
	var director models.Director
	res := db.session.Preload("Movies").Find(&director, id)
	return director, res.Error
}
