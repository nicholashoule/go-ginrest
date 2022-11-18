package main

import (
	"fmt"

	"github.com/nicholashoule/go-ginrest/build"
	"github.com/nicholashoule/go-ginrest/routes"
)

// Variables
var version = "Development"

// main ...
func main() {
	// Version and build information
	fmt.Println("Version:\t", version)
	build.Run()

	// Call routes.Run
	routes.Run()
}
