package auth

import (
	"context"
	"time"

	"github.com/fatjan/tutuplapak/internal/models"
	"github.com/fatjan/tutuplapak/internal/pkg/exceptions"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type repository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) PostEmail(ctx context.Context, user *models.User) (int, error) {
	var newID int
	now := time.Now()

	query := `
			INSERT INTO users (email, password, created_at, updated_at)
			VALUES ($1, $2, $3, $4)
			RETURNING id`

	err := r.db.QueryRowContext(ctx, query,
		user.Email,
		user.Password,
		now,
		now,
	).Scan(&newID)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == pq.ErrorCode("23505") {
				return 0, exceptions.ErrConflict
			}
		}
		return 0, err
	}

	return newID, nil
}

func (r *repository) PostPhone(ctx context.Context, user *models.User) (int, error) {
	var newID int
	now := time.Now()

	query := `
			INSERT INTO users (phone, password, created_at, updated_at)
			VALUES ($1, $2, $3, $4)
			RETURNING id`

	err := r.db.QueryRowContext(ctx, query,
		user.Phone,
		user.Password,
		now,
		now,
	).Scan(&newID)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == pq.ErrorCode("23505") {
				return 0, exceptions.ErrConflict
			}
		}
		return 0, err
	}

	return newID, nil
}
