package build

import "fmt"

// Exported build variables
var (
	Environment string
)

// Info ...
func Info(environment string, version string, githash string, datetime string) {
	// Version and build information
	// If environment is not set, assume Development
	if environment == "" {
		Environment = "Development"
	} else {
		Environment = environment
	}

	// build information
	fmt.Printf("\nEnvironment:\t%s", Environment)
	fmt.Printf("\nbuild.Version:\t%s", version)
	fmt.Printf("\nbuild.Time:\t%s", datetime)
	fmt.Printf("\nbuild.GitHash:\t%s\n\n", githash)
}
