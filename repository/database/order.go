package database

import (
	"Project-Mini/config"
	"Project-Mini/models"
	"errors"
)

// create order (admin)

func CreateOrderByAdmin(order *models.Order) error {
	// Query user berdasarkan ID user
	if err := config.DB.First(&order.User, order.ID_user).Error; err != nil {
		return errors.New("failed to create order")
	}

	// Query baju berdasarkan ID baju
	if err := config.DB.First(&order.Baju, order.ID_baju).Error; err != nil {
		return errors.New("failed to create order")
	}

	if err := config.DB.Save(&order).Error; err != nil {
		return err
	}

	return nil
}

// get all order (admin)

func GetOrder() (order []models.Order, err error) {
	err = config.DB.Preload("User").Preload("Baju").Find(&order).Error

	if err != nil {
		return []models.Order{}, err
	}
	return
}

// get order by id (admin)

func GetOrderById(id int) (models.Order, error) {
	var order models.Order

	err := config.DB.Preload("User").Preload("Baju").Where("id = ?", id).First(&order).Error

	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

// update order by id (admin)

func UpdateOrder(order models.Order, id interface{}) (models.Order, error) {

	if err := config.DB.Where("id = ?", id).Updates(&order).Error; err != nil {
		return models.Order{}, err
	}

	err := config.DB.Preload("User").Preload("Baju").First(&order, id).Error
	if err != nil {
		return models.Order{}, err
	}

	return order, nil

}

// delete order by id (admin)

func DeleteOrder(id int) (int64, error) {
	order := models.Order{}

	if err := config.DB.First(&order, id).Error; err != nil {
		return 0, err
	}

	if err := config.DB.Delete(&order).Error; err != nil {
		return 0, err
	}

	return config.DB.RowsAffected, nil
}
