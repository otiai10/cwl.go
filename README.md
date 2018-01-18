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

# Tests

## Prerequisite

`xtest.sh` requires Go package `github.com/otiai10/mint` 

To install it.

```
go get -u github.com/otiai10/mint
```

## Why xtest.sh and How to do test with it.

Because there are both array and dictionary in CWL specification, and as you know Golang can't keep order of map keys, the test fails sometimes by order problem. Therefore, [`./xtest.sh`](https://github.com/otiai10/cwl.go/blob/master/xtest.sh) tries testing each case several times eagerly unless it passes.

For all cases,

```sh
./xtest.sh
```

For only 1 case which matches `_wf3`,

```sh
./xtest.sh _wf3
```

Or if you want to execute single test for just 1 time (NOT eagerly),

```sh
go test ./tests -run _wf3
```
