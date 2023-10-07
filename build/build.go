package build

import "fmt"

// Exported build variables
var (
	Environment string
	Version     string
	GitHash     string
	Datetime    string
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
	Version = version
	GitHash = githash
	Datetime = datetime

	// build information
	fmt.Printf("\nEnvironment:\t%s", Environment)
	fmt.Printf("\nbuild.Version:\t%s", Version)
	fmt.Printf("\nbuild.Time:\t%s", Datetime)
	fmt.Printf("\nbuild.GitHash:\t%s\n\n", GitHash)
}
