package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest-api-soap/config"
	"time"
)

// SOAPService interface defines methods for interacting with SOAP
type SOAPService interface {
	SendSOAPRequest(soapBody string) (string, error)
}

// soapServiceImpl implements the SOAPService interface
type soapServiceImpl struct {
	config config.SOAPConfig
}

// NewSOAPService creates a new instance of SOAPService
func NewSOAPService(cfg config.SOAPConfig) SOAPService {
	return &soapServiceImpl{
		config: cfg,
	}
}

// SendSOAPRequest sends a SOAP request to the configured SOAP service
func (s *soapServiceImpl) SendSOAPRequest(soapBody string) (string, error) {
	client := &http.Client{
		Timeout: time.Duration(s.config.Timeout) * time.Second,
	}

	req, err := http.NewRequest("POST", s.config.Endpoint, bytes.NewBuffer([]byte(soapBody)))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Add necessary headers and set Basic Auth
	req.Header.Add("Content-Type", "text/xml")
	req.SetBasicAuth(s.config.Username, s.config.Password)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending SOAP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	return string(body), nil
}
