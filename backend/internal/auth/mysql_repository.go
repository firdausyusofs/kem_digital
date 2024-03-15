package auth

import (
	"context"
	"firdausyusofs/kem_digital/internal/models"
)

type Repository interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	FindByEmail(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByID(ctx context.Context, id int64) (*models.User, error)
}
