package service

// pkg service operates the business logic of the interface, including validation, and sanity checks

import (
	"TAPI/model"
	"TAPI/repository"
	"database/sql"
	"errors"
	"fmt"
)

// ==================================================
// ERROR TYPES FOR SERVICE LAYER
// ==================================================
var ErrInvalidSno = fmt.Errorf("validation failed: SNo must be 8 digits")
var ErrInvalidFV = fmt.Errorf("firmware version is incorrect")
var ErrNoRows = fmt.Errorf("cannot update MeshConfig, no device found with this serial number")

// ServiceContractDefinition is the contract between the handler and the service
type ServiceContractDefinition interface {
	RegisterDevice(sno, firmwareVersion int) (*model.ModelInstance, error)
	UpdateMeshStatus(sno int) (*model.ModelInstance, error)
	RetrieveById(sno int) (*model.ModelInstance, error)
}

// RepoContractInstance consumes the contract between the repository and the service
type RepoContractInstance struct {
	RCinst repository.RepoContractDefinition
}

// NewRepoContractInstance creates a new instance of the Repo Contract
func NewRepoContractInstance(r repository.RepoContractDefinition) *RepoContractInstance {
	return &RepoContractInstance{
		RCinst: r,
	}
}

// RegisterDevice registers the received serial number and firmware version
func (rci *RepoContractInstance) RegisterDevice(sno, firmwareVersion int) (*model.ModelInstance, error) {
	if err := validateSno(sno); err != nil {
		return nil, err
	}
	if err := validateFirmwareVersion(firmwareVersion); err != nil {
		return nil, err
	}

	// Create a new model instance with the provided data and default values.
	m := model.ModelInstance{
		SNo:                    sno,
		FirmwareVersion:        firmwareVersion,
		CurrentFirmwareVersion: true, // A new device has the current version
	}

	// saving the values to the database
	err := rci.RCinst.CreateModel(&m)

	return &m, err
}

// UpdateMeshStatus validates the SNo, calls the repository to update the mesh status, and handles not-found errors.
func (rci *RepoContractInstance) UpdateMeshStatus(sno int) (*model.ModelInstance, error) {
	if err := validateSno(sno); err != nil {
		return nil, err
	}
	model, err := rci.RCinst.UpdateModelbySNO(sno)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRows
		}
		return nil, err
	}
	return model, err
}

// RetrieveByID retrieves the device if it exists
func (rci *RepoContractInstance) RetrieveById(sno int) (*model.ModelInstance, error) {
	if err := validateSno(sno); err != nil {
		return nil, err
	}

	// Call the repository to get the model by its SNo
	m, err := rci.RCinst.GetModelbySNO(sno)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Translate the database error to a service-level error
			return nil, ErrNoRows
		}
		return nil, err
	}
	return m, nil
}
