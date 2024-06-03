package dtos

type VerifyEmailDto struct {
	VerificationToken string `json:"verification_token" validate:"required,min=1"`
}
