package dto

import (
	"errors"
	"regexp"

	"github.com/fatjan/tutuplapak/internal/pkg/exceptions"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	Email    string
	Phone    string
	Password string
}

type AuthResponse struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
	Token string `json:"token"`
}

func (d *AuthRequest) ValidatePayloadAuth(isPhoneRegister bool) error {
	if d.Password == "" || !isValidPasswordLength(d.Password, 8, 32) {
		return exceptions.ErrorBadRequest
	}

	if isPhoneRegister {
		if d.Phone == "" || !isValidPhone(d.Phone) {
			return exceptions.ErrorBadRequest
		}
		return nil
	}

	if d.Email == "" || !isValidEmail(d.Email) {
		return exceptions.ErrorBadRequest
	}
	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidPhone(phone string) bool {
	re := regexp.MustCompile(`^\+[0-9]{10,}$`)
	return re.MatchString(phone)
}

func isValidPasswordLength(password string, minLength, maxLength int) bool {
	passwordLength := len(password)
	return passwordLength >= minLength && passwordLength <= maxLength
}

func (d *AuthRequest) HashPassword() error {
	resultHash, err := bcrypt.GenerateFromPassword([]byte(d.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error hashing")
	}
	d.Password = string(resultHash)

	return nil
}
