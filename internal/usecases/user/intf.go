package user

import (
	"context"

	"github.com/fatjan/tutuplapak/internal/dto"
)

type UseCase interface {
	GetUser(context.Context, *dto.UserRequest) (*dto.UserPatchResponse, error)
	UpdateUser(context.Context, int, *dto.UserPatchRequest) (*dto.UserPatchResponse, error)
}
