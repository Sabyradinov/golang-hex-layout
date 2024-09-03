package repository

import (
	"context"
	"github.com/Sabyradinov/golang-hex-layout/internal/adapter/storage/entity"
)

type Manager struct {
	Product  IProduct
	Shipment IShipment
}

type IProduct interface {
	GetById(ctx context.Context, id int64) (res entity.Product, err error)
	Create(ctx context.Context, req entity.Product) (err error)
	Update(ctx context.Context, req entity.Product) (err error)
	Delete(ctx context.Context, req entity.Product) (err error)
}

type IShipment interface {
	GetById(ctx context.Context, id int64) (res entity.Shipment, err error)
	Create(ctx context.Context, req entity.Shipment) (err error)
	Update(ctx context.Context, req entity.Shipment) (err error)
	Delete(ctx context.Context, req entity.Shipment) (err error)
}
