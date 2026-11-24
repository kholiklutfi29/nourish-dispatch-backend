package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
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

func (r *RestaurantRepository) UpdateRestaurantStatus(
	ctx context.Context,
	restaurantID uuid.UUID,
	ownerID uuid.UUID,
	isOpen bool,
) error {
	query := `
		UPDATE restaurants
		SET
			is_open = $1,
			updated_at = NOW()
		WHERE id = $2
			AND owner_id = $3
			AND is_active = true
	`
	result, err := r.db.ExecContext(
		ctx, query, isOpen, restaurantID, ownerID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrRestaurantNotFound
	}
	return nil
}

func (r *RestaurantRepository) GetMyRestaurants(
	ctx context.Context,
	ownerID uuid.UUID,
) ([]*models.Restaurant, error) {
	query := `
		SELECT 
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
		FROM restaurants
		WHERE owner_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	restaurants := []*models.Restaurant{}
	for rows.Next() {
		restaurant := &models.Restaurant{}
		err := rows.Scan(
			&restaurant.ID,
			&restaurant.OwnerID,
			&restaurant.Name,
			&restaurant.Description,
			&restaurant.Phone,
			&restaurant.Address,
			&restaurant.Latitude,
			&restaurant.Longitude,
			&restaurant.IsOpen,
			&restaurant.IsActive,
			&restaurant.CreatedAt,
			&restaurant.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		restaurants = append(restaurants, restaurant)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (r *RestaurantRepository) UpdateRestaurant(
	ctx context.Context,
	restaurant *models.Restaurant,
) error {
	query := `
		UPDATE restaurants
		SET
			name = $1,
			description = $2,
			phone = $3,
			address = $4,
			latitude = $5,
			longitude = $6,
			updated_at = NOW()
		WHERE id = $7
	`

	_, err := r.db.ExecContext(ctx, query,
		restaurant.Name,
		restaurant.Description,
		restaurant.Phone,
		restaurant.Address,
		restaurant.Latitude,
		restaurant.Longitude,
		restaurant.ID,
	)
	return err
}

func (r *RestaurantRepository) SoftDeleteRestaurant(
	ctx context.Context,
	restaurantID uuid.UUID,
	ownerID uuid.UUID,
) error {
	query := `
		UPDATE restaurants
		SET
			is_active = false,
			updated_at = NOW()
		WHERE id = $1
			AND owner_id = $2
			AND is_active = true
	`

	result, err := r.db.ExecContext(ctx, query, restaurantID, ownerID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrRestaurantNotFound
	}

	return nil
}

func (r *RestaurantRepository) GetRestaurantByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Restaurant, error) {
	query := `
		SELECT 
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
		FROM restaurants
		WHERE id = $1
	`

	restaurant := &models.Restaurant{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&restaurant.ID,
		&restaurant.OwnerID,
		&restaurant.Name,
		&restaurant.Description,
		&restaurant.Phone,
		&restaurant.Address,
		&restaurant.Latitude,
		&restaurant.Longitude,
		&restaurant.IsOpen,
		&restaurant.IsActive,
		&restaurant.CreatedAt,
		&restaurant.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return restaurant, nil
}

func (r *RestaurantRepository) GetAllRestaurants(
	ctx context.Context,
) ([]*models.Restaurant, error) {
	query := `
		SELECT 
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
		FROM restaurants
		WHERE is_active = true
		ORDER BY name ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	restaurants := []*models.Restaurant{}

	for rows.Next() {
		restaurant := &models.Restaurant{}

		err := rows.Scan(
			&restaurant.ID,
			&restaurant.OwnerID,
			&restaurant.Name,
			&restaurant.Description,
			&restaurant.Phone,
			&restaurant.Address,
			&restaurant.Latitude,
			&restaurant.Longitude,
			&restaurant.IsOpen,
			&restaurant.IsActive,
			&restaurant.CreatedAt,
			&restaurant.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		restaurants = append(restaurants, restaurant)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return restaurants, nil
}

func (r *RestaurantRepository) GetRestaurantIDByOwner(
    ctx context.Context,
    ownerID uuid.UUID,
) (uuid.UUID, error) {
    query := `
        SELECT id
        FROM restaurants
        WHERE owner_id = $1
    `

    var restaurantID uuid.UUID

    err := r.db.QueryRowContext(
        ctx,
        query,
        ownerID,
    ).Scan(&restaurantID)

    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return uuid.Nil, apperrors.ErrRestaurantNotFound
        }

        return uuid.Nil, err
    }

    return restaurantID, nil
}
