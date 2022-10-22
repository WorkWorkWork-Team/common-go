package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

type Config struct {
	ListeningPort string
}

func NewHttpServer(config Config) *gin.Engine {
	router := gin.Default()
	router.GET("/healthcheck", healthCheckHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
