package repos

import "api/models"

func InitCommentRepo() ICommentRepo {
	db := conn()
	MakeMigrations(db)
	return &database{
		session: db,
	}
}

func (db *database) CreateComment(comment models.Comment) error {
	err := db.session.Create(&comment).Error
	return err
}

func (db *database) GetCommentsByMovieId(movieId uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := db.session.Where("movieId <> ?", movieId).Find(&comments).Error
	return comments, err
}
