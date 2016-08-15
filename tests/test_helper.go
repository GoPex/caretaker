package tests

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	"github.com/GoPex/caretaker/helpers"
)

// Global variables to be used by all tests of this package
var (
	ContextLogger = log.WithFields(log.Fields{
		"environment": "test",
	})

	AppConfigTest helpers.Specification
)

// Initialze test for the whole package
func init() {
	// Force gin in test mode
	gin.SetMode(gin.TestMode)

	// Force logrus to log only from warnings
	log.SetLevel(log.WarnLevel)

	// Mock App configuration
	AppConfigTest = helpers.Specification{LogLevel: "warning"}
	helpers.Config = &AppConfigTest
}
