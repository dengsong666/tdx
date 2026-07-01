package main

import (
	"github.com/gin-gonic/gin"
	"github.com/injoyai/tdx/api/handler"
	"github.com/injoyai/tdx/api/middleware"
)

func NewRouter(app *App) *gin.Engine {
	router := gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		panic(err)
	}
	router.Use(middleware.CORS())

	h := handler.New(app.Client, app.Codes, app.Gbbq, Success, Fail)

	// 所有公开接口统一使用 /api 前缀。
	v1 := router.Group("/api")
	{
		v1.GET("/health", h.Health)
		v1.GET("/quote", h.Quote)
		v1.GET("/kline", h.Kline)
		v1.GET("/minute", h.Minute)
		v1.GET("/trade", h.Trade)
		v1.GET("/codes", h.Codes)
	}

	return router
}
