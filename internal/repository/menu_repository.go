package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
)

type MenuRepository struct {
	db *sql.DB
}

func NewMenuRepository(db *sql.DB) *MenuRepository {
	return &MenuRepository{
		db: db,
	}
}

func (r *MenuRepository) CreateMenu(
	ctx context.Context,
	menu *models.Menu,
) error {
	query := `
		INSERT INTO menus (
			id,
			restaurant_id,
			name,
			description,
			price,
			image_url,
			is_available,
			created_at,
			updated_at
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, NOW(), NOW()
		)
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		menu.ID,
		menu.RestaurantID,
		menu.Name,
		menu.Description,
		menu.Price,
		menu.ImageURL,
		menu.IsAvailable,
	)

	return err
}

func (r *MenuRepository) GetMyMenu(
	ctx context.Context,
	restaurantID uuid.UUID,
) ([]*models.Menu, error) {
	query := `
		SELECT 
			id,
			restaurant_id,
			name,
			description,
			price,
			image_url,
			is_available,
			created_at,
			updated_at
		FROM menus
		WHERE restaurant_id = $1
			AND deleted_at IS NULL
			AND is_available = TRUE
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(
		ctx, query, restaurantID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	menus := []*models.Menu{}

	for rows.Next() {
		menu := &models.Menu{}

		err := rows.Scan(
			&menu.ID,
			&menu.RestaurantID,
			&menu.Name,
			&menu.Description,
			&menu.Price,
			&menu.ImageURL,
			&menu.IsAvailable,
			&menu.CreatedAt,
			&menu.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		menus = append(menus, menu)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return menus, nil
}

func (r *MenuRepository) GetMenuDetail(
	ctx context.Context,
	menuID uuid.UUID,
) (*models.Menu, error) {
	query := `
		SELECT
			id,
			restaurant_id,
			name,
			description,
			price,
			image_url,
			is_available
		FROM menus
		WHERE id = $1
			AND deleted_at IS NULL
	`
	menu := &models.Menu{}

	err := r.db.QueryRowContext(
		ctx, query, menuID,
	).Scan(
		&menu.ID,
		&menu.RestaurantID,
		&menu.Name,
		&menu.Description,
		&menu.Price,
		&menu.ImageURL,
		&menu.IsAvailable,
	)

	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (r *MenuRepository) UpdateMenuInfo(
	ctx context.Context,
	restaurantID uuid.UUID,
	menuID uuid.UUID,
	menuName string,
	menuDescription string,
) error {
	query := `
		UPDATE menus
		SET 
			name = $1,
			description = $2,
			updated_at = NOW()
		WHERE id = $3
			AND restaurant_id = $4
			AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(
		ctx, query, menuName, menuDescription, menuID, restaurantID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrMenuNotFound
	}

	return nil
}

func (r *MenuRepository) UpdateMenuPrice(
	ctx context.Context,
	restaurantID uuid.UUID,
	menuID uuid.UUID,
	price float64,
) error {
	query := `
		UPDATE menus
		SET 
			price = $1,
			updated_at = NOW()
		WHERE id = $2
			AND restaurant_id = $3
			AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(
		ctx, query, price, menuID, restaurantID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrMenuNotFound
	}

	return nil
}

func (r *MenuRepository) UpdateMenuAvailability(
	ctx context.Context,
	restaurantID uuid.UUID,
	menuID uuid.UUID,
	isAvailable bool,
) error {
	query := `
		UPDATE menus
		SET 
			is_available = $1,
			updated_at = NOW()
		WHERE id = $2
			AND restaurant_id = $3
			AND deleted_at IS NULL
	`
	result, err := r.db.ExecContext(
		ctx, query, isAvailable, menuID, restaurantID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrMenuNotFound
	}

	return nil
}

func (r *MenuRepository) SoftDeleteMenu(
	ctx context.Context,
	restaurantID uuid.UUID,
	menuID uuid.UUID,
) error {
	query := `
		UPDATE menus
		SET
			deleted_at = NOW(),
			updated_at = NOW()
		WHERE id = $1
			AND restaurant_id = $2
			AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(
		ctx, query, menuID, restaurantID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrMenuNotFound
	}

	return nil
}

func (r *MenuRepository) GetRestaurantMenus(
	ctx context.Context,
	restaurantID uuid.UUID,
) ([]*models.Menu, error) {
	query := `
		SELECT 
			id,
			restaurant_id,
			name,
			description,
			price,
			image_url,
			is_available
		FROM menus
		WHERE restaurant_id = $1
			AND deleted_at IS NULL
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(
		ctx, query, restaurantID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	menus := []*models.Menu{}

	for rows.Next() {
		menu := &models.Menu{}

		err := rows.Scan(
			&menu.ID,
			&menu.RestaurantID,
			&menu.Name,
			&menu.Description,
			&menu.Price,
			&menu.ImageURL,
			&menu.IsAvailable,
		)

		if err != nil {
			return nil, err
		}

		menus = append(menus, menu)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return menus, nil
}
