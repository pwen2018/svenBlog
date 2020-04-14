package main

import (
	"svenBlog/middlewares"
	"svenBlog/router"
	"svenBlog/tool"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tool.ParseConfig("config/config.json")
	if err != nil {
		panic(err)
	}
	// 初始化数据库
	dsn := tool.GetDSN(cfg)
	err = tool.InitGorm(cfg, dsn)
	if err != nil {
		panic(err)
	}
	defer tool.DB.Close()
	app := gin.Default()
	// 使用jwt 中间件
	app.Use(middlewares.Cors())
	router.Router(app)
	//app.Use(middlewares.JWTAuth())
	gin.ForceConsoleColor()
	app.Run(cfg.AppHost + ":" + cfg.AppPort)

}
