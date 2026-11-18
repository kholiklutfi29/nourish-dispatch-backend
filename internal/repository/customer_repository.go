package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) CreateTx(ctx context.Context, tx *sql.Tx, customer *models.Customer) error {

	query := `
		INSERT INTO customers (
			user_id,
			default_address,
			latitude,
			longitude,
			created_at,
			updated_at
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			NOW(),
			NOW()
		)
	`
	_, err := tx.ExecContext(
		ctx,
		query,
		customer.UserID,
		customer.DefaultAddress,
		customer.Latitude,
		customer.Longitude,
	)

	return err

}

func (r *CustomerRepository) GetCustomerProfile(
	ctx context.Context,
	customerID uuid.UUID,
) (*models.CustomerDetail, error) {
	query := `
		SELECT 
			u.id,
			u.name,
			u.email,
			u.phone,
			u.role,
			c.default_address,
			c.latitude,
			c.longitude,
			c.created_at,
			c.updated_at
		FROM customers c
		INNER JOIN users u ON u.id = c.user_id
		WHERE c.user_id = $1
			AND u.role = 'CUSTOMER'
	`

	customer := &models.CustomerDetail{}

	err := r.db.QueryRowContext(
		ctx, query, customerID,
	).Scan(
		&customer.UserID,
		&customer.Name,
		&customer.Email,
		&customer.Phone,
		&customer.Role,
		&customer.DefaultAddress,
		&customer.Latitude,
		&customer.Longitude,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *CustomerRepository) GetCustomerAddress(
	ctx context.Context,
	customerID uuid.UUID,
) (*models.Customer, error) {
	query := `
		SELECT 
			user_id,
			default_address,
			latitude,
			longitude
		FROM customers
		WHERE user_id = $1
	`
	customer := &models.Customer{}

	err := r.db.QueryRowContext(
		ctx, query, customerID,
	).Scan(
		&customer.UserID,
		&customer.DefaultAddress,
		&customer.Latitude,
		&customer.Longitude,
	)

	if err != nil {
		return nil, err
	}

	 return customer, nil
}

func (r *CustomerRepository) UpdateCustomerAddress(
	ctx context.Context,
	customerID uuid.UUID,
	defaultAddress string,
	latitude float64,
	longitude float64,
) error {
	query := `
		UPDATE customers
		SET
			default_address = $1,
			latitude = $2,
			longitude = $3,
			updated_at = NOW()
		WHERE user_id = $4
	`

	result, err := r.db.ExecContext(
		ctx, query, defaultAddress, latitude, longitude, customerID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrCustomerNotFound
	}

	return nil

}
