package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-test-demo/constants"
	"go-test-demo/internal/models"
	"go-test-demo/internal/repository"
	"net/http"
)

func GetUser(c *gin.Context) {
	var u struct {
		Id int `json:"id"`
	}
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, models.ResultError(constants.SystemErrorCode, constants.SystemErrorMessage))
		return
	}
	res := repository.GetUserById(u.Id)
	for _, r := range res {
		fmt.Println(r)
	}
	c.JSON(http.StatusOK, models.ResultSuccess(res))
}
