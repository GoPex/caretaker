package engine

import (
	"html/template"
	"path/filepath"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"

	"github.com/GoPex/caretaker/controllers"
	"github.com/GoPex/caretaker/helpers"
)

// Application struct holding everything needed to serve Caretaker application
type Application struct {
	Engine *gin.Engine
	Config *helpers.Specification
}

// loadTemplates loads every templates applying includes while loading them
func loadTemplates(templatesDir string) multitemplate.Render {
	r := multitemplate.New()

	layouts, err := filepath.Glob(filepath.Join(templatesDir, "layouts/*.tmpl"))
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(filepath.Join(templatesDir, "includes/*.tmpl"))
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		layoutName := strings.TrimSuffix(filepath.Base(layout), filepath.Ext(layout))
		r.Add(layoutName, template.Must(template.ParseFiles(files...)))
	}

	return r
}

// Initialize to be executed before the application runs
func (caretaker *Application) Initialize(config *helpers.Specification) error {

	// Set the log level to debug
	logLevel, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}
	log.SetLevel(logLevel)

	// Print all configuration variables
	config.Describe()

	// Assign the incoming configuration
	caretaker.Config = config

	// FIXME: Attribute the configuration globally for ease of use
	helpers.Config = config

	return nil
}

// New initialize the Application application based on the gin http framework
func New() *Application {

	// Will be used to hold everything needed to serve Caretaker
	var application Application

	// Create an empty configuration to avoid panic
	application.Config = &helpers.Specification{}

	// Create a default gin stack
	application.Engine = gin.Default()

	// Load templates
	application.Engine.HTMLRender = loadTemplates("templates")

	// Routes
	// Ping route
	application.Engine.GET("/ping", controllers.GetPing)

	// Info routes
	info := application.Engine.Group("/info")
	info.GET("/status", controllers.GetStatus)
	info.GET("/version", controllers.GetVersion)

	// Root
	application.Engine.GET("/", controllers.GetHome)

	// Load all static assets
	application.Engine.Static("/static", "./static")

	return &application
}
