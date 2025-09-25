package repository

// This package will define the contract encapsulated in the interface

import (
	"TAPI/model"
)

// The model has 4 defined actions:
// create a new model
// update the mesh config of an existing model(identified by SNO)
// get the model data by SNO
type RepoContractDefinition interface {
	createModel(*model.ModelInstance) error
	updateModelbySNO(sno int) (*model.ModelInstance, error)
	getModelbySNO(sno int) (*model.ModelInstance, error)
}
