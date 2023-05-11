package database

import (
	"Project-Mini/config"
	"Project-Mini/models"
)

// create user (register)
func CreateUser(user models.User) (models.User, error) {

	err := config.DB.Create(&user).Error

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// get all user
func GetAllUser() (user []models.User, err error) {
	err = config.DB.Find(&user).Error

	if err != nil {
		return []models.User{}, err
	}
	return
}

// get user by id
func GetUserById(id interface{}) (models.User, error) {
	var user models.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// update user by id
func UpdateUser(user models.User, id int) (models.User, error) {
	err := config.DB.Where("id = ?", id).Updates(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// delete user by id
func DeleteUserById(id interface{}) error {
	var user models.User

	if err := config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

// login user

func LoginUser(user models.User) (models.User, error) {

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
