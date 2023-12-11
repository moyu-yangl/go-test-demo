package models

type User struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

func (_ User) TableName() string {
	return "user"
}
