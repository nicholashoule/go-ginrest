package main

import (
	"github.com/nicholashoule/go-ginrest/build"
	"github.com/nicholashoule/go-ginrest/routes"
)

// main ...
func main() {
	// Build information
	build.Run()

	// Call routes.Run
	routes.Run()
}
