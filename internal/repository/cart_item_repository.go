package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
)

type CartItemRepository struct {
	db *sql.DB
}

func NewCartItemRepository(db *sql.DB) *CartItemRepository {
	return &CartItemRepository{
		db: db,
	}
} 

func (r *CartItemRepository) CreateCartItem(
	ctx context.Context,
	cartItem *models.CartItem,
) error {
	query := `
		INSERT INTO cart_items (
			id,
			cart_id,
			menu_id,
			quantity,
			created_at,
			updated_at
		)
		VALUES (
			$1, $2, $3, $4, NOW(), NOW()
		)
	`
	_, err := r.db.ExecContext(
		ctx, query, cartItem.ID, cartItem.CartID, cartItem.MenuID, cartItem.Quantity,
	)

	return err
}

func (r *CartItemRepository) GetCartItems(
	ctx context.Context,
	cartID uuid.UUID,
) ([]*models.CartItem, error) {
	query := `
		SELECT 
			id,
			cart_id,
			menu_id,
			quantity,
			created_at,
			updated_at
		FROM cart_items
		WHERE cart_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(
		ctx, query, cartID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cartItems := []*models.CartItem{}
	for rows.Next() {
		cartItem := &models.CartItem{}
		err := rows.Scan(
			&cartItem.ID,
			&cartItem.CartID,
			&cartItem.MenuID,
			&cartItem.Quantity,
			&cartItem.CreatedAt,
			&cartItem.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (r *CartItemRepository) AddCartItem(
	ctx context.Context,
	cartItem *models.CartItem,
) error {
	query := `
		INSERT INTO cart_items (
			id,
			cart_id,
			menu_id,
			quantity
		)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (cart_id, menu_id)
		DO UPDATE SET
			quantity = cart_items.quantity + EXCLUDED.quantity,
			updated_at = NOW()
	`
	_, err := r.db.ExecContext(
        ctx,
        query,
        cartItem.ID,
        cartItem.CartID,
        cartItem.MenuID,
        cartItem.Quantity,
    )
    return err
}

// delete one
func (r *CartItemRepository) DeleteCartItem(
	ctx context.Context,
	cartItemID uuid.UUID,
	cartID uuid.UUID,
) error  {
	query := `
		DELETE FROM cart_items
		WHERE id = $1
		AND cart_id = $2
	`

	result, err := r.db.ExecContext(
		ctx, query, cartItemID, cartID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrCartItemNotFound
	}

	return nil
}

func (r *CartItemRepository) ClearCartItems(
    ctx context.Context,
    cartID uuid.UUID,
) error {
    query := `
        DELETE FROM cart_items
        WHERE cart_id = $1
    `

    _, err := r.db.ExecContext(
        ctx,
        query,
        cartID,
    )
	
    if err != nil {
        return err
    }

    return nil
}