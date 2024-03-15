package repository

import (
	"context"
	"firdausyusofs/kem_digital/internal/models"
	"firdausyusofs/kem_digital/internal/role"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) role.Repository {
	return &roleRepository{db}
}

func (r *roleRepository) GetRoles(ctx context.Context) ([]*models.Role, error) {
	roles := make([]*models.Role, 0)
	result := r.db.WithContext(ctx).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}

func (r *roleRepository) CreateRole(ctx context.Context, role *models.Role) (*models.Role, error) {
	result := r.db.WithContext(ctx).Create(&role)
	if result.Error != nil {
		return nil, result.Error
	}

	return role, nil
}
