package helpers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/kelseyhightower/envconfig"
)

var (
	// Config make things easier but is bad
	Config *Specification

	// AppName describes the name of the application
	AppName = "Caretaker"

	// AppVersion describes the version of the application
	AppVersion = "0.1.0"
)

// Describe will log all variables parsed with envconfig
func (specification *Specification) Describe() {
	log.WithFields(log.Fields{
		"Port":     specification.Port,
		"LogLevel": specification.LogLevel,
	}).Info("caretaker initialized !")
}

// Specification to hold the configuration of the application
type Specification struct {
	Port     string `default:"3000"`
	LogLevel string `envconfig:"log_level" default:"debug"`
}

// ParseConfiguration will parse the configuration of Unleash based on environment variables
func ParseConfiguration() (Specification, error) {
	// Gather the configuration
	var config Specification
	if err := envconfig.Process("caretaker", &config); err != nil {
		return config, err
	}
	return config, nil
}
