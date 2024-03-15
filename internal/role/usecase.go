package role

import (
	"context"
	"firdausyusofs/kem_digital/internal/models"
)

type UseCase interface {
	GetRoles(ctx context.Context) ([]*models.Role, error)
}
