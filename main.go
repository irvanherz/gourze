package main

import (
	"github.com/gin-gonic/gin"
	"github.com/irvanherz/gourze/config"
	"github.com/irvanherz/gourze/core"
	"github.com/irvanherz/gourze/modules/auth"
	"github.com/irvanherz/gourze/modules/course"
	"github.com/irvanherz/gourze/modules/media"
	"github.com/irvanherz/gourze/modules/order"
	"github.com/irvanherz/gourze/modules/user"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module, // Provide config and routing
		core.Module,   // Provide core module dependencies
		user.Module,   // Provide user module dependencies
		auth.Module,   // Provide auth module dependencies
		media.Module,  // Provide media module dependencies
		course.Module, // Provide course module dependencies
		order.Module,  // Provide order module dependencies
		fx.Invoke(func(router *gin.Engine) {
			router.Run(":8080") // Start Gin server
		}),
	)

	app.Run()
}
