package config

import (
	"testing"

	"github.com/kisom/goutils/assert"
)

var testConfig = `compilers:
  cxx: clang++
  cc: clang

targets:
  cc:
    ch01ex01: []
    ch01ex02:
      - ch01ex04.cc
  c:
    vm:
      - vm.c
      - stack.c
      - isa.c
`

func TestParse(t *testing.T) {
	cfg, err := Parse([]byte(testConfig))
	assert.NoErrorT(t, err)

	assert.BoolT(t, cfg.Compilers.LD == "clang++")
}
