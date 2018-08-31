package jsonindent

import (
	"os"
	"testing"

	. "github.com/otiai10/mint"
)

func TestNewEncoder(t *testing.T) {
	encoder := NewEncoder(os.Stdout)
	Expect(t, encoder).TypeOf("*json.Encoder")
}
