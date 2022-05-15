package mappers

import (
	"api/dto"
	"api/models"

	"gorm.io/gorm"
)

func CreateMovieDtoToMovieModel(movieDto dto.CreateMovieDto) models.Movie {
	return models.Movie{
		Title:        movieDto.Title,
		Description:  movieDto.Description,
		PremiereDate: movieDto.PremiereDate,
		Location:     movieDto.Location,
		User:         AuthorDtoToUserModel(movieDto.User),
		Genre:        SimpleGenreDtoToGenreModel(movieDto.Genre),
		Directors:    CreateDirectorModelList(movieDto.Directors),
		Poster:       []byte(movieDto.Poster),
		Teaser:       movieDto.Teaser,
	}
}

func MovieModelToSimpleMovieDto(movie models.Movie) *dto.SimpleMovieDto {
	return &dto.SimpleMovieDto{
		Id:           movie.ID,
		Title:        movie.Title,
		Description:  movie.Description,
		PremiereDate: movie.PremiereDate,
		Location:     movie.Location,
		Genre:        GenreModelToSimpleGenreDto(movie.Genre),
	}
}

func SimpleMovieDtoToMovieModel(movieDto *dto.SimpleMovieDto) models.Movie {
	return models.Movie{
		Model: gorm.Model{
			ID: movieDto.Id,
		},
		Title:        movieDto.Title,
		Description:  movieDto.Description,
		PremiereDate: movieDto.PremiereDate,
		Location:     movieDto.Location,
		Genre:        SimpleGenreDtoToGenreModel(movieDto.Genre),
	}
}

func MovieModelToMovieDto(movie *models.Movie) *dto.MovieDto {
	return &dto.MovieDto{
		Id:           movie.ID,
		Title:        movie.Title,
		Description:  movie.Description,
		PremiereDate: movie.PremiereDate,
		Location:     movie.Location,
		User:         UserModelToAuthorDto(movie.User),
		Genre:        GenreModelToSimpleGenreDto(movie.Genre),
		Comments:     CreateSimpleCommentDtoList(movie.Comments),
		Directors:    CreateSimpleDiectorDtoList(movie.Directors),
		Poster:       string(movie.Poster),
		Teaser:       movie.Teaser,
	}
}

func CreateSimpleMovieDtoList(movies []models.Movie) (movieDtos []*dto.SimpleMovieDto) {
	for _, item := range movies {
		movieDtos = append(movieDtos, MovieModelToSimpleMovieDto(item))
	}
	return movieDtos
}
