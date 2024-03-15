package repository

import (
	"context"
	"firdausyusofs/kem_digital/internal/models"
	"firdausyusofs/kem_digital/internal/supplier"

	"gorm.io/gorm"
)

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) supplier.Repository {
	return &supplierRepository{
		db: db,
	}
}

func (r *supplierRepository) GetSuppliers(ctx context.Context) ([]*models.Supplier, error) {
	suppliers := make([]*models.Supplier, 0)
	result := r.db.Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}

	return suppliers, nil
}

func (r *supplierRepository) CreateSupplier(ctx context.Context, supplier *models.Supplier) (*models.Supplier, error) {
	result := r.db.Create(&supplier)
	if result.Error != nil {
		return nil, result.Error
	}

	return supplier, nil
}
