package config

import (
	"go.uber.org/fx"
)

// **Proper Fx Module**
var Module = fx.Module("config",
	fx.Provide(ProvideConfig),
)
