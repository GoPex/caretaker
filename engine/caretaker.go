package engine

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// Caretaker struct holding everything needed to serve Unleash application
type Caretaker struct {
	Engine *gin.Engine
}

// Initialize to be executed before the application runs
func (caretaker *Caretaker) Initialize() error {

	// Set the log level to debug
	log.SetLevel(log.DebugLevel)

	return nil
}

// New initialize the Unleash application based on the gin http framework
func New() *Caretaker {

	// Will be used to hold everything needed to serve Unleash
	var caretaker Caretaker

	// Create a default gin stack
	caretaker.Engine = gin.Default()

	// Load html templates
	caretaker.Engine.LoadHTMLGlob("templates/*")

	// Add routes
	caretaker.Engine.GET("/title", func(c *gin.Context) {
		c.HTML(http.StatusOK, "title.tmpl", gin.H{
			"title": "Main website",
		})
	})

	return &caretaker
}
