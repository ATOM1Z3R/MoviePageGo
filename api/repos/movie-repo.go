package repos

import "api/models"

type IMovieRepo interface {
	GetAllMovies() (movies []models.Movie, err error)
	GetMovieById(id uint) (movie models.Movie, err error)
	CreateMovie(movie models.Movie) (err error)
	FilterMoviesByGenreId(genreId uint) (movies []models.Movie, err error)
	UpdateMovie(movie *models.Movie) (err error)
	GetMovieCollection(offset int, numberOfItems int) (movies []models.Movie, err error)
	GetMovieByName(name string) (movie models.Movie, err error)
}

func InitMovieRepo() IMovieRepo {
	db := conn()
	return &database{
		session: db,
	}
}

func (db *database) GetAllMovies() (movies []models.Movie, err error) {
	err = db.session.Preload("Genre").Find(&movies).Error
	return movies, err
}

func (db *database) GetMovieById(id uint) (movie models.Movie, err error) {
	preloads := db.session.Preload("Comments").Preload("User").Preload("Genre").Preload("Directors")
	err = preloads.First(&movie, id).Error
	return movie, err
}

func (db *database) CreateMovie(movie models.Movie) (err error) {
	return db.session.Create(&movie).Error
}

func (db *database) UpdateMovie(movie *models.Movie) (err error) {
	return db.session.Save(&movie).Error
}

func (db *database) FilterMoviesByGenreId(genreId uint) (movies []models.Movie, err error) {
	err = db.session.Preload("Genre").Where("genre_Id = ?", genreId).Find(&movies).Error
	return movies, err
}

func (db *database) GetMovieCollection(offset int, numberOfItems int) (movies []models.Movie, err error) {
	preloads := db.session.Preload("Genre").Preload("Comments").Preload("Directors").Preload("User")
	err = preloads.Offset(offset).Limit(numberOfItems).Find(&movies).Error
	return movies, err
}

func (db *database) GetMovieByName(name string) (movie models.Movie, err error) {
	preloads := db.session.Preload("Genre").Preload("Comments").Preload("Directors").Preload("User")
	err = preloads.Where("Title LIKE ?", name).First(&movie).Error
	return movie, err
}
