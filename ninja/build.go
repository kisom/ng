package ninja

import "strings"

type Build struct {
	output string
	inputs []string
	rule   string
}

func (b *Build) String() string {
	return "build " + b.output + ": " + b.rule + " " + strings.Join(b.inputs, " ")
}

func (b *Build) Output() string {
	return b.output
}

func NewBuild(output, rule string, inputs ...string) *Build {
	return &Build{
		output: output,
		rule:   rule,
		inputs: inputs,
	}
}
