package httpserver

import (
	"github.com/gin-gonic/gin"
)

func NewHttpServer(proxy string) *gin.Engine {
	router := gin.Default()
	api := router.Group(proxy)
	api.GET("/healthcheck", healthCheckHandler)
	return router
}
