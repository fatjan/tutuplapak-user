package user

import (
	"context"
	"fmt"
	"log"

	"github.com/fatjan/tutuplapak/internal/dto"
	"github.com/fatjan/tutuplapak/internal/repositories/user"
)

type useCase struct {
	userRepository user.Repository
}

func NewUseCase(userRepository user.Repository) UseCase {
	return &useCase{userRepository: userRepository}
}

func (u *useCase) GetUser(ctx context.Context, userRequest *dto.UserRequest) (*dto.UserPatchResponse, error) {
	user, err := u.userRepository.GetUser(ctx, userRequest.UserID)
	if err != nil {
		return nil, err
	}

	return &dto.UserPatchResponse{
		Email:             user.Email,
		Phone:             user.Phone,
		FileId:            user.FileId,
		FileUri:           user.FileUrl,
		FileThumbnailUri:  user.FileThumbnailUrl,
		BankAccountName:   user.BankAccountName,
		BankAccountHolder: user.BankAccountHolder,
		BankAccountNumber: user.BankAccountNumber,
	}, nil
}

func (u *useCase) UpdateUser(ctx context.Context, userID int, request *dto.UserPatchRequest) (*dto.UserPatchResponse, error) {
	err := request.ValidatePayload()
	if err != nil {
		return nil, err
	}

	// Get existing user
	user, err := u.userRepository.GetUser(ctx, userID)
	if err != nil {
		log.Println(fmt.Errorf("failed to get user: %w", err))
		return nil, err
	}

	// Update user in repository
	if err = u.userRepository.Update(ctx, userID, request); err != nil {
		log.Println(fmt.Errorf("failed to update user: %w", err))
		return nil, err
	}

	// Return new user data
	return &dto.UserPatchResponse{
		Email:             user.Email,
		Phone:             user.Phone,
		FileId:            user.FileId,
		FileUri:           user.FileUrl,
		FileThumbnailUri:  user.FileThumbnailUrl,
		BankAccountName:   user.BankAccountName,
		BankAccountHolder: user.BankAccountHolder,
		BankAccountNumber: user.BankAccountNumber,
	}, nil
}
