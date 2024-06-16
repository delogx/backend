package dtos

type CreateAppDto struct {
	Name     string `json:"name" validate:"required,min=2,max=255"`
	HostName string `json:"host_name" validate:"required,min=2,max=255,hostname"`
}
