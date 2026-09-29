package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
	"github.com/kholiklutfi29/nourish-dispatch/internal/dto"
	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
	"github.com/kholiklutfi29/nourish-dispatch/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db             *sql.DB
	userRepository *repository.UserRepository
}

func NewUserService(
	db *sql.DB,
	userRepository *repository.UserRepository,
) *UserService {
	return &UserService{
		db:             db,
		userRepository: userRepository,
	}
}

func (s *UserService) ChangeUserName(
	ctx context.Context,
	req dto.UserChangeNameRequest,
	id string,
) (*models.User, error) {
	// ==========================================
	// 1. NORMALIZE INPUT
	// ==========================================
	name := strings.TrimSpace(req.NewName)

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	// ==========================================
	// 2. PARSE USER ID
	// ==========================================
	userID, err := uuid.Parse(id)
	
	if err != nil {
		return nil, errors.New("Invalid user id")
	}

	// ==========================================
	// 3. UPDATE USER
	// ==========================================
	return s.userRepository.UpdateName(ctx, userID, name)

}

func (s *UserService) ChangeUserPassword(
	ctx context.Context,
	req dto.UserChangePasswordRequest,
	id string,
) error {
	// ==========================================
	// 1. NIL CHECK & NORMALIZE INPUT
	// ==========================================
	if req.OldPassword == nil {
		return errors.New("old password is required")
	}

	if req.NewPassword == nil {
		return errors.New("new password is required")
	}

	// ==========================================
	// 2. NORMALIZE INPUT & CHECK 
	// ==========================================
	oldPassword := strings.TrimSpace(*req.OldPassword)
	newPassword := strings.TrimSpace(*req.NewPassword)

	if oldPassword == "" {
		return errors.New("old password cannot be empty")
	}

	if newPassword == "" {
		return errors.New("new password cannot be empty")
	}

	if oldPassword == newPassword {
		return errors.New("new password must be different from old password")
	}


	// ==========================================
	// 3. PARSE USER ID
	// ==========================================
	UserID, err := uuid.Parse(id)

	if err != nil {
		return  errors.New("Invalid user id")
	}

	// ==========================================
	// 4. GET EXISTING PASSWORD HASH & COMPARE
	// ==========================================
	user, err := s.userRepository.FindByID(ctx, UserID)

	if err != nil {
		return fmt.Errorf("%w: %v", apperrors.ErrInternal, err)
	}

	if err:= bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(oldPassword),
	); err != nil {
		return errors.New("old password is incorrect")
	}

	// ==========================================
	// 5. HASH NEW PASSWORD
	// ==========================================
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(newPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return fmt.Errorf("%w: %v", apperrors.ErrInternal, err)
	}

	
	// ==========================================
	// 5. UPADATE PASSWORD
	// ==========================================

	return s.userRepository.UpdatePassword(ctx, UserID, string(hashedPassword))
}


func (s *UserService) ChangeUserPhone(
	ctx context.Context,
	req dto.UserChangePhoneRequest,
	id string,
) (*models.User, error) {
	
	// ==========================================
	// 1. NIL CHECK & NORMALIZE INPUT
	// ==========================================
	if req.NewPhone == nil {
		return nil, errors.New("phone is required")
	}

	// ==========================================
	// 2. NORMALIZE INPUT
	// ==========================================
	phone := strings.TrimSpace(*req.NewPhone)

	if phone == "" {
		return nil, errors.New("phone cannot be empty")
	}

	// ==========================================
	// 3. PARSE USER ID
	// ==========================================
	userID, err := uuid.Parse(id)
	
	if err != nil {
		return nil, errors.New("Invalid user id")
	}

	// ==========================================
	// 4. UPDATE USER
	// ==========================================
	return s.userRepository.UpdatePhone(ctx, userID, phone)
}