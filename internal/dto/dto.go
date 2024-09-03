package dto

import (
	"github.com/Sabyradinov/golang-hex-layout/internal/adapter/storage/entity"
	"github.com/Sabyradinov/golang-hex-layout/internal/model"
)

// ToProductEntity converts ProductRequest to Product entity
func ToProductEntity(req model.ProductRequest) (res entity.Product) {
	return entity.Product(req)
}

// ToShipmentEntity converts ShipmentRequest to Shipment entity
func ToShipmentEntity(req model.ShipmentRequest) (res entity.Shipment) {
	return entity.Shipment(req)
}
