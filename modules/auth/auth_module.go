package auth

import "go.uber.org/fx"

// Module exports dependencies for the user module
var Module = fx.Module("auth",
	fx.Provide(NewAuthService),
	fx.Provide(NewAuthController),
	fx.Provide(NewAuthMiddleware),
)
