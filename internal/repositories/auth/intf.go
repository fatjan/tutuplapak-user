package auth

import (
	"context"

	"github.com/fatjan/tutuplapak/internal/models"
)

type Repository interface {
	PostEmail(context.Context, *models.User) (int, error)
	PostPhone(context.Context, *models.User) (int, error)
}
