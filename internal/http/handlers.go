package http

import (
	"github.com/Sabyradinov/golang-hex-layout/config"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/logger"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/service"
	"github.com/Sabyradinov/golang-hex-layout/internal/handler"
)

type Handlers struct {
	middleware *handler.Middleware
	market     *handler.Market
	shipment   *handler.Shipment
}

func newHandler(cfg *config.Configs, logger logger.AppLogger, services *service.Builder) Handlers {
	return Handlers{
		middleware: handler.NewMiddleware(logger),
		market:     handler.NewMarket(&handler.Options{Cfg: cfg, Logger: logger, Services: services}),
		shipment:   handler.NewShipment(&handler.Options{Cfg: cfg, Logger: logger, Services: services}),
	}
}
