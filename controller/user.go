package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-test-demo/db"
	"go-test-demo/models"
	"net/http"
)

func GetUser(c *gin.Context) {
	//var user map[string]interface{}
	//var u User
	var u struct {
		Age         int
		Name, Phone string
		//Name string
		//Name string `json:"name"`
	}
	err := c.BindJSON(&u)
	if err != nil {
		c.String(http.StatusOK, "err")
		return
	}
	//fmt.Println(u.name)
	res := db.GetUserByName(u.Name)
	for _, r := range res {
		fmt.Println(r)
	}
	c.JSON(http.StatusOK, models.ResultSuccess(res))
}
