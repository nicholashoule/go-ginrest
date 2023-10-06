package build

import "fmt"

// Exported
var (
	Time    string
	GitHash string
	Version string
)

// Run ...
func Run() {
	// Version and build information
	Version = "Development"

	// build information
	fmt.Printf("\nVersion:\t%s", Version)
	fmt.Printf("\nbuild.Time:\t%s", Time)
	fmt.Printf("\nbuild.GitHash:\t%s\n", GitHash)
}
