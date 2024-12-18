package main

import (
	"delivery/configs"
	"delivery/constants"
	admincontroller "delivery/controllers/admin"
	"delivery/handlers"
	"delivery/logger"
	"delivery/routers"
	"delivery/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	//configuration settings
	cfg := configs.Config()
	validate := validator.New()

	// take environment from config then set gin mode according to it
	switch cfg.Environment {
	case constants.DebugMode:
		gin.SetMode(gin.DebugMode)
	case constants.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	//logger
	log := logger.NewLogger(cfg.AppName, cfg.LogLevel)
	defer logger.Cleanup(log)

	//storage init
	strg := storage.New(cfg)

	//controllers init
	admincontroller := admincontroller.NewAdminController(log, strg)

	//handlers init
	h := handlers.New(
		cfg,
		log,
		admincontroller,
		validate,
	)

	//routers
	router := routers.New(h, cfg, log)

	router.Start()

}
