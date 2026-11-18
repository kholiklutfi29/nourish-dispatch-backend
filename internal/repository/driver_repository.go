package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/kholiklutfi29/nourish-dispatch/internal/apperrors"
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

func (r *DriverRepository) GetDriverProfile(
	ctx context.Context,
	driverID uuid.UUID,
) (*models.Driver, error) {
	query := `
		SELECT 
			user_id,
			vehicle_type,
			vehicle_plate,
			status,
			created_at,
			updated_at,
		FROM drivers
		WHERE user_id = $1
	`
	driver := &models.Driver{}
	err := r.db.QueryRowContext(ctx, query, driverID).Scan(
		&driver.UserID,
		&driver.VehicleType,
		&driver.VehiclePlate,
		&driver.Status,
		&driver.CreatedAt,
		&driver.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return	driver, nil

}

func (r *DriverRepository) UpdateDriverProfile(
	ctx context.Context,
	driverID uuid.UUID,
	vehicleType string,
	vehiclePlate string,
) error {
	query := `
		UPDATE drivers
		SET 
			vehicle_type = $1,
			vehicle_type = $2,
			updated_at = NOW()
		WHERE user_id = $3;
	`

	result, err := r.db.ExecContext(
		ctx, query, vehicleType, vehiclePlate, driverID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrDriverNotFound
	}

	return nil
}

func (r *DriverRepository) UpdateDriverStatus(
	ctx context.Context,
	driverID uuid.UUID,
	status models.DriverStatus,
) error {
	query := `
		UPDATE drivers
		SET
			status = $2
			update_at = NOW()
		WHERE user_id = $1
	`

	result, err := r.db.ExecContext(
		ctx, query, driverID, status,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrDriverNotFound
	}

	return nil
}

func (r *DriverRepository) GetAllDrivers(
	ctx context.Context,
) ([]*models.Driver, error) {
	query := `
		SELECT
			user_id,
			vehicle_type,
			vehicle_plate,
			status,
			created_at,
			updated_at
		FROM drivers
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	drivers := []*models.Driver{}

	for rows.Next() {
		driver := &models.Driver{}

		err := rows.Scan(
			&driver.UserID,
			&driver.VehicleType,
			&driver.VehiclePlate,
			&driver.Status,
			&driver.CreatedAt,
			&driver.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		drivers = append(drivers, driver)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return drivers, nil
}

func (r *DriverRepository) GetDriverByStatus(
	ctx context.Context,
	status models.DriverStatus,
) ([]*models.Driver, error) {
	query := `
		SELECT 
			user_id,
			vehicle_type,
			vehicle_plate,
			status,
			created_at,
			updated_at,
		FROM drivers
		WHERE status = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(
		ctx, query, status,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// make instance struct
	drivers := []*models.Driver{} // slice of pointer that refer to models.driver

	for rows.Next() {

		// make instance struct and produce to the struct in same time
		driver := &models.Driver{} // address operator, to get address

		err := rows.Scan(
			&driver.UserID,
			&driver.VehicleType,
			&driver.VehiclePlate,
			&driver.Status,
			&driver.CreatedAt,
			&driver.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		drivers = append(drivers, driver)
	}

	return drivers, nil
}

func (r *DriverRepository) GetDriverDetails(
	ctx context.Context,
	driverID uuid.UUID,
) (*models.DriverDetail, error) {
	query := `
		SELECT
            u.id,
            u.name,
            u.email,
            u.phone,
            u.is_active,
            d.vehicle_type,
            d.vehicle_plate,
            d.status,
            u.created_at,
            d.updated_at
        FROM drivers d
        INNER JOIN users u ON u.id = d.user_id
        WHERE d.user_id = $1
          AND u.role = 'DRIVER'
	`
	driver := &models.DriverDetail{}

	err := r.db.QueryRowContext(
		ctx, query, driverID,
	).Scan(
		&driver.UserID,
        &driver.Name,
        &driver.Email,
        &driver.Phone,
        &driver.IsActive,
        &driver.VehicleType,
        &driver.VehiclePlate,
        &driver.Status,
        &driver.CreatedAt,
        &driver.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return driver, nil
}