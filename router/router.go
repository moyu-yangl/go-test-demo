package router

import (
	"github.com/gin-gonic/gin"
	"go-test-demo/db"
	"go-test-demo/internal/middleware"
	server2 "go-test-demo/internal/server"
	"log"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	db.NewMySQL()
	apiRouter := r.Group("/v1")
	apiRouter.Use(middleware.Recover)
	{
		ticketRouter := apiRouter.Group("/ticket")
		ticketRouter.GET("/info/:id", logging(server2.TicketInfo))

	}
	{
		userRouter := apiRouter.Group("/user")
		userRouter.POST("/user", logging(server2.GetUser))
	}

	return r
}

func logging(f gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.Request.URL.Path, c.Request.URL.Host)
		f(c)
	}
}
