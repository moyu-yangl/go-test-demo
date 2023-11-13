package models

type User struct {
	Name  string `json:"name" gorm:""`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (_ User) TableName() string {
	return "user"
}
