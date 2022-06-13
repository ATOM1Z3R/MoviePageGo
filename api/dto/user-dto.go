package dto

type UserDto struct {
	Id           uint   `json:"id"`
	UserName     string `json:"userName"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	CreatedDate  string `json:"createdDate"`
}

type CreateUserDto struct {
	UserName  string `json:"userName" validate:"required,min=2,max=100"`
	FirstName string `json:"firstName" validate:"required,min=2,max=100"`
	LastName  string `json:"lastName" validate:"required,min=2,max=100"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required,min=6"`
}

type AuthorDto struct {
	Id        uint   `json:"id"`
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type LoginDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type RefreshTokenDto struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
