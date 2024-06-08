package dtos

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email,min=2"`
	Password string `json:"password" validate:"required,min=2,max=50"`
}
