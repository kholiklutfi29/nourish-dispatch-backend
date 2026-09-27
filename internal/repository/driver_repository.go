package repository

import (
	"context"
	"database/sql"

	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
)

type DriverRepository struct {
	db *sql.DB
}

func NewDriverRepository(db *sql.DB) *DriverRepository {
	return &DriverRepository{
		db: db,
	}
}

func (r *DriverRepository) CreateTx(
	ctx context.Context,
	tx *sql.Tx,
	driver *models.Driver,
) error {

	query := `
		INSERT INTO drivers (
			user_id,
			vehicle_type,
			vehicle_plate,
			status,
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
		driver.UserID,
		driver.VehicleType,
		driver.VehiclePlate,
		driver.Status,
	)

	return err
}