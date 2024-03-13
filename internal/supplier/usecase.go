package supplier

import (
	"context"
	"firdausyusofs/kem_digital/internal/models"
)

type UseCase interface {
	GetSuppliers(ctx context.Context) ([]*models.Supplier, error)
}
