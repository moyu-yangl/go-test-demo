package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func World(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"res": "world"})
}
