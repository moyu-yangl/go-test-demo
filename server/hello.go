package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "res", "code": 200, "message": "OK"})
}
