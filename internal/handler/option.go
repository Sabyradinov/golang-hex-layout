package handler

import (
	"github.com/Sabyradinov/golang-hex-layout/config"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/logger"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/service"
)

type Options struct {
	Services *service.Builder
	Logger   logger.AppLogger
	Cfg      *config.Configs
}
