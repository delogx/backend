package apps

import "backend/src/common/types"

type AuthService interface {
	AuthMiddleware() types.Middleware
	VerifiedDashboardUserMiddleware(db types.DB) types.Middleware
}

type Provider struct {
	AuthService AuthService
	DB          types.DB
}
