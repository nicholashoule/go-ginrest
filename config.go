package main

import (
	"fmt"
	"net"
	"os"
	"runtime"

	"github.com/thinkerou/favicon"

	"github.com/nicholashoule/go-ginrest/build"
	"github.com/nicholashoule/go-ginrest/crypto"
	"github.com/nicholashoule/go-ginrest/routes"
)

// Exported config variables
var (
	Pwd  string
	Host string
	Ipv4 net.IP
	Url  string
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
	Host = host
	// IPv4 address
	var ipv4 net.IP
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 = addr.To4(); ipv4 != nil {
			ipv4 = addr.To4()
		}
	}
	Ipv4 = ipv4

	// Golang working directory
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Pwd = pwd

	// Display information
	fmt.Println("[Environment Information]")

	// Development
	if build.Environment == "Development" {
		Url = "http://" + ipv4.String() + ":5000/"
	}

	fmt.Printf("Hostname: %s\n", host)
	fmt.Printf("Path: %s\n", pwd)
	fmt.Printf("IPv4: %s\n", ipv4)
	fmt.Printf("Running with %d CPUs\n", numCPU)
	fmt.Printf("URL: %s\n\n", Url)

	// SetTrustedProxies set a list of network origins (IPv4 addresses, IPv4 CIDRs, IPv6 addresses or IPv6 CIDRs)
	routes.R.SetTrustedProxies([]string{"localhost", "127.0.0.1", "::1"})
	routes.R.Use(favicon.New(pwd + "/favicon.ico"))
	routes.R.HandleMethodNotAllowed = true
	routes.R.ForwardedByClientIP = true
}

// startTLS ...
// Listen and serve HTTPS requests
func startTLS() {
	// Configure
	// SSL/TLS
	crypto.Run()
	Url = "https://" + Ipv4.String() + ":5000/"
	// Listen and serve
	fmt.Printf("URL: %s\n", Url)
	routes.R.RunTLS(":5000", "tls/server.crt", "tls/server.key")
}

// Run ...
// Configure, and start the server
func Run() {
	// Configure
	configRuntime()
	routes.GetRoutes()

	// Listen and serve
	routes.R.Run(":5000")
	// startTLS()
}
