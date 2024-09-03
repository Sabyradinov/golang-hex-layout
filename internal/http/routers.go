package http

import (
	"github.com/Sabyradinov/golang-hex-layout/config"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/logger"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

type Router struct {
	*gin.Engine
}

func NewRouter(cfg *config.Configs, log logger.AppLogger, services *service.Builder) (*Router, error) {
	//create gin instance
	router := gin.New()
	//create handlers
	handler := newHandler(cfg, log, services)

	//config gin logger
	if cfg.WebServer.GIN.UseLogger {
		router.Use(gin.Logger())
	}
	//config panic handler
	if cfg.WebServer.GIN.UseRecovery {
		router.Use(gin.CustomRecovery(log.HttpPanicHandler))
	}

	//config swagger ui
	docs.SwaggerInfo.Host = cfg.SwaggerUI.Host
	docs.SwaggerInfo.Description = cfg.SwaggerUI.Description
	docs.SwaggerInfo.Title = cfg.SwaggerUI.PageTitle

	docs.SwaggerInfo.InstanceName()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//health check
	router.GET("/health")

	//config routes
	rg := router.Group("/api/v1")
	{
		rg.Use(handler.middleware.UserTokenVerify).
			Use(handler.middleware.Locale)
		market := rg.Group("/market")
		{
			market.POST("/buy", handler.market.Buy)
			market.GET("/:id", handler.market.GetById)
		}

		shipment := rg.Group("/shipment")
		{
			shipment.POST("/delivery", handler.shipment.DeliveryToClient)
			shipment.GET("/:id", handler.shipment.GetMyDelivery)
		}
	}

	return &Router{router}, nil
}
