package ninja

import (
	"testing"

	"github.com/kisom/goutils/assert"
)

func TestNewRule(t *testing.T) {
	expected := `rule: test
  command = echo "hello, world"
  description = Print a test message.
`
	r := NewRule("test", `echo "hello, world"`)
	r.Describe("Print a test message.")
	assert.BoolT(t, r.String() == expected)
}
