package main

import (
	"flag"
	"io"
	"io/ioutil"
	"os"

	"github.com/kisom/goutils/die"
	"github.com/kisom/ng/config"
	"github.com/kisom/ng/planner"
)

const (
	defaultSpec = "build.yaml"
	defaultOut  = "build.ninja"
)

func main() {
	var (
		outf    io.WriteCloser = os.Stdout
		nw, dbg bool
		spec    string
		err     error
	)

	flag.BoolVar(&dbg, "d", false, "perform a debug build")
	flag.StringVar(&spec, "f", defaultSpec, "specify the build file")
	flag.BoolVar(&nw, "n", false, "don't write output file, just print to standard output")
	flag.Parse()

	in, err := ioutil.ReadFile(spec)
	die.If(err)

	cfg, err := config.Parse(in)
	die.If(err)

	if dbg {
		cfg.Debug = dbg
	}

	bp := &planner.Plan{}
	err = bp.Execute(cfg)
	die.If(err)

	if flag.NArg() == 1 {
		outf, err = os.Create(flag.Arg(0))
		die.If(err)
	} else if !nw {
		outf, err = os.Create(defaultOut)
		die.If(err)
	}
	defer outf.Close()

	err = bp.Emit(outf)
	die.If(err)
}
