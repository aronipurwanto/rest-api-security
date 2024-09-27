package config

import (
	"log"

	"github.com/spf13/viper"
)

// AppConfig holds the application's configuration
type AppConfig struct {
	SOAP SOAPConfig
}

// SOAPConfig holds the SOAP service configuration
type SOAPConfig struct {
	Endpoint string
	Username string
	Password string
	Timeout  int
}

// LoadConfig loads configuration from the config.yaml file
func LoadConfig() (*AppConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return nil, err
	}

	config := &AppConfig{
		SOAP: SOAPConfig{
			Endpoint: viper.GetString("soap.endpoint"),
			Username: viper.GetString("soap.username"),
			Password: viper.GetString("soap.password"),
			Timeout:  viper.GetInt("soap.timeout"),
		},
	}
	return config, nil
}
