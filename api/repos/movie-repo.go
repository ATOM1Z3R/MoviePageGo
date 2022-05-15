package repos

import "api/models"

func InitMovieRepo() IMovieRepo {
	db := conn()
	MakeMigrations(db)
	return &database{
		session: db,
	}
}

func (db *database) GetAllMovies() ([]models.Movie, error) {
	var movies []models.Movie
	res := db.session.Preload("Genre").Find(&movies)
	return movies, res.Error
}

func (db *database) GetMovieById(id uint) (models.Movie, error) {
	var movie models.Movie
	res := db.session.Preload("Comments").Preload("User").Preload("Genre").Preload("Directors").First(&movie, id)
	return movie, res.Error
}

func (db *database) CreateMovie(movie models.Movie) error {
	res := db.session.Create(&movie)
	return res.Error
}

func (db *database) UpdateMovie(movie *models.Movie) error {
	err := db.session.Save(&movie).Error
	return err
}

func (db *database) FilterMoviesByGenreId(genreId uint) ([]models.Movie, error) {
	var movies []models.Movie
	res := db.session.Where("movieId <> ?", genreId).Find(&movies)
	return movies, res.Error
}
