package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"test/controller"
	"test/db"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	db.NewMySQL()
	apiRouter := r.Group("/v1")
	apiRouter.GET("/hello", logging(controller.Hello))
	apiRouter.GET("/world", logging(controller.World))
	apiRouter.GET("/test/:isMan", logging(controller.PathMan))
	apiRouter.POST("/from", logging(controller.From))
	apiRouter.POST("/user", logging(controller.GetUser))

	return r
}

func logging(f gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.Request.URL.Path, c.Request.URL.Host)
		f(c)
	}
}
