package mappers

import (
	"api/dto"
	"api/models"

	"gorm.io/gorm"
)

func DirectorModelToDiretorDto(director *models.Director) *dto.DirectorDto {
	return &dto.DirectorDto{
		Id:           director.ID,
		FirstName:    director.FirstName,
		LastName:     director.LastName,
		BirthDate:    director.BirthDate,
		BirthCountry: director.BirthCountry,
		Movies:       CreateSimpleMovieDtoList(director.Movies),
	}
}

func DirectorModelToSimpleDirectorDto(director *models.Director) *dto.SimpleDirectorDto {
	return &dto.SimpleDirectorDto{
		Id:           director.ID,
		FirstName:    director.FirstName,
		LastName:     director.LastName,
		BirthDate:    director.BirthDate,
		BirthCountry: director.BirthCountry,
	}
}

func SimpleDirectorDtoToDirectorModel(directorDto *dto.SimpleDirectorDto) models.Director {
	return models.Director{
		Model: gorm.Model{
			ID: directorDto.Id,
		},
		FirstName:    directorDto.FirstName,
		LastName:     directorDto.LastName,
		BirthDate:    directorDto.BirthDate,
		BirthCountry: directorDto.BirthCountry,
	}
}

func CreateSimpleDiectorDtoList(directors []models.Director) (directorDtos []*dto.SimpleDirectorDto) {
	for _, item := range directors {
		directorDtos = append(directorDtos, DirectorModelToSimpleDirectorDto(&item))
	}
	return directorDtos
}

func CreateDirectorModelList(directorDtos []*dto.SimpleDirectorDto) (directors []models.Director) {
	for _, item := range directorDtos {
		directors = append(directors, SimpleDirectorDtoToDirectorModel(item))
	}
	return directors
}
