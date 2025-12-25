package router

import (
	"session-redis/internal/handlers"
	"session-redis/internal/middlewares"
	"session-redis/internal/session"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	rdb := session.CreateClient()
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.POST("/login", handlers.LoginHandler(rdb))
	auth := r.Group("/panel")
	auth.Use(middlewares.AuthMiddleware(rdb))
	auth.GET("/", handlers.PanelHandler())
	auth.POST("/logout", handlers.LogoutHandler(rdb))
	return r
}