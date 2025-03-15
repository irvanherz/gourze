package user

import "go.uber.org/fx"

// Module exports dependencies for the user module
var Module = fx.Module("user",
	fx.Provide(NewUserService),
	fx.Provide(NewUserController),
)
