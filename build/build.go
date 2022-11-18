package build

import "fmt"

// Exported
var (
	Time    string
	GitHash string
)

// Run ...
func Run() {
	// build information
	fmt.Printf("build.Time:\t%s", Time)
	fmt.Printf("\nbuild.GitHash:\t%s\n\n", GitHash)
}
