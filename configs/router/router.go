package router

import (
	"L-F/app/controllers"
	"L-F/app/middles"

	"github.com/gin-gonic/gin"
)

func Router(c *gin.Engine) {
	const pre = "/api"
	api := c.Group(pre)
	api.Use(middles.GlobalResponseError)
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.Use(middles.ParseJwt())
		{
			api.POST("/post", controllers.Post)
			api.GET("/list/:page", controllers.Manage)
		}
	}
}
