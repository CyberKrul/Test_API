package repository

// This package will implement the SQL that will fulfill the contract's demands
import (
	"TAPI/model"
	"database/sql"
)

// PostgresModelRepo will hold the database connection
type PGModelRepo struct {
	PGConn *sql.DB
}

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
    kc_configuration,
	) VALUES ($1,$2,$3,$4,$5,$6)
	RETURNING sno`
	return pg.PGConn.Query(query, m.SNo, m.FirmwareVersion, m.CurrentFirmwareVersion, m.MeshConfig, m.AppConfig, m.KCConfig).Scan(&m.sno)
}
