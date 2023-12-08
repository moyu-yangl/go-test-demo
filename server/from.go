package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func From(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Println(name)
	c.String(http.StatusOK, "ok")
}
