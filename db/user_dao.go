package db

import "test/models"

func GetUserByName(name string) []models.User {
	var res []models.User
	DB.Where(&models.User{Name: name}).Find(&res)
	return res
}
