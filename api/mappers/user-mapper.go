package mappers

import (
	"api/dto"
	"api/models"

	"gorm.io/gorm"
)

func AuthorDtoToUserModel(authorDto *dto.AuthorDto) models.User {
	return models.User{
		Model: gorm.Model{
			ID: authorDto.Id,
		},
		UserName:  authorDto.UserName,
		FirstName: authorDto.FirstName,
		LastName:  authorDto.LastName,
		Email:     authorDto.Email,
	}
}

func UserModelToAuthorDto(author models.User) *dto.AuthorDto {
	return &dto.AuthorDto{
		Id:        author.ID,
		UserName:  author.UserName,
		FirstName: author.FirstName,
		LastName:  author.LastName,
		Email:     author.Email,
	}
}

func UserModelToUserDto(user models.User) *dto.UserDto {
	return &dto.UserDto{
		Id:          user.ID,
		UserName:    user.UserName,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		CreatedDate: user.CreatedAt.String(),
	}
}

func CreateUserDtoToUserModel(userDto dto.CreateUserDto) models.User {
	return models.User{
		UserName:  userDto.UserName,
		FirstName: userDto.FirstName,
		LastName:  userDto.LastName,
		Email:     userDto.Email,
		Password:  userDto.Password,
	}
}
