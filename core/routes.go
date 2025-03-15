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
	AuthMiddleware   *auth.AuthMiddleware
	UserController   *user.UserController
	MediaController  *media.MediaController
	CourseController *course.CourseController
	OrderController  *order.OrderController
}

func ProvideRouter(params RouterParams) *gin.Engine {
	r := gin.Default()
	r.Use(params.AuthMiddleware.Authenticate())

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signin", params.AuthController.Signin)
		authRoutes.POST("/signup", params.AuthController.Signup)
	}

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", params.AuthMiddleware.Authorize(true, user.Admin), params.UserController.FindManyUsers)
		userRoutes.POST("/", params.AuthMiddleware.Authorize(true, user.Admin), params.UserController.CreateUsers)
	}

	mediaRoutes := r.Group("/media")
	{
		mediaRoutes.GET("/", params.AuthMiddleware.Authorize(false), params.MediaController.FindManyMedia)
		mediaRoutes.POST("/upload-photo", params.AuthMiddleware.Authorize(true), params.MediaController.UploadPhoto)
	}

	courseRoutes := r.Group("/courses")
	{
		courseRoutes.GET("/", params.CourseController.FindManyCourse)
		courseRoutes.POST("/", params.CourseController.CreateCourse)
	}

	orderRoutes := r.Group("/orders")
	{
		orderRoutes.GET("/", params.OrderController.FindManyOrders)
		orderRoutes.POST("/", params.OrderController.CreateOrder)
		orderRoutes.GET("/:id", params.OrderController.FindOrderByID)
		orderRoutes.PUT("/:id", params.OrderController.UpdateOrderByID)
		orderRoutes.DELETE("/:id", params.OrderController.DeleteOrderByID)
	}

	return r
}
