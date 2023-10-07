package main

import (
	"github.com/nicholashoule/go-ginrest/build"
)

// main ...
func main() {
	// Build information
	build.Info()

	// Call routes.Run
	Run()
}
