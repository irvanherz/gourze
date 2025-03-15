package course

import "go.uber.org/fx"

// Module exports dependencies for the course module
var Module = fx.Module("course",
	fx.Provide(NewCourseService),
	fx.Provide(NewCourseController),
)
