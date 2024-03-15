package supplier

import (
	"context"
	"firdausyusofs/kem_digital/internal/models"
)

type Repository interface {
	GetSuppliers(ctx context.Context) ([]*models.Supplier, error)
	CreateSupplier(ctx context.Context, supplier *models.Supplier) (*models.Supplier, error)
}
