package routes

import (
	"github.com/gin-gonic/gin"
)

// Exported routes variables
var (
	R = gin.Default()
)

// getRoutes ...
// creates our routes for our entire application. This way every
// group of routes can be defined in their own files.
func GetRoutes() {
	v0 := R.Group("/")
	addDefault(v0) // default

	v1 := R.Group("/v1")
	addGetDate(v1)    // default
	addPingRoutes(v1) // default

	v1_1 := R.Group("/v1.1")
	addGetDate(v1_1)    // default
	addPingRoutes(v1_1) // default

	v2 := R.Group("/v2")
	addGetDate(v2)    // default
	addPingRoutes(v2) // default
	addGetCert(v2)    // crt
	addHash(v2)       // hash
}
