package routes

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"

	"github.com/nicholashoule/go-ginrest/crypto"
)

// Variables
var (
	r = gin.Default()
)

// ConfigRuntime ...
func configRuntime() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	//gin.DisableConsoleColor()
	// Sets the number of operating system threads.
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)

	// SetTrustedProxies set a list of network origins (IPv4 addresses, IPv4 CIDRs, IPv6 addresses or IPv6 CIDRs)
	r.SetTrustedProxies([]string{"localhost", "127.0.0.1", "::1"})
	r.Use(favicon.New("./favicon.ico"))
	r.HandleMethodNotAllowed = true
	r.ForwardedByClientIP = true
}

// getRoutes ...
// creates our routes for our entire application. This way every
// group of routes can be defined in their own files.
func getRoutes() {
	v0 := r.Group("/")
	addDefault(v0) // default

	v1 := r.Group("/v1")
	addGetDate(v1)    // default
	addPingRoutes(v1) // default

	v1_1 := r.Group("/v1.1")
	addGetDate(v1_1)    // default
	addPingRoutes(v1_1) // default

	v2 := r.Group("/v2")
	addGetDate(v2)    // default
	addPingRoutes(v2) // default
	addGetCert(v2)    // crt
	addHash(v2)       // hash
}

// Run ...
// Configure, and start the server
func Run() {
	// Configure
	crypto.Run()
	configRuntime()
	getRoutes()

	// Listen and serve
	//r.RunTLS(":5000", "./tls/server.crt", "./tls/server.key")
	r.Run(":5000")
}
