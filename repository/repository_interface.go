package repository

// This package will define the contract encapsulated in the interface

import (
	"TAPI/model"
)

// The model has 3 defined actions:
// create a new model
// update the mesh config of an existing model(identified by SNO)
// get the model data by SNO
type RepoContractDefinition interface {
	CreateModel(m *model.ModelInstance) error
	UpdateModelbySNO(sno int) (*model.ModelInstance, error)
	GetModelbySNO(sno int) (*model.ModelInstance, error)
}
