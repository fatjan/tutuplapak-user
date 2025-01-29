package auth

import (
	"context"

	"github.com/fatjan/tutuplapak/internal/dto"
)

type UseCase interface {
	Register(context.Context, *dto.AuthRequest, bool) (*dto.AuthResponse, error)
}
