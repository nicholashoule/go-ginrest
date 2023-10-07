package build

import "fmt"

// Exported build variables
var (
	Time        string
	GitHash     string
	Version     string
	Environment string
)

// Info ...
func Info() {
	// Version and build information
	Version = "1.0.0"
	Environment = "Development"

	// build information
	fmt.Printf("\nEnvironment:\t%s", Environment)
	fmt.Printf("\nVersion:\t%s", Version)
	fmt.Printf("\nbuild.Time:\t%s", Time)
	fmt.Printf("\nbuild.GitHash:\t%s\n\n", GitHash)
}
