# cwl.go

[![Build Status](https://travis-ci.org/otiai10/cwl.go.svg?branch=master)](https://travis-ci.org/otiai10/cwl.go) [![GoDoc](https://godoc.org/github.com/otiai10/cwl.go?status.svg)](https://godoc.org/github.com/otiai10/cwl.go)

`cwl.go` is just a parser of CWL file and input files based on [CWL](https://github.com/common-workflow-language/common-workflow-language), for example [1st-tool.yaml](https://github.com/common-workflow-language/common-workflow-language/blob/master/v1.0/examples/1st-tool.cwl) and [echo-job.yml](https://github.com/common-workflow-language/common-workflow-language/blob/master/v1.0/examples/echo-job.yml).

Fully documented [here!](https://godoc.org/github.com/otiai10/cwl.go)

# Example

```go
package main

import (
	"fmt"
	"os"

	cwl "github.com/otiai10/cwl.go"
)

func main() {
	file, _ := os.Open("hello.cwl")
	doc := cwl.NewCWL()
	doc.Decode(file)
	fmt.Printf("%+v\n", doc)
}
```
