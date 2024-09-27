package service

import (
	"rest-api-soap/config"
	"testing"
)

func TestSendSOAPRequest(t *testing.T) {
	cfg := config.SOAPConfig{
		Endpoint: "http://example.com/soap",
		Username: "testuser",
		Password: "testpassword",
		Timeout:  30,
	}

	soapService := NewSOAPService(cfg)

	soapBody := `
		<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://www.example.org/webservice/">
			<soapenv:Header/>
			<soapenv:Body>
				<web:MyRequest>
					<!-- SOAP request body here -->
				</web:MyRequest>
			</soapenv:Body>
		</soapenv:Envelope>
	`

	response, err := soapService.SendSOAPRequest(soapBody)
	if err != nil {
		t.Fatalf("Failed to send SOAP request: %v", err)
	}

	if len(response) == 0 {
		t.Errorf("Expected non-empty response")
	}
}
