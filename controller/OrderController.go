package controller

import (
	"Project-Mini/models"
	"Project-Mini/repository/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// CreateOrder = untuk admin membuat orderan
func CreateOrderController(c echo.Context) error {
	order := models.Order{}
	c.Bind(&order)

	if err := database.CreateOrderByAdmin(&order); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create order",
		"order":   order,
	})
}

// update order
func UpdateOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	order := models.Order{}
	c.Bind(&order)

	order, err := database.UpdateOrder(order, id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update order",
		"order":   order,
	})
}

//delete orderan
func DeleteOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := database.DeleteOrder(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete order",
	})
}

//get orderan by id = bisa untuk admin dan pembeli melihat detail orderan
// func GetOrderById(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	order := models.Order{}
// 	if err := config.DB.First(&order, id).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success get order",
// 		"order":   order,
// 	})
// }

func GetOrderByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := database.GetOrderById(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get order",
		"order":   order,
	})
}

//get all orderan
// func GetAllOrder(c echo.Context) error {
// 	var order []models.Order

// 	if err := config.DB.Find(&order).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success get all order",
// 		"order":   order,
// 	})
// }

func GetAllOrder(c echo.Context) error {
	order, err := database.GetOrder()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all order",
		"order":   order,
	})
}
