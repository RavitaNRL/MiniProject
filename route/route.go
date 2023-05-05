package route

import (
	"MiniProject/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/order", controller.GetOrderById)
	e.POST("/order", controller.CreateOrder)

	return e

}
