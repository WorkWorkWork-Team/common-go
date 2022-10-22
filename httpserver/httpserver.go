package httpserver

import (
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

type Config struct {
	ListeningPort string
}

type Server struct {
	config    Config
	ginEngine *gin.Engine
}

func NewHttpServer(config Config) (server Server) {
	router := gin.Default()
	router.GET("/healthcheck", healthCheckHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	server.config = config
	server.ginEngine = router
	return server
}
