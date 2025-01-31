package dto

import (
	"github.com/fatjan/tutuplapak/internal/pkg/exceptions"
)

type UserRequest struct {
	UserID int `json:"id"`
}

type UserPatchRequest struct {
	FileId            string `json:"fileId"`
	BankAccountName   string `json:"bankAccountName"`
	BankAccountHolder string `json:"bankAccountHolder"`
	BankAccountNumber string `json:"bankAccountNumber"`
}

type UserPatchResponse struct {
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	FileId            string `json:"fileId"`
	FileUri           string `json:"fileUri"`
	FileThumbnailUri  string `json:"fileThumbnailUri"`
	BankAccountName   string `json:"bankAccountName"`
	BankAccountHolder string `json:"bankAccountHolder"`
	BankAccountNumber string `json:"bankAccountNumber"`
}

func (d *UserPatchRequest) ValidatePayload() error {
	if d.BankAccountName == "" || d.BankAccountHolder == "" || d.BankAccountNumber == "" {
		return exceptions.ErrorBadRequest
	}
	if len(d.BankAccountHolder) <= 3 || len(d.BankAccountName) <= 3 || len(d.BankAccountNumber) <= 3 {
		return exceptions.ErrorBadRequest
	}
	return nil
}
