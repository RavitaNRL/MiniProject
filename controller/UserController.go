package controller

import (
	"Project-Mini/middleware"
	"Project-Mini/models"
	"Project-Mini/repository/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//create user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	user, err := database.CreateUser(user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

//get all user
func GetAllUserController(c echo.Context) error {
	user, err := database.GetAllUser()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all user",
		"user":    user,
	})
}

//get user by id
func GetUserByIDController(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := database.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"user":    user,
	})
}

// update user
func UpdateUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := models.User{}
	c.Bind(&user)

	user, err := database.UpdateUser(user, id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

//delete user by id (admin)
func DeleteUserController(c echo.Context) error {
	// your solution here

	id, _ := strconv.Atoi(c.Param("id"))

	err := database.DeleteUserById(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success  delete user",
	})

}

// login user
func LoginUsers(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	user, err := database.LoginUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail login user",
			"status":  err.Error(),
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT TOken",
			"status":  err.Error(),
		})
	}

	user.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"user":   user,
	})
}
