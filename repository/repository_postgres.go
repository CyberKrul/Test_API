package repository

// Package repository implements the data access layer for the application,
// providing concrete implementations of the repository interface using a PostgreSQL database.
import (
	"TAPI/model"
	"context"
	"database/sql"
	"errors"
)

// PGModelRepo provides methods for interacting with the device data in the database.
type PGModelRepo struct {
	PGConn *sql.DB
}

// ErrNotFound is returned when a requested item is not found in the database.
var ErrNotFound = errors.New("repository: item not found")

// NewPGModelRepo creates a new PGModelRepo with the given database connection.
func NewPGModelRepo(pmr *sql.DB) *PGModelRepo {
	return &PGModelRepo{PGConn: pmr}
}

// CreateModel inserts a new device record into the database.
// It uses RETURNING to scan the inserted sno back into the model.
func (pg *PGModelRepo) CreateModel(ctx context.Context, m *model.ModelInstance) error {
	query := `
		INSERT INTO deviceDBTwo (
			sno,
			firmware_version,
			current_firmware_version,
			mesh_configuration,
			app_configuration,
			kc_configuration
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING sno`

	// Use QueryRowContext to make the database call cancellable.
	err := pg.PGConn.QueryRowContext(ctx, query,
		m.SNo,
		m.FirmwareVersion,
		m.CurrentFirmwareVersion,
		m.MeshConfig,
		m.AppConfig,
		m.KCConfig,
	).Scan(&m.SNo)

	return err
}

// UpdateModelbySNO toggles the mesh_configuration for a device identified by its SNo.
// It returns the entire updated record.
func (pg *PGModelRepo) UpdateModelbySNO(ctx context.Context, sno int) (*model.ModelInstance, error) {
	query := `
		UPDATE deviceDBTwo
		SET mesh_configuration = NOT mesh_configuration
		WHERE sno = $1
		RETURNING sno, firmware_version, current_firmware_version, mesh_configuration, app_configuration, kc_configuration`

	var m model.ModelInstance

	err := pg.PGConn.QueryRowContext(ctx, query, sno).Scan(
		&m.SNo,
		&m.FirmwareVersion,
		&m.CurrentFirmwareVersion,
		&m.MeshConfig,
		&m.AppConfig,
		&m.KCConfig,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &m, err
}

// GetModelbySNO retrieves a device record from the database by its SNo.
func (pg *PGModelRepo) GetModelbySNO(ctx context.Context, sno int) (*model.ModelInstance, error) {
	query := `
		SELECT
			sno, firmware_version, current_firmware_version, mesh_configuration, app_configuration, kc_configuration
		FROM deviceDBTwo
		WHERE sno = $1`

	var m model.ModelInstance

	err := pg.PGConn.QueryRowContext(ctx, query, sno).Scan(
		&m.SNo,
		&m.FirmwareVersion,
		&m.CurrentFirmwareVersion,
		&m.MeshConfig,
		&m.AppConfig,
		&m.KCConfig,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &m, err
}
