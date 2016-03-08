package main

import (
	"github.com/spf13/viper"

	"github.com/GoPex/caretaker/engine"
)

func main() {
	viper.SetEnvPrefix("caretaker")

	viper.SetDefault("PORT", 3000)

	viper.AutomaticEnv()

	// Create a new Unleash application
	application := engine.New()

	// Initialize the application
	if err := application.Initialize(); err != nil {
		panic("Not able to initialize the application ! Cause: " + err.Error())
	}

	// Listen and serve on port defined by environment variable UNLEASH_PORT
	if err := application.Engine.Run(":" + viper.GetString("port")); err != nil {
		panic("Error while starting unleash ! Cause: " + err.Error())
	}
}
