package media

import "go.uber.org/fx"

// Module exports dependencies for the media module
var Module = fx.Module("media",
	fx.Provide(NewMediaService),
	fx.Provide(NewMediaController),
	fx.Provide(NewBunnyService),
)
