package dto

type DirectorDto struct {
	Id           uint              `json:"id"`
	FirstName    string            `json:"firstName"`
	LastName     string            `json:"lastName"`
	BirthDate    string            `json:"birthDate"`
	BirthCountry string            `json:"birthCountry"`
	Movies       []*SimpleMovieDto `json:"movies,omnitempty"`
}

type SimpleDirectorDto struct {
	Id           uint   `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	BirthDate    string `json:"birthDate"`
	BirthCountry string `json:"birthCountry"`
}
