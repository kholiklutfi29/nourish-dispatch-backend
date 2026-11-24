package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) CreateCart(
	ctx context.Context,
	cart *models.Cart,
) (*models.Cart, error) {
	query := `
		INSERT INTO carts (
			id,
			customer_id,
			restaurant_id,
			created_at,
			updated_at
		)
		VALUES (
			$1, $2, $3, NOW(), NOW()
		)
		RETURNING
			id,
			customer_id,
			restaurant_id,
			created_at,
			updated_at
	`

	createdCart := &models.Cart{}

	err := r.db.QueryRowContext(
		ctx,
		query,
		cart.ID,
		cart.CustomerID,
		cart.RestaurantID,
	).Scan(
		&createdCart.ID,
		&createdCart.CustomerID,
		&createdCart.RestaurantID,
		&createdCart.CreatedAt,
		&createdCart.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return createdCart, nil
}

func (r *CartRepository) GetCartByCustomer(
	ctx context.Context,
	customerID uuid.UUID,
) (*models.Cart, error) {
	query := `
		SELECT
			id,
			customer_id,
			restaurant_id,
			created_at,
			updated_at
		FROM carts
		WHERE customer_id = $1
	`
	cart := &models.Cart{}

	err := r.db.QueryRowContext(
		ctx, query, customerID,
	).Scan(
		&cart.ID,
		&cart.CustomerID,
		&cart.RestaurantID,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (r *CartRepository) DeleteCart(
	ctx context.Context,
	customerID uuid.UUID,
	cartID uuid.UUID,
) error {
	query := `
		DELETE FROM carts
		WHERE id = $1
		AND customer_id = $2
	`

	result, err := r.db.ExecContext(
		ctx, query, cartID, customerID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}