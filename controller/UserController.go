package controller

import (
	"MiniProject/config"
	"MiniProject/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//create user
func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

//get all user
func GetAllUser(c echo.Context) error {
	var user []models.User

	if err := config.DB.Find(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all user",
		"user":    user,
	})
}

//get user by id
func GetUserByID(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	user := models.User{}
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user":    user,
	})
}

//update user
func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := models.User{}
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

//delete user
func DeleteUser(c echo.Context) error {
	// your solution here

	id, _ := strconv.Atoi(c.Param("id"))

	user := models.User{}
	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success  delete user",
	})

}

// // login user
// func LoginUsersController(c echo.Context) error {
// 	user := models.User{}
// 	c.Bind(&user)

// 	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": "Fail login user",
// 			"status":  err.Error(),
// 		})
// 	}

// 	token, err := middleware.CreateToken(int(user.ID), user.Name)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message": "Fail create JWT TOken",
// 			"status":  err.Error(),
// 		})
// 	}

// 	user.Token = token

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"status": "success login",
// 		"user":   user,
// 	})
// }
