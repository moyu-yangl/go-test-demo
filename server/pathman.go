package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PathMan(c *gin.Context) {
	i, err := strconv.ParseInt(c.Param("isMan"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "isMan err"})
		return
	}
	fmt.Println(i)
	userName := c.DefaultQuery("userName", "")
	fmt.Println(userName)
	c.String(http.StatusOK, "ok")
}
