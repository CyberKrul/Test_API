package service

// Package service implements the business logic for the application.
// It orchestrates data flow between the handlers and the repository and enforces business rules.
import (
	"TAPI/model"
	"TAPI/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// ==================================================
// ERROR TYPES FOR SERVICE LAYER
// ==================================================
var ErrInvalidSno = fmt.Errorf("validation failed: SNo must be 8 digits")
var ErrInvalidFV = fmt.Errorf("firmware version is incorrect")
var ErrNoRows = fmt.Errorf("service: device not found")

// ServiceContractDefinition defines the interface for the service layer,
// outlining the business operations that can be performed.
type ServiceContractDefinition interface {
	RegisterDevice(ctx context.Context, sno, firmwareVersion int) (*model.ModelInstance, error)
	UpdateMeshStatus(ctx context.Context, sno int) (*model.ModelInstance, error)
	RetrieveById(ctx context.Context, sno int) (*model.ModelInstance, error)
}

// RepoContractInstance is a concrete implementation of the ServiceContractDefinition.
// It holds a reference to a repository to interact with the data layer.
type RepoContractInstance struct {
	RCinst repository.RepoContractDefinition
}

// NewRepoContractInstance creates a new service instance with the given repository.
func NewRepoContractInstance(r repository.RepoContractDefinition) *RepoContractInstance {
	return &RepoContractInstance{
		RCinst: r,
	}
}

// RegisterDevice validates input data, creates a new device model, and persists it via the repository.
func (rci *RepoContractInstance) RegisterDevice(ctx context.Context, sno, firmwareVersion int) (*model.ModelInstance, error) {
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

	// Persist the new model to the database.
	err := rci.RCinst.CreateModel(ctx, &m)

	return &m, err
}

// UpdateMeshStatus validates the SNo and calls the repository to toggle the device's mesh status.
// It translates repository-level errors into service-level errors.
func (rci *RepoContractInstance) UpdateMeshStatus(ctx context.Context, sno int) (*model.ModelInstance, error) {
	if err := validateSno(sno); err != nil {
		return nil, err
	}
	model, err := rci.RCinst.UpdateModelbySNO(ctx, sno)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRows
		}
		return nil, err
	}
	return model, err
}

// RetrieveById validates the SNo and retrieves the corresponding device from the repository.
// It translates repository-level errors into service-level errors.
func (rci *RepoContractInstance) RetrieveById(ctx context.Context, sno int) (*model.ModelInstance, error) {
	if err := validateSno(sno); err != nil {
		return nil, err
	}

	// Call the repository to get the model by its SNo
	m, err := rci.RCinst.GetModelbySNO(ctx, sno)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Translate the database error to a service-level error
			return nil, ErrNoRows
		}
		return nil, err
	}
	return m, nil
}
