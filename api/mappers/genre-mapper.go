package mappers

import (
	"api/dto"
	"api/models"

	"gorm.io/gorm"
)

func SimpleGenreDtoToGenreModel(genreDto *dto.SimpleGenreDto) models.Genre {
	return models.Genre{
		Model: gorm.Model{
			ID: genreDto.Id,
		},
		Name: genreDto.Name,
	}
}

func GenreModelToGenreDto(genre *models.Genre) *dto.GenreDto {
	return &dto.GenreDto{
		Id:     genre.ID,
		Name:   genre.Name,
		Movies: CreateSimpleMovieDtoList(genre.Movies),
	}
}

func GenreModelToSimpleGenreDto(genre models.Genre) *dto.SimpleGenreDto {
	return &dto.SimpleGenreDto{
		Id:   genre.ID,
		Name: genre.Name,
	}
}

func CreateGenreDtoList(genres []models.Genre) (genresDto []*dto.GenreDto) {
	for _, item := range genres {
		genresDto = append(genresDto, GenreModelToGenreDto(&item))
	}
	return genresDto
}
