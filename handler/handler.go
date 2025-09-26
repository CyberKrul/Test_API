package handler

// Package handler contains the HTTP handlers for the API.
// It is responsible for parsing requests, calling the service layer, and writing HTTP responses.
import (
	"TAPI/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ==================================================
// REQUEST/RESPONSE MODELS
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

// ServiceContractInstance is a concrete implementation of the HTTP handlers.
// It holds a reference to the service layer interface to perform business logic.
type ServiceContractInstance struct {
	sci service.ServiceContractDefinition
}

// NewServiceContractInstance creates a new handler instance with the given service.
func NewServiceContractInstance(sci service.ServiceContractDefinition) *ServiceContractInstance {
	return &ServiceContractInstance{
		sci: sci,
	}
}

// HandleCreateDeviceRequest handles the HTTP request to register a new device.
func (s *ServiceContractInstance) HandleCreateDeviceRequest(c *gin.Context) {
	// Bind the incoming JSON request to the RegisterDeviceRequest struct.
	var newDevice RegisterDeviceRequest

	// Gin will handle decoding and error response
	if err := c.ShouldBindJSON(&newDevice); err != nil {
		// Provide a more detailed error message for debugging.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	// Call the service layer to perform the registration logic.
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

// HandleUpdateMeshRequest handles the HTTP request to toggle the mesh status of a device.
func (s *ServiceContractInstance) HandleUpdateMeshRequest(c *gin.Context) {
	// Bind the incoming JSON request to the UpdateMeshRequest struct.
	var device UpdateMeshRequest

	// Gin will handle decoding and error response
	if err := c.ShouldBindJSON(&device); err != nil {
		// Provide a more detailed error message for debugging.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	// Call the service layer to perform the update logic.
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

// HandleDeviceRetrieval handles the HTTP request to retrieve a device's information by its serial number.
func (s *ServiceContractInstance) HandleDeviceRetrieval(c *gin.Context) {
	// Bind the incoming JSON request to the RetrieveDeviceRequest struct.
	var device RetrieveDeviceRequest

	// Gin will handle decoding and error response
	if err := c.ShouldBindJSON(&device); err != nil {
		// Provide a more detailed error message for debugging.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	// Call the service layer to perform the retrieval logic.
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
