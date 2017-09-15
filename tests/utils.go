package cwlgotest

import (
	"fmt"
	"os"
)

const version = "1.0"

// Provides file object for testable official .cwl files.
func load(name string) *os.File {
	fpath := fmt.Sprintf("../cwl/v%[1]s/v%[1]s/%s", version, name)
	f, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}
	return f
}
