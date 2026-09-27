package repository

import (
	"context"
	"database/sql"

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