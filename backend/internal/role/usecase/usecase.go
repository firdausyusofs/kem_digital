package usecase

import (
	"context"
	"firdausyusofs/kem_digital/config"
	"firdausyusofs/kem_digital/internal/models"
	"firdausyusofs/kem_digital/internal/role"
	"firdausyusofs/kem_digital/pkg/logger"
)

type roleUseCase struct {
	cfg      *config.Config
	roleRepo role.Repository
	logger   logger.Logger
}

func NewRoleUseCase(cfg *config.Config, roleRepo role.Repository, logger logger.Logger) role.UseCase {
	return &roleUseCase{
		cfg:      cfg,
		roleRepo: roleRepo,
		logger:   logger,
	}
}

func (u *roleUseCase) GetRoles(ctx context.Context) ([]*models.Role, error) {
	roles, err := u.roleRepo.GetRoles(ctx)
	if err != nil {
		u.logger.Error(err)
		return nil, err
	}

	return roles, nil
}
