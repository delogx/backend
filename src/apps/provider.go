package apps

import "backend/src/common/types"

type AuthService interface {
	AuthMiddleware() types.Middleware
	VerifiedDashboardUserMiddleware() types.Middleware
}

type Provider struct {
	AuthService AuthService
}
