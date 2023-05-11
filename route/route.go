package route

import (
	"Project-Mini/constants"
	"Project-Mini/controller"
	"Project-Mini/middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {

	e := echo.New()

	// Middleware
	middleware.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	// User Controllers = dapat digunakan untuk semua user
	user := e.Group("/user")
	user.POST("/register", controller.CreateUserController)
	user.POST("/login", controller.LoginUsers)
	user.GET("", controller.GetAllUserController, mid.JWT([]byte(constants.SCREAT_JWT)))
	user.GET("/:id", controller.GetUserByIDController, mid.JWT([]byte(constants.SCREAT_JWT)))
	user.PUT("/:id", controller.UpdateUserById, mid.JWT([]byte(constants.SCREAT_JWT)))
	user.DELETE("/:id", controller.DeleteUserController, mid.JWT([]byte(constants.SCREAT_JWT)))

	// Baju Controllers = dapat digunakan untuk admin dan buyer
	baju := e.Group("/baju")
	baju.POST("/buyer", controller.CreateOrderBaju, mid.JWT([]byte(constants.SCREAT_JWT)))
	baju.GET("/admin", controller.GetAllBajuController, mid.JWT([]byte(constants.SCREAT_JWT)))
	baju.GET("/:id", controller.GetBajuByID, mid.JWT([]byte(constants.SCREAT_JWT)))
	baju.PUT("/:id", controller.UpdateBajuById, mid.JWT([]byte(constants.SCREAT_JWT)))

	// Order Controllers = dapat digunakan untuk admin dan buyer
	order := e.Group("/order")
	order.POST("", controller.CreateOrderController)
	order.GET("", controller.GetAllOrder)
	order.GET("/:id", controller.GetOrderByIdController)
	order.PUT("/:id", controller.UpdateOrder)
	order.DELETE("/:id", controller.DeleteOrder)

	return e
}
