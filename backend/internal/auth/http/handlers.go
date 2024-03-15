package http

import (
	"firdausyusofs/kem_digital/config"
	"firdausyusofs/kem_digital/internal/auth"
	"firdausyusofs/kem_digital/internal/models"
	"firdausyusofs/kem_digital/pkg/api_response"
	"firdausyusofs/kem_digital/pkg/logger"
	"firdausyusofs/kem_digital/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	cfg    *config.Config
	authUC auth.UseCase
	logger logger.Logger
}

func NewAuthHandler(cfg *config.Config, authUC auth.UseCase, logger logger.Logger) auth.Handlers {
	return &authHandler{
		cfg:    cfg,
		authUC: authUC,
		logger: logger,
	}
}

func (h *authHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := &models.User{}
		if err := utils.ReadBody(c, user); err != nil {
			return c.JSON(api_response.MakeErrorResponse(http.StatusBadRequest, "Invalid request body", nil))
		}

		createdUser, err := h.authUC.Register(c.Request().Context(), user)
		if err != nil {
			errMsg := err.Error()
			return c.JSON(api_response.MakeErrorResponse(http.StatusInternalServerError, "Failed to register user", &errMsg))
		}

		return c.JSON(api_response.MakeSuccessResponse(http.StatusCreated, "Successfully registered user", &createdUser))
	}
}

func (h *authHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		type Login struct {
			Email    string `json:"email" db:"email" validate:"required,email"`
			Password string `json:"password" db:"password" validate:"required,gte=6"`
		}
		loginReq := &Login{}
		if err := utils.ReadBody(c, loginReq); err != nil {
			h.logger.Errorf("Failed to bind user: %v", err)
			return c.JSON(api_response.MakeErrorResponse(http.StatusBadRequest, "Invalid request body", nil))
		}

		loggedInUser, err := h.authUC.Login(c.Request().Context(), &models.User{
			Email:    loginReq.Email,
			Password: loginReq.Password,
		})
		if err != nil {
			errMsg := err.Error()
			return c.JSON(api_response.MakeErrorResponse(http.StatusUnauthorized, "Failed to login", &errMsg))
		}

		return c.JSON(api_response.MakeSuccessResponse(http.StatusOK, "Successfully logged in", &loggedInUser))
	}
}
