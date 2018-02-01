package main

import (
	"flag"
	"io"
	"os"

	"github.com/kisom/goutils/die"
)

const (
	defaultSpec = "build.yaml"
	defaultOut  = "build.ninja"
)

func main() {
	var outf *io.WriteCloser = os.Stdout
	var spec string
	var err error

	flag.StringVar(&spec, "f", defaultSpec, "specify the build file")
	flag.Parse()
	if flag.NArg() == 1 {
		outf, err = os.Create(flag.Arg(0))
		die.If(err)
	} else {
		outf, err = os.Create(defaultOut)
		die.If(err)
	}
}
