package role

import (
	"context"
	"firdausyusofs/kem_digital/internal/models"
)

type Repository interface {
	GetRoles(ctx context.Context) ([]*models.Role, error)
	CreateRole(ctx context.Context, role *models.Role) (*models.Role, error)
}
