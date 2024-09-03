package repo

import (
	"context"
	"github.com/Sabyradinov/golang-hex-layout/internal/adapter/storage/entity"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/repository"
	"gorm.io/gorm"
	"time"
)

// shipment info about shipment in store
type shipment struct {
	db *gorm.DB
}

// NewShipment constructor to create shipment repository instance
func NewShipment(db *gorm.DB) repository.IShipment {
	return &shipment{db: db}
}

// GetById function to get shipment by id with context
func (r *shipment) GetById(ctx context.Context, id int64) (res entity.Shipment, err error) {
	tx := r.db.WithContext(ctx)
	err = tx.First(&res, id).Error
	return
}

// Create function to insert new row with context
func (r *shipment) Create(ctx context.Context, req entity.Shipment) error {
	//Create transaction with timeout
	ctxTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	tx := r.db.WithContext(ctxTimeout)
	err := tx.Create(&req).Error

	return err
}

// Update function to update row with context
func (r *shipment) Update(ctx context.Context, req entity.Shipment) (err error) {
	//Create transaction with timeout
	ctxTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	tx := r.db.WithContext(ctxTimeout)
	err = tx.Updates(&req).Error

	return
}

// Delete function to delete row with context
func (r *shipment) Delete(ctx context.Context, req entity.Shipment) (err error) {
	tx := r.db.WithContext(ctx)

	err = tx.Delete(&req).Error
	return
}
