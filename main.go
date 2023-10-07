package main

import (
	"github.com/nicholashoule/go-ginrest/build"
)

// variables
var (
	environment string
	version     string
	githash     string
	datetime    string
)

// main ...
func main() {
	// Build information
	build.Info(environment, version, githash, datetime)

	// Call routes.Run
	Run()
}
