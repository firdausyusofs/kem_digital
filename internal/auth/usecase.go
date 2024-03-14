package auth

import (
	"context"
	"firdausyusofs/kem_digital/internal/models"
)

type UseCase interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Login(ctx context.Context, user *models.User) (*models.User, error)
}
