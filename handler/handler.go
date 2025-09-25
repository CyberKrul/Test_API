package handler

// this package will handle all things HTTP, and acts as the highest, front facing interface

import (
	"TAPI/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ==================================================
// SERVICE STRUCT TYPES
// ==================================================
// RegisterDeviceRequest defines the shape of the JSON body for a device registration request.
type RegisterDeviceRequest struct {
	SNo             int `json:"sno"`
	FirmwareVersion int `json:"firmwareVersion"`
}

// UpdateMeshRequest defines the shape of the JSON body for a mesh update request
type UpdateMeshRequest struct {
	SNo int `json:"sno"`
}

// RetrieveDeviceRequest  defines the shape of the JSON body for a device retrieval request
type RetrieveDeviceRequest struct {
	SNo int `json:"sno"`
}

// ==================================================

// ServiceContractInstance consumes the service contract between the handler and service layer
type ServiceContractInstance struct {
	sci service.ServiceContractDefinition
}

// NewServiceContractInstance creates a new instance of ServiceContractInstance
func NewServiceContractInstance(sci service.ServiceContractDefinition) *ServiceContractInstance {
	return &ServiceContractInstance{
		sci: sci,
	}
}

// HandleCreateDeviceRequest will handle the HTTP request to add a new device to the DB
func (s *ServiceContractInstance) HandleCreateDeviceRequest(c *gin.Context) {
	// instantiate a new object to bind the HTTP JSON to
	var newDevice RegisterDeviceRequest

	// Gin will handle decoding and error response
	if err := c.ShouldBindJSON(&newDevice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body"})
		return
	}

	// call the service layer here
	m, err := s.sci.RegisterDevice(newDevice.SNo, newDevice.FirmwareVersion)
	if err != nil {
		if errors.Is(err, service.ErrInvalidSno) || errors.Is(err, service.ErrInvalidFV) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// For any other unexpected error, return a 500.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal server error occurred"})
		return
	}

	c.JSON(http.StatusCreated, m)
}

// HandleUpdateMeshRequest will handle the HTTP request to flip the mesh status of device with given serial number
func (s *ServiceContractInstance) HandleUpdateMeshRequest(c *gin.Context) {
	// instantiate a new object to bind the HTTP JSON to
	var device UpdateMeshRequest

	// Gin will handle decoding and error response
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body"})
		return
	}

	// call the service layer here
	m, err := s.sci.UpdateMeshStatus(device.SNo)
	if err != nil {
		if errors.Is(err, service.ErrInvalidSno) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, service.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal server error occurred"})
		return
	}

	c.JSON(http.StatusOK, m)
}

// HandleDeviceRetrieval will handle the HTTP request to retrieve information on a particular device
func (s *ServiceContractInstance) HandleDeviceRetrieval(c *gin.Context) {
	// instantiate a new object to bind the HTTP JSON to
	var device RetrieveDeviceRequest

	// Gin will handle decoding and error response
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body"})
		return
	}

	// call the service layer here
	m, err := s.sci.RetrieveById(device.SNo)
	if err != nil {
		// Check for a validation error first.
		if errors.Is(err, service.ErrInvalidSno) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Check if the device was not found.
		if errors.Is(err, service.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		// For any other unexpected error, return a 500.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal server error occurred"})
		return
	}

	c.JSON(http.StatusOK, m)
}
