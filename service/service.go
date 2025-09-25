package service

// pkg service operates the business logic of the interface, including validation, and sanity checks

import (
	"TAPI/model"
	"TAPI/repository"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

// ==================================================
// ERROR TYPES FOR SERVICE LAYER
// ==================================================
var ErrInvalidSno = fmt.Errorf("Length of SNo must be 8 digits or more")
var ErrInvalidFV = fmt.Errorf("Firmware version is incorrect")
var ErrNoRows = fmt.Errorf("Cannot update MeshConfig, no device found with this Serial number")

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
	// validating that the sno is 8 digits long:
	if len(strconv.Itoa(sno)) != 8 {
		return nil, ErrInvalidSno
	}
	// validating that the firmware version is valid
	if firmwareVersion > 8 {
		return nil, ErrInvalidFV
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

// RegisterDevice registers the received serial number and firmware version
func (rci *RepoContractInstance) UpdateMeshStatus(sno int) (*model.ModelInstance, error) {
	// validating that the sno is 8 digits long:
	if len(strconv.Itoa(sno)) != 8 {
		return nil, ErrInvalidSno
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
func (rci *RepoContractInstance) RetrieveByID(sno int) (*model.ModelInstance, error) {
	// validating that the sno is 8 digits long:
	if len(strconv.Itoa(sno)) != 8 {
		return nil, ErrInvalidSno
	}

	// Create a new model instance with given data and default values.
	m := model.ModelInstance{
		SNo:                    sno,
		FirmwareVersion:        0,
		CurrentFirmwareVersion: true, // A new device has the current version
	}

	// saving the values to the database
	err := rci.RCinst.CreateModel(&m)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRows
		}
		return nil, err
	}
	return &m, err
}
