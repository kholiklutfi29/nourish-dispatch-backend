package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
	"github.com/kholiklutfi29/nourish-dispatch/internal/dto"
	"github.com/kholiklutfi29/nourish-dispatch/internal/helper"
	"github.com/kholiklutfi29/nourish-dispatch/internal/jwt"
	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
	"github.com/kholiklutfi29/nourish-dispatch/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db             *sql.DB
	userRepository *repository.UserRepository
	profileService *ProfileService
	jwtService     *jwt.JWTService
}

func NewAuthService(
	db *sql.DB,
	userRepository *repository.UserRepository,
	profileService *ProfileService,
	jwtService *jwt.JWTService,
) *AuthService {
	return &AuthService{
		db:             db,
		userRepository: userRepository,
		profileService: profileService,
		jwtService:     jwtService,
	}
}

func (s *AuthService) Register(
	ctx context.Context,
	req dto.RegisterRequest,
) (*dto.RegisterResponse, error) {

	// ==========================================
	// 1. NORMALIZE INPUT
	// ==========================================

	email := strings.ToLower(
		strings.TrimSpace(req.Email),
	)

	name := strings.TrimSpace(req.Name)

	// ==========================================
	// 2. VALIDATE ROLE
	// ==========================================

	if !helper.IsValidRole(req.Role) {
		return nil, apperrors.ErrInvalidRole
	}

	// ==========================================
	// 3. CHECK EMAIL
	// ==========================================

	existingUser, err := s.userRepository.FindByEmail(
		ctx,
		email,
	)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New(
			"failed to check existing user",
		)
	}

	if existingUser != nil {
		return nil, apperrors.ErrEmailAlreadyRegistered
	}

	// ==========================================
	// 4. HASH PASSWORD
	// ==========================================

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, errors.New(
			"failed to hash password",
		)
	}

	// ==========================================
	// 5. CREATE USER MODEL
	// ==========================================

	user := &models.User{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: string(passwordHash),
		Phone:        nil,
		Role:         req.Role,
		IsActive:     true,
	}

	if phone := strings.TrimSpace(req.Phone); phone != "" {
		user.Phone = &phone
	}

	// ==========================================
	// 6. BEGIN TRANSACTION
	// ==========================================

	tx, err := s.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, errors.New(
			"failed to begin transaction",
		)
	}

	// ==========================================
	// 7. AUTOMATIC ROLLBACK
	// ==========================================

	defer tx.Rollback()

	// ==========================================
	// 8. CREATE USER (use transaction)
	// ==========================================

	err = s.userRepository.CreateTx(
		ctx,
		tx,
		user,
	)

	if err != nil {
		return nil, errors.New(
			"failed to create user",
		)
	}

	// ==========================================
	// 9. CREATE ROLE PROFILE (use transaction)
	// ==========================================

	err = s.profileService.CreateForRoleTx(
		ctx,
		tx,
		user,
	)

	if err != nil {
		return nil, err
	}

	// ==========================================
	// 10. COMMIT TRANSACTION
	// ==========================================

	if err := tx.Commit(); err != nil {
		return nil, errors.New(
			"failed to commit transaction",
		)
	}

	// ==========================================
	// 11. RETURN RESPONSE
	// ==========================================

	return &dto.RegisterResponse{
		User: dto.UserResponse{
			ID:       user.ID.String(),
			Name:     user.Name,
			Email:    user.Email,
			Phone:    user.Phone,
			Role:     user.Role,
			IsActive: user.IsActive,
		},
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))

	// cari user berdasarkan email
	user, err := s.userRepository.FindByEmail(ctx, email)

	if err != nil {
		return nil, apperrors.ErrInvalidCredentials
	}

	// check account
	if !user.IsActive {
		return nil, apperrors.ErrAccountInactive
	}

	// compare password
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, apperrors.ErrInvalidCredentials
	}

	// generate JWT
	token, err := s.jwtService.GenerateJWTToken(user)

	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// response
	return &dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:       user.ID.String(),
			Name:     user.Name,
			Email:    user.Email,
			Phone:    user.Phone,
			Role:     user.Role,
			IsActive: user.IsActive,
		},
	}, nil
}
