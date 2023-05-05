package route

import (
	"MiniProject/constants"
	"MiniProject/controller"
	"MiniProject/middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	middleware.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	// User Controllers
	user := e.Group("/user")
	user.POST("", controller.CreateUser)
	user.POST("/login", controller.LoginUsers)
	user.GET("", controller.GetAllUser, mid.JWT([]byte(constants.SCREAT_JWT)))
	user.GET("/:id", controller.GetUserByID, mid.JWT([]byte(constants.SCREAT_JWT)))
	user.PUT("/:id", controller.UpdateUser, mid.JWT([]byte(constants.SCREAT_JWT)))
	user.DELETE("/:id", controller.DeleteUser, mid.JWT([]byte(constants.SCREAT_JWT)))

	// Baju Controllers
	baju := e.Group("/baju")
	baju.POST("", controller.CreateBaju, mid.JWT([]byte(constants.SCREAT_JWT)))
	baju.GET("", controller.GetAllBaju, mid.JWT([]byte(constants.SCREAT_JWT)))
	baju.GET("/:id", controller.GetBajuByID, mid.JWT([]byte(constants.SCREAT_JWT)))
	baju.PUT("/:id", controller.UpdateBaju, mid.JWT([]byte(constants.SCREAT_JWT)))

	return e
}
