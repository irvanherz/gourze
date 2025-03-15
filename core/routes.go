package core

import (
	"github.com/gin-gonic/gin"
	"github.com/irvanherz/gourze/modules/auth"
	"github.com/irvanherz/gourze/modules/course"
	"github.com/irvanherz/gourze/modules/media"
	"github.com/irvanherz/gourze/modules/order"
	"github.com/irvanherz/gourze/modules/user"
	"go.uber.org/fx"
)

type RouterParams struct {
	fx.In
	AuthController   *auth.AuthController
	UserController   *user.UserController
	MediaController  *media.MediaController
	CourseController *course.CourseController
	OrderController  *order.OrderController
}

func ProvideRouter(params RouterParams) *gin.Engine {
	r := gin.Default()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signin", params.AuthController.Signin)
		authRoutes.POST("/signup", params.AuthController.Signup)
	}

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", params.UserController.FindMany)
		userRoutes.POST("/", params.UserController.Create)
	}

	mediaRoutes := r.Group("/media")
	{
		mediaRoutes.GET("/", params.MediaController.FindMany)
		mediaRoutes.POST("/", params.MediaController.Create)
		mediaRoutes.POST("/upload-photo", params.MediaController.UploadPhoto)
	}

	courseRoutes := r.Group("/courses")
	{
		courseRoutes.GET("/", params.CourseController.FindManyCourse)
		courseRoutes.POST("/", params.CourseController.CreateCourse)
	}

	orderRoutes := r.Group("/orders")
	{
		orderRoutes.GET("/", params.OrderController.FindMany)
		orderRoutes.POST("/", params.OrderController.Create)
		orderRoutes.GET("/:id", params.OrderController.FindOne)
		orderRoutes.PUT("/:id", params.OrderController.UpdateByID)
		orderRoutes.DELETE("/:id", params.OrderController.DeleteByID)
	}

	return r
}
