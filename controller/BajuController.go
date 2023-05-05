package controller

import (
	"MiniProject/config"
	"MiniProject/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//GetAllBaju = untuk admin melihat data orderan

func GetAllBaju(c echo.Context) error {
	var baju []models.Baju
	if err := config.DB.Find(&baju).Error; err != nil {
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

	baju := models.Baju{}
	if err := config.DB.Where("id = ?", id).First(&baju).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get baju by id",
		"baju":    baju,
	})
}

//CreateBaju = untuk pembeli membuat orderan

func CreateBaju(c echo.Context) error {
	baju := models.Baju{}
	c.Bind(&baju)

	if err := config.DB.Save(&baju).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create baju",
		"baju":    baju,
	})
}
