package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-test-demo/internal/repository"
	"go-test-demo/models"
	"net/http"
)

func GetUser(c *gin.Context) {
	//var user map[string]interface{}
	//var u User
	var u struct {
		Id int `json:"id"`
		//Name string
		//Name string `json:"name"`
	}
	err := c.BindJSON(&u)
	if err != nil {
		c.String(http.StatusOK, "err")
		return
	}
	//fmt.Println(u.name)
	res := repository.GetUserById(u.Id)
	for _, r := range res {
		fmt.Println(r)
	}
	c.JSON(http.StatusOK, models.ResultSuccess(res))
}
