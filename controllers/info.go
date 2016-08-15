package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/GoPex/caretaker/bindings"
	"github.com/GoPex/caretaker/helpers"
)

// GetPing is the handler for the GET /info/ping route.
// This will respond by a pong JSON message if the server is alive
func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, bindings.PingResponse{Pong: "OK"})
}

// GetStatus is an handler for the GET /info/status route.
// This will respond  by the status of the server and of the docker host in a
// JSON message.
func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK,
		bindings.StatusResponse{Status: "OK"},
	)
}

// GetVersion is an handler for the GET /info/version route. This will respond a
// JSON message with application version and the version of Docker running in the Docker host.
func GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK,
		bindings.VersionResponse{Version: helpers.AppVersion},
	)
}
