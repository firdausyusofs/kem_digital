package usecase

import (
	"context"
	"firdausyusofs/kem_digital/config"
	"firdausyusofs/kem_digital/internal/inventory"
	"firdausyusofs/kem_digital/internal/models"
	"firdausyusofs/kem_digital/pkg/logger"
)

type inventoryUseCase struct {
	cfg           *config.Config
	inventoryRepo inventory.Repository
	logger        logger.Logger
}

func NewInventoryUseCase(cfg *config.Config, inventoryRepo inventory.Repository, logger logger.Logger) inventory.UseCase {
	return &inventoryUseCase{
		cfg:           cfg,
		inventoryRepo: inventoryRepo,
		logger:        logger,
	}
}

func (u *inventoryUseCase) GetProducts(ctx context.Context) ([]*models.Product, error) {
	return u.inventoryRepo.GetProducts(ctx)
}

func (u *inventoryUseCase) GetProductByID(ctx context.Context, id int64) (*models.Product, error) {
	return u.inventoryRepo.GetProductByID(ctx, id)
}

func (u *inventoryUseCase) CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	return u.inventoryRepo.CreateProduct(ctx, product)
}
