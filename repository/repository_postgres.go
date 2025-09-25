package repository

// This package will implement the SQL that will fulfill the contract's demands
import (
	"TAPI/model"
	"database/sql"
	"errors"
)

// PostgresModelRepo will hold the database connection
type PGModelRepo struct {
	PGConn *sql.DB
}

// ErrNotFound is a public error variable indicating data of given SNo does not exist in database
var ErrNotFound = errors.New("requested item not found")

// NewPGModelRepo will construct the database connection object
func NewPGModelRepo(pmr *sql.DB) *PGModelRepo {
	return &PGModelRepo{PGConn: pmr}
}

// CreateModel creates a model instance in the database
func (pg *PGModelRepo) CreateModel(m *model.ModelInstance) error {
	// query is the sql query to execute to insert the new entry
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

	err := pg.PGConn.QueryRow(query,
		m.SNo,
		m.FirmwareVersion,
		m.CurrentFirmwareVersion,
		m.MeshConfig,
		m.AppConfig,
		m.KCConfig,
	).Scan(&m.SNo)

	return err
}

// UpdateModelbySNO will flip the meshconfig bool and return the updated model.
func (pg *PGModelRepo) UpdateModelbySNO(sno int) (*model.ModelInstance, error) {
	query := `
		UPDATE deviceDBTwo
		SET mesh_configuration = NOT mesh_configuration
		WHERE sno = $1
		RETURNING sno, firmware_version, current_firmware_version, mesh_configuration, app_configuration, kc_configuration`

	var m model.ModelInstance

	err := pg.PGConn.QueryRow(query, sno).Scan(
		&m.SNo,
		&m.FirmwareVersion,
		&m.CurrentFirmwareVersion,
		&m.MeshConfig,
		&m.AppConfig,
		&m.KCConfig,
	)

	return &m, err
}

// GetModelbySNO will retrieve the data associated with the given serial number, else it will return a data does not exist error
func (pg *PGModelRepo) GetModelbySNO(sno int) (*model.ModelInstance, error) {
	query := `
		SELECT
			sno, firmware_version, current_firmware_version, mesh_configuration, app_configuration, kc_configuration
		FROM deviceDBTwo
		WHERE sno = $1`

	var m model.ModelInstance

	err := pg.PGConn.QueryRow(query, sno).Scan(
		&m.SNo,
		&m.FirmwareVersion,
		&m.CurrentFirmwareVersion,
		&m.MeshConfig,
		&m.AppConfig,
		&m.KCConfig,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &m, err
}
