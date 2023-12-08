package router

import (
	"github.com/gin-gonic/gin"
	"go-test-demo/db"
	"go-test-demo/server"
	"log"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	db.NewMySQL()
	apiRouter := r.Group("/v1")
	apiRouter.GET("/hello", logging(server.Hello))
	apiRouter.GET("/world", logging(server.World))
	apiRouter.GET("/test/:isMan", logging(server.PathMan))
	apiRouter.POST("/from", logging(server.From))
	apiRouter.POST("/user", logging(server.GetUser))

	return r
}

func logging(f gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.Request.URL.Path, c.Request.URL.Host)
		f(c)
	}
}
