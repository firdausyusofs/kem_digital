package repository

import (
	"context"
	"firdausyusofs/kem_digital/internal/auth"
	"firdausyusofs/kem_digital/internal/models"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.Repository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Register(ctx context.Context, user *models.User) (*models.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *authRepository) FindByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	found := &models.User{}
	result := r.db.Where("email = ?", user.Email).First(&found)
	if result.Error != nil {
		return nil, result.Error
	}

	return found, nil
}

func (r *authRepository) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	return nil, nil
}
