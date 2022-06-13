package repos

import (
	"api/models"
)

func InitUserRepo() IUserRepo {
	db := conn()
	return &database{
		session: db,
	}
}

func (db *database) CreateUser(user models.User) error {
	res := db.session.Create(&user)
	return res.Error
}

func (db *database) GetAllUsers() ([]models.User, error) {
	var users []models.User
	res := db.session.Find(&users)
	return users, res.Error
}

func (db *database) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	res := db.session.Where("Email LIKE ?", email).Find(&user)
	return user, res.Error
}

func (db *database) GetUserById(id uint) (models.User, error) {
	var user models.User
	res := db.session.Find(&user, id)
	return user, res.Error
}

func (db *database) UpdateUser(user *models.User) error {
	err := db.session.Save(&user).Error
	return err
}
