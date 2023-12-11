package main

import (
	"github.com/gin-gonic/gin"
	"go-test-demo/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8080")

	//logger := log.GetLogger()
	//logger.Error("error")
}
