package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/dto"
	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
	"github.com/kholiklutfi29/nourish-dispatch/internal/repository"
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