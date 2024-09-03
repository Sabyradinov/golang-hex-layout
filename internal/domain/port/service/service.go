package service

import (
	"context"
	"github.com/Sabyradinov/golang-hex-layout/internal/model"
)

type IMarket interface {
	GetProductById(ctx context.Context, Id int64) (res interface{}, err error)
	Buy(ctx context.Context, req model.ProductRequest) (err error)
	GetShipmentById(ctx context.Context, Id int64) (res interface{}, err error)
	DeliveryToClient(ctx context.Context, req model.ShipmentRequest) (err error)
}
