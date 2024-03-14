package usecase

import (
	"context"
	"errors"
	"firdausyusofs/kem_digital/config"
	"firdausyusofs/kem_digital/internal/auth"
	"firdausyusofs/kem_digital/internal/models"
	"firdausyusofs/kem_digital/pkg/logger"
	"firdausyusofs/kem_digital/pkg/utils"
)

type authUseCase struct {
	cfg      *config.Config
	authRepo auth.Repository
	logger   logger.Logger
}

func NewAuthUseCase(cfg *config.Config, authRepo auth.Repository, logger logger.Logger) auth.UseCase {
	return &authUseCase{
		cfg:      cfg,
		authRepo: authRepo,
		logger:   logger,
	}
}

func (u *authUseCase) Register(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	existingUser, err := u.authRepo.FindByEmail(ctx, user)
	if existingUser != nil || err == nil {
		u.logger.Error(err)
		return nil, errors.New("User already exists")
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	createdUser, err := u.authRepo.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	createdUser.ClearPassword()

	token, err := utils.GenerateToken(existingUser, u.cfg)
	if err != nil {
		return nil, err
	}

	return &models.UserWithToken{
		User:  createdUser,
		Token: token,
	}, nil
}

func (u *authUseCase) Login(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	existingUser, err := u.authRepo.FindByEmail(ctx, user)
	if err != nil {
		return nil, err
	}
	if err = existingUser.ComparePassword(user.Password); err != nil {
		return nil, err
	}

	existingUser.ClearPassword()

	token, err := utils.GenerateToken(existingUser, u.cfg)
	if err != nil {
		return nil, err
	}

	return &models.UserWithToken{
		User:  existingUser,
		Token: token,
	}, nil
}
