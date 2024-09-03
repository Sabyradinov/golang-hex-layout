package service

import (
	"github.com/Sabyradinov/golang-hex-layout/config"
	"github.com/Sabyradinov/golang-hex-layout/internal/adapter/httpUtilities"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/repository"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/service"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/service/market"
)

type Builder struct {
	Market service.IMarket
}

func Init(cfg *config.Configs, db *repository.Manager) *Builder {
	return &Builder{
		Market: market.New(cfg, db, httpUtilities.New(cfg.MarketUrls.BaseAddress.HttpTimeout)),
	}
}
