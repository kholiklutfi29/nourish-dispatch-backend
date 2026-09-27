package repository

import (
	"context"
	"database/sql"

	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
)

type RestaurantRepository struct {
	db *sql.DB
}

func NewRestaurantRepository(db *sql.DB) *RestaurantRepository {
	return &RestaurantRepository{
		db: db,
	}
}

func (r *RestaurantRepository) CreateTx(
	ctx context.Context,
	tx *sql.Tx,
	restaurant *models.Restaurant,
) error {

	query := `
		INSERT INTO restaurants (
			id,
			owner_id,
			name,
			description,
			phone,
			address,
			latitude,
			longitude,
			is_open,
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
			$8,
			$9,
			$10,
			NOW(),
			NOW()
		)
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		restaurant.ID,
		restaurant.OwnerID,
		restaurant.Name,
		restaurant.Description,
		restaurant.Phone,
		restaurant.Address,
		restaurant.Latitude,
		restaurant.Longitude,
		restaurant.IsOpen,
		restaurant.IsActive,
	)

	return err
}