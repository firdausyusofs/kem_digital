package inventory

import (
	"context"
	"firdausyusofs/kem_digital/internal/models"
)

type UseCase interface {
	GetProducts(ctx context.Context) ([]*models.Product, error)
	GetProductByID(ctx context.Context, id int64) (*models.Product, error)
}
