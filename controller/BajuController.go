package controller

import (
	"Project-Mini/models"
	"Project-Mini/repository/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//GetAllBaju = untuk admin melihat data orderan

func GetAllBajuController(c echo.Context) error {
	baju, err := database.GetBaju()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all baju",
		"baju":    baju,
	})

}

//GetBajuByID = untuk pembeli melihat atau mengedit data orderan

func GetBajuByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	baju, err := database.GetBajuById(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get baju by id",
		"baju":    baju,
	})
}

//CreateBaju = untuk pembeli membuat orderan

func CreateOrderBaju(c echo.Context) error {
	baju := models.Baju{}
	c.Bind(&baju)

	baju, err := database.CreateOrder(baju)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create baju",
		"baju":    baju,
	})
}

// UpdateBaju = untuk pembeli mengedit data orderan
func UpdateBajuById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	baju := models.Baju{}
	c.Bind(&baju)

	baju, err := database.UpdateBaju(baju, id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update baju",
		"baju":    baju,
	})
}
