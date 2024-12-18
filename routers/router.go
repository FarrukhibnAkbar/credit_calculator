package routers

import (
	"delivery/configs"
	"delivery/handlers"
	"delivery/logger"
	"delivery/middlewares"

	"github.com/gin-gonic/gin"
)

type Router struct {
	handler handlers.Handler
	config  *configs.Configuration
	router  *gin.Engine
	logger  logger.LoggerI
}

// New creates a new router
func New(h handlers.Handler, cfg *configs.Configuration, logger logger.LoggerI) Router {
	r := gin.New()

	return Router{
		handler: h,
		router:  r,
		logger:  logger,
		config:  cfg,
	}

}

func (r Router) Start() {

	r.router.Use(gin.Logger())
	r.router.Use(gin.Recovery())
	r.router.Use(middlewares.CustomCORSMiddleware())

	r.AdminRouters()

	r.logger.Info("HTTP: Server being started...", logger.String("port", r.config.HTTPPort))

	err := r.router.Run(r.config.HTTPPort)
	if err != nil {
		panic(err)
	}
}
