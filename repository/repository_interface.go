package repository

// Package repository defines the interface for the data access layer.
// This abstraction allows the service layer to be decoupled from the specific database implementation.
import (
	"TAPI/model"
)

// RepoContractDefinition defines the set of operations available for device data.
type RepoContractDefinition interface {
	CreateModel(m *model.ModelInstance) error
	UpdateModelbySNO(sno int) (*model.ModelInstance, error)
	GetModelbySNO(sno int) (*model.ModelInstance, error)
}
