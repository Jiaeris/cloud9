package ginserver

import (
	"cloud9/handler"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	mainGroup := router.Group("/cloud")
	{ //账户相关部分
		accountGroup := mainGroup.Group("/api/account")
		{
			accountGroup.GET("/unLogin", handler.UnLogin)
			accountGroup.POST("/login", handler.Login)
			accountGroup.POST("/register", handler.Register)
		}
		accountGroup.Use(loginMiddle())
		{
			accountGroup.POST("/logout", handler.Logout)
			accountGroup.GET("/detail", handler.Detail)
			accountGroup.POST("/detail/update", handler.DetailUpdate)

		}
	}
	router.NoRoute(handler.ServerPing)
	router.Run(":8080")
}
