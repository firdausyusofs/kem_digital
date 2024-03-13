package usecase

import (
	"context"
	"firdausyusofs/kem_digital/config"
	"firdausyusofs/kem_digital/internal/models"
	"firdausyusofs/kem_digital/internal/supplier"
	"firdausyusofs/kem_digital/pkg/logger"
)

type supplierUseCase struct {
	cfg          *config.Config
	supplierRepo supplier.Repository
	logger       logger.Logger
}

func NewSupplierUseCase(cfg *config.Config, supplierRepo supplier.Repository, logger logger.Logger) supplier.UseCase {
	return &supplierUseCase{
		cfg:          cfg,
		supplierRepo: supplierRepo,
		logger:       logger,
	}
}

func (u *supplierUseCase) GetSuppliers(ctx context.Context) ([]*models.Supplier, error) {
	return u.supplierRepo.GetSuppliers(ctx)
}
