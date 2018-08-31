package jsonindent

import "os"

func ExampleNewEncoder() {
	v := map[string]interface{}{"name": "otiai10", "age": 30, "foo": map[string]int{"bar": 100}}
	NewEncoder(os.Stdout).Encode(v)
	// Output:
	// {
	// 	"age": 30,
	// 	"foo": {
	// 		"bar": 100
	// 	},
	// 	"name": "otiai10"
	// }
}
