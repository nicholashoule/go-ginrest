package build

import "fmt"

var (
	Time    string
	GitHash string
)

func Run() {
	// build information
	fmt.Printf("build.Time:\t%s", Time)
	fmt.Printf("\nbuild.GitHash:\t%s\n\n", GitHash)
}
