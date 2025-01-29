package auth

import (
	"context"

	"github.com/fatjan/tutuplapak/internal/config"
	"github.com/fatjan/tutuplapak/internal/dto"
	"github.com/fatjan/tutuplapak/internal/models"
	"github.com/fatjan/tutuplapak/internal/pkg/jwt_helper"
	"github.com/fatjan/tutuplapak/internal/repositories/auth"
)

type useCase struct {
	authRepository auth.Repository
	cfg            *config.Config
}

func NewUseCase(authRepository auth.Repository, cfg *config.Config) UseCase {
	return &useCase{
		authRepository: authRepository,
		cfg:            cfg,
	}
}

func (uc *useCase) Register(ctx context.Context, authRequest *dto.AuthRequest, isRegisterPhone bool) (*dto.AuthResponse, error) {
	err := authRequest.ValidatePayloadAuth(isRegisterPhone)
	if err != nil {
		return nil, err
	}

	err = authRequest.HashPassword()
	if err != nil {
		return nil, err
	}

	newAuth := &models.User{
		Email:    authRequest.Email,
		Phone:    authRequest.Phone,
		Password: authRequest.Password,
	}

	var id int
	if isRegisterPhone {
		id, err = uc.authRepository.PostPhone(ctx, newAuth)
	} else {
		id, err = uc.authRepository.PostEmail(ctx, newAuth)
	}
	if err != nil {
		return nil, err
	}

	token, err := jwt_helper.SignJwt(uc.cfg.JwtKey, id)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Email: authRequest.Email,
		Phone: authRequest.Phone,
		Token: token,
	}, nil
}
