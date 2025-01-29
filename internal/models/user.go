package models

import "time"

type User struct {
	ID                int       `db:"id"`
	Email             string    `db:"email"`
	Phone             string    `db:"phone"`
	Password          string    `db:"password"`
	FileId            string    `db:"file_id"`
	FileUrl           string    `db:"file_url"`
	FileThumbnailUrl  string    `db:"file_thumbnail_url"`
	BankAccountName   string    `db:"bank_account_name"`
	BankAccountHolder string    `db:"bank_account_holder"`
	BankAccountNumber string    `db:"bank_account_number"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}
