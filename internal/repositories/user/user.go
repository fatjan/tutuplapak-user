package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/fatjan/tutuplapak/internal/dto"
	"github.com/fatjan/tutuplapak/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const (
	PG_DUPLICATE_ERROR = "23505"
)

type repository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetUser(ctx context.Context, id int) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, phone, file_id, file_url, file_thumbnail_url, bank_account_name, bank_account_holder, bank_account_number 
		FROM users 
		WHERE id = $1;
	`

	err := r.db.GetContext(ctx, user, query, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) Update(ctx context.Context, userID int, request *dto.UserPatchRequest) error {
	baseQuery := `UPDATE users SET `
	var setClauses []string
	var args []interface{}
	var argIndex int = 1

	if request != nil {
		if request.FileId != "" {
			setClauses = append(setClauses, fmt.Sprintf(`file_id = $%d`, argIndex))
			args = append(args, *&request.FileId)
			argIndex++
		}
		if request.BankAccountName != "" {
			setClauses = append(setClauses, fmt.Sprintf(`bank_account_name = $%d`, argIndex))
			args = append(args, *&request.FileId)
			argIndex++
		}
		if request.BankAccountHolder != "" {
			setClauses = append(setClauses, fmt.Sprintf(`bank_account_holder = $%d`, argIndex))
			args = append(args, *&request.FileId)
			argIndex++
		}
		if request.BankAccountNumber != "" {
			setClauses = append(setClauses, fmt.Sprintf(`bank_account_number = $%d`, argIndex))
			args = append(args, *&request.FileId)
			argIndex++
		}
	}

	if len(setClauses) == 0 {
		return errors.New("no fields to update")
	}

	baseQuery += strings.Join(setClauses, ", ")
	baseQuery += fmt.Sprintf(` WHERE id = $%d`, argIndex)
	args = append(args, userID)

	result, err := r.db.ExecContext(ctx, baseQuery, args...)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == PG_DUPLICATE_ERROR {
			return fmt.Errorf("duplicate email")
		}

		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("error query")
		return err
	}
	if rowsAffected == 0 {
		log.Println("failed update user")
		return errors.New("update query failed")
	}

	return nil
}
