package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-soap/service"
)

// SOAPController handles SOAP-related requests
type SOAPController struct {
	soapService service.SOAPService
}

// NewSOAPController creates a new instance of SOAPController
func NewSOAPController(soapService service.SOAPService) *SOAPController {
	return &SOAPController{
		soapService: soapService,
	}
}

// HandleSOAPRequest handles the incoming SOAP request
func (ctrl *SOAPController) HandleSOAPRequest(c *gin.Context) {
	var soapBody string
	if err := c.ShouldBind(&soapBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Call service to process SOAP request
	response, err := ctrl.soapService.SendSOAPRequest(soapBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with SOAP response
	c.String(http.StatusOK, response)
}
