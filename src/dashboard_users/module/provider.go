package module

import (
	"backend/src/common/models"
)

type AppsService interface {
	FindOne() models.App
}

type Provider struct {
	AppsService AppsService
}
