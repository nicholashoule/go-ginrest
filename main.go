package main

import (
	"fmt"

	"github.com/nicholashoule/go-ginrest/build"
	"github.com/nicholashoule/go-ginrest/routes"
)

var version = "Development"

func main() {
	// Version and build information
	fmt.Println("Version:\t", version)
	build.Run()
	// Our server will live in routes
	routes.Run()
}
