package controller

import (
	"MiniProject/config"
	"MiniProject/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// CreateOrder = untuk admin membuat orderan
func CreateOrder(c echo.Context) error {
	order := models.Order{}

	c.Bind(&order)
	if err := config.DB.Save(&order).Error; err != nil {
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
	if err := config.DB.First(&order, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&order)
	if err := config.DB.Save(&order).Error; err != nil {
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

	order := models.Order{}
	if err := config.DB.First(&order, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&order).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete order",
	})
}

//get orderan by id = bisa untuk admin dan pembeli melihat detail orderan
func GetOrderById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	order := models.Order{}
	if err := config.DB.First(&order, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get order",
		"order":   order,
	})
}

//get all orderan
func GetAllOrder(c echo.Context) error {
	var order []models.Order

	if err := config.DB.Find(&order).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all order",
		"order":   order,
	})
}

//get orderan by id user = bisa untuk admin dan pembeli melihat detail orderan
