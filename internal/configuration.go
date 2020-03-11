package internal

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

// Config : Configuration extracted from environment variables
type Config struct {
	Host string `default:"0.0.0.0"`
	Port int    `default:"80"`

	WorkDir         string `default:"/data"`
	WKHTMLTOPDFPath string `default:"wkhtmltopdf"`
}

// GetConfig read the configuration from the environment
func GetConfig() *Config {
	var config = &Config{}

	if err := envconfig.Process("", config); err != nil {
		log.Fatalf("failed to parse configuration from environment: %v", err)
	}

	return config
}
