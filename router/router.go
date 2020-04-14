package router

import (
	"svenBlog/controller"
	"svenBlog/middlewares"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	//api := engine.Group("api")
	// 用户接口
	member := engine.Group("v1")
	{
		member.POST("/login", controller.Login)
		member.POST("/register", controller.Register)
		member.POST("/update", middlewares.JWTAuthMiddleware(), controller.CheckPassword)
	}
}
