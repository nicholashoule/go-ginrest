package main

import (
	"github.com/nicholashoule/go-ginrest/build"
)

// Exported build variables
var (
	Environment string
	Version     string
	GitHash     string
	DateTime    string
)

// main ...
func main() {
	// Build information
	build.Info(Environment, Version, GitHash, DateTime)

	// Call routes.Run
	Run()
}
