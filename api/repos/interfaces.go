package repos

import "api/models"

type ICommentRepo interface {
	CreateComment(comment models.Comment) error
	GetCommentsByMovieId(movieId uint) ([]models.Comment, error)
}

type IMovieRepo interface {
	GetAllMovies() ([]models.Movie, error)
	GetMovieById(id uint) (models.Movie, error)
	CreateMovie(movie models.Movie) error
	FilterMoviesByGenreId(id uint) ([]models.Movie, error)
	UpdateMovie(movie *models.Movie) error
}

type IDirectorRepo interface {
	GetAllDirectors() ([]models.Director, error)
	GetDirectorById(id uint) (models.Director, error)
}

type IUserRepo interface {
	CreateUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserById(id uint) (models.User, error)
	UpdateUser(user *models.User) error
}

type IGenreRepo interface {
	GetAllGenres() ([]models.Genre, error)
	GetGenreById(id uint) (models.Genre, error)
}
