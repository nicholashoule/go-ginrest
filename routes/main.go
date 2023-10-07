package routes

import (
	"fmt"
	"net"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"

	"github.com/nicholashoule/go-ginrest/build"
	"github.com/nicholashoule/go-ginrest/crypto"
)

// Variables
var (
	r    = gin.Default()
	Pwd  string
	Host string
	Ipv4 net.IP
)

// ConfigRuntime ...
func configRuntime() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	//gin.DisableConsoleColor()
	// Sets the number of operating system threads.
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	// Get some OS information
	// Hostname
	host, _ := os.Hostname()
	// IPv4 address
	var ipv4 net.IP
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 = addr.To4(); ipv4 != nil {
			ipv4 = addr.To4()
		}
	}

	// Golang working directory
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Display information
	Pwd = pwd
	Host = host
	Ipv4 = ipv4
	fmt.Println("[Environment Information]")
	fmt.Printf("Hostname: %s\n", host)
	fmt.Printf("Path: %s\n", pwd)
	fmt.Printf("IPv4: %s\n", ipv4)
	fmt.Printf("Running with %d CPUs\n\n", numCPU)

	// Development
	if build.Environment == "Development" {
		// SSL/TLS
		crypto.Run()
		fmt.Printf("Try loading: https://%s:5000/v2/date/\n", ipv4)
	}

	// SetTrustedProxies set a list of network origins (IPv4 addresses, IPv4 CIDRs, IPv6 addresses or IPv6 CIDRs)
	r.SetTrustedProxies([]string{"localhost", "127.0.0.1", "::1"})
	r.Use(favicon.New(pwd + "/favicon.ico"))
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

// startTLS ...
// Listen and serve HTTPS requests
func startTLS() {
	// Configure
	// Listen and serve
	r.RunTLS(":5000", "tls/server.crt", "tls/server.key")
}

// Run ...
// Configure, and start the server
func Run() {
	// Configure
	configRuntime()
	getRoutes()

	// Listen and serve
	// r.Run(":5000")
	startTLS()
}
