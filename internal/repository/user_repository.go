package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
)

type UserRepository struct {
	db *sql.DB // use pointer because dont want to always copy the db object
}

func NewUserRepository(db *sql.DB) *UserRepository {
	// take the address of the struct
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (
			id,
			name,
			email,
			password_hash,
			phone,
			role,
			is_active,
			created_at,
			updated_at
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			NOW(),
			NOW()
		)
		RETURNING created_at, updated_at
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		user.ID,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.Phone,
		user.Role,
		user.IsActive,
	).Scan(
		&user.CreatedAt, // take returning value form query and assign to user model
		&user.UpdatedAt,
	)
}

func (r *UserRepository) CreateTx(ctx context.Context, tx *sql.Tx, user *models.User) error {

	query := `
		INSERT INTO users (
			id,
			name,
			email,
			password_hash,
			phone,
			role,
			is_active,
			created_at,
			updated_at
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			NOW(),
			NOW()
		)
	`
	// use tx (transaction) to execute all 
	// query in transaction (ensure all query in transaction succeed)
	_, err := tx.ExecContext(
		ctx,
		query,
		user.ID,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.Phone,
		user.Role,
		user.IsActive,
	)

	return err

}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT
			id,
			name,
			email,
			password_hash,
			phone,
			role,
			is_active,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
	`
	
	var user models.User

	err := r.db.QueryRowContext(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Phone,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := `
		SELECT
			id,
			name,
			email,
			password_hash,
			phone,
			role,
			is_active,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
	`

	var user models.User

	// QueryRowContext return one row data (SELECT)
	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Phone,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdatePassword(
	ctx context.Context,
	userID uuid.UUID,
	passwordHash string,
) error {
	_, err := r.db.ExecContext(
		ctx,
		`
			UPDATE users
			SET password_hash = $1,
				updated_at = NOW()
			WHERE id = $2
		`, passwordHash, userID,
	)
	return err
}

func (r *UserRepository) UpdatePhone(
	ctx context.Context,
	userID uuid.UUID,
	newPhone string,
) (*models.User, error) {
	var user models.User

	query := `
		UPDATE users
		SET phone = $1,
		    updated_at = NOW()
		WHERE id = $2
		RETURNING
		    id,
		    name,
		    email,
		    phone,
		    role,
		    is_active,
		    created_at,
		    updated_at
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		newPhone,
		userID,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateName(
	ctx context.Context,
	userID uuid.UUID,
	name string,
) (*models.User, error) {
	var user models.User

	query := `
		UPDATE users
		SET name = $1,
		    updated_at = NOW()
		WHERE id = $2
		RETURNING
		    id,
		    name,
		    email,
		    phone,
		    role,
		    is_active,
		    created_at,
		    updated_at
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		name,
		userID,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

