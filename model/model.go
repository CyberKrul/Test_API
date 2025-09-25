package model

// pkg model will define the internal representation of the data

//modelInstance is the internal representation of a unit
type ModelInstance struct {
	SNo                    int
	FirmwareVersion        int
	CurrentFirmwareVersion bool
	MeshConfig             bool
	AppConfig              bool
	KCConfig               bool
}

func NewModelInstance() ModelInstance {
	return ModelInstance{
		SNo:                    0,
		FirmwareVersion:        0,
		CurrentFirmwareVersion: false,
		MeshConfig:             false,
		AppConfig:              false,
		KCConfig:               false,
	}
}
