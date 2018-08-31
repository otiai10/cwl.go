package yaml2json

import (
	"fmt"
	"os"
)

func ExampleY2J() {
	src, _ := os.Open("./testdata/001.yaml")
	b, err := Y2J(src)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	// Output:
	// {"baz":[100,200],"foo":"bar","spam":{"ham":[2,{"age":17,"name":"otiai10"},"5"]}}
}
