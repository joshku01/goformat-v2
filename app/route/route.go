package route

import (
	"github.com/gin-gonic/gin"
	"goformat-v2/app/global/helper"
	"goformat-v2/app/usecase/response"
)

func SetupRouter(engine *gin.Engine) {
	// middleware
	engine.Use()

	//404 response
	engine.NoRoute(func(c *gin.Context) {
		utileGin := response.Gin{Ctx: c}
		utileGin.Response(404, "Request not exist", nil)
	})

	engine.GET("/ping", func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(1, "pong", nil)
	})

	api := engine.Group("api")
	{
		api.GET("", func(c *gin.Context) {
			_, err2 := helper.Atoi()
			c.JSON(200, gin.H{
				"message": err2,
			})
		})
	}


}
