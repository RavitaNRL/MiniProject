package database

import (
	"Project-Mini/config"
	"Project-Mini/models"
)

// create pesanan baju oleh pembeli
func CreateOrder(baju models.Baju) (models.Baju, error) {

	err := config.DB.Create(&baju).Error

	if err != nil {
		return models.Baju{}, err
	}
	return baju, nil
}

// get all pesanan baju oleh admin
func GetBaju() (baju []models.Baju, err error) {
	err = config.DB.Find(&baju).Error

	if err != nil {
		return []models.Baju{}, err
	}
	return
}

// get pesanan baju by id
func GetBajuById(id int) (models.Baju, error) {
	var baju models.Baju

	if err := config.DB.Where("id = ?", id).First(&baju).Error; err != nil {
		return models.Baju{}, err
	}
	return baju, nil
}

// update pesanan baju oleh pembeli
func UpdateBaju(baju models.Baju, id int) (models.Baju, error) {
	err := config.DB.Where("id = ?", id).Updates(&baju).Error

	if err != nil {
		return models.Baju{}, err
	}

	return baju, nil
}
