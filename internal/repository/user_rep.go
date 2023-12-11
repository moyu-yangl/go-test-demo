package repository

import (
	"go-test-demo/db"
	"go-test-demo/models"
)

func GetUserByName(name string) []models.User {
	var res []models.User
	db.DB.Where(&models.User{Name: name}).Find(&res)
	return res
}

func GetUserById(id int) []models.User {
	var res []models.User
	db.DB.Where(&models.User{Id: id}).Find(&res)
	return res
}
