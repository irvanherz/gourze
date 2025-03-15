package core

import (
	"go.uber.org/fx"
)

// **Proper Fx Module**
var Module = fx.Module("core",
	fx.Provide(ProvideDatabase),
	fx.Provide(ProvideRouter),
)
