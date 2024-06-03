package dtos

type LoginDTO struct {
	Username string `json:"username" validate:"required,min=2"`
	Password string `json:"password" validate:"required,min=2,max=50"`
}
