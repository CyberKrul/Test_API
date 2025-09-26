package repository

// Package repository defines the interface for the data access layer.
// This abstraction allows the service layer to be decoupled from the specific database implementation.
import (
	"TAPI/model"
	"context"
)

// RepoContractDefinition defines the set of operations available for device data.
type RepoContractDefinition interface {
	// Add context.Context as the first argument to all methods.
	CreateModel(ctx context.Context, m *model.ModelInstance) error
	UpdateModelbySNO(ctx context.Context, sno int) (*model.ModelInstance, error)
	GetModelbySNO(ctx context.Context, sno int) (*model.ModelInstance, error)
}
