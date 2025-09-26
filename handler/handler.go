package handler

// Package handler contains the HTTP handlers for the API.
// It is responsible for parsing requests, calling the service layer, and writing HTTP responses.
import (
	"TAPI/service"
	"errors"
	"net/http"
	"strconv"

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
	// Map the string passed via PATCH verb
	snoStr := c.Param("sno")
	snoInt, err := strconv.Atoi(snoStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid serial number format"})
		return
	}

	// Call the service layer to perform the update logic.
	m, err := s.sci.UpdateMeshStatus(snoInt)
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
	// take in the number:
	snoStr := c.Param("sno")
	snoInt, err := strconv.Atoi(snoStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid serial number format"})
		return
	}

	m, err := s.sci.RetrieveById(snoInt)
	if err != nil {
		if errors.Is(err, service.ErrInvalidSno) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, service.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an internal server error occured"})
		return
	}
	c.JSON(http.StatusOK, m)
	// code for successful retrieval is StatusOK, not StatusFound which is for cases of redirection
}
