package main

import (
	"github.com/GoPex/caretaker/engine"
	"github.com/GoPex/caretaker/helpers"
)

func main() {
	// Create a new Unleash application
	application := engine.New()

	// Parse the configuration
	config, err := helpers.ParseConfiguration()
	if err != nil {
		panic("Not able to parse the configuration ! Cause: " + err.Error())
	}

	// Initialize the application
	if err = application.Initialize(&config); err != nil {
		panic("Not able to initialize the application ! Cause: " + err.Error())
	}

	// Check the configuration
	if application.Config.Port == "" {
		panic("No port was given !")
	}

	// Listen and serve on port defined by environment variable UNLEASH_PORT
	if err := application.Engine.Run(":" + application.Config.Port); err != nil {
		panic("Error while starting unleash ! Cause: " + err.Error())
	}
}
