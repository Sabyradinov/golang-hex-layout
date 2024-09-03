package repo

import (
	"context"
	"github.com/Sabyradinov/golang-hex-layout/internal/adapter/storage/entity"
	"github.com/Sabyradinov/golang-hex-layout/internal/domain/port/repository"
	"gorm.io/gorm"
	"time"
)

// product info about store goods
type product struct {
	db *gorm.DB
}

// NewProduct constructor to create market repository instance
func NewProduct(db *gorm.DB) repository.IProduct {
	return &product{db: db}
}

// GetById function to get market by id with context
func (r *product) GetById(ctx context.Context, id int64) (res entity.Product, err error) {
	tx := r.db.WithContext(ctx)
	err = tx.First(&res, id).Error
	return
}

// Create function to insert to new row with context
func (r *product) Create(ctx context.Context, req entity.Product) error {
	//Create transaction with timeout
	ctxTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	tx := r.db.WithContext(ctxTimeout)
	err := tx.Create(&req).Error

	return err
}

// Update function to update row with context
func (r *product) Update(ctx context.Context, req entity.Product) (err error) {
	//Create transaction with timeout
	ctxTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	tx := r.db.WithContext(ctxTimeout)
	err = tx.Updates(&req).Error

	return
}

// Delete function to delete row with context
func (r *product) Delete(ctx context.Context, req entity.Product) (err error) {
	tx := r.db.WithContext(ctx)

	err = tx.Delete(&req).Error
	return
}
