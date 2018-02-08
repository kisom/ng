package planner

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/kisom/goutils/die"
	"github.com/kisom/ng/compilers"
	"github.com/kisom/ng/config"
	"github.com/kisom/ng/ninja"
)

// Plan tell ng how to generate a build file.
type Plan struct {
	Compilers map[string]*compilers.Compiler
	Targets   []*ninja.Build

	// Targets is used for ordering, sources is used to prevent
	// duplicates.
	sources map[string]*ninja.Build
}

func (bp *Plan) addCompiler(name string, comp *compilers.Compiler) {
	if bp.Compilers[name] == nil {
		bp.Compilers[name] = comp
	}
}

func (bp *Plan) getObjectFor(source, rule string) (string, *ninja.Build) {
	objFile := source[0:len(source)-len(filepath.Ext(source))] + ".o"
	return objFile, ninja.NewBuild(objFile, rule, source)
}

func (bp *Plan) generateBuildsFor(target, rule string, sources ...string) []*ninja.Build {
	objs := make([]string, 0, len(sources))
	builds := make([]*ninja.Build, 0, len(sources)+1)
	for _, source := range sources {
		o, b := bp.getObjectFor(source, rule)
		objs = append(objs, o)
		builds = append(builds, b)
	}

	builds = append(builds, ninja.NewBuild(target, "ld", objs...))
	return builds
}

func (bp *Plan) addTargetsFor(confTargets map[string][]string, rule string) {
	for target, sources := range confTargets {
		if bp.sources[target] != nil {
			continue
		}

		builds := bp.generateBuildsFor(target, rule, sources...)
		for _, build := range builds {
			bp.Targets = append(bp.Targets, build)
			bp.sources[build.Output()] = build
		}
	}
}

func (bp *Plan) Execute(cfg *config.Config) error {
	if bp.Compilers == nil {
		bp.Compilers = map[string]*compilers.Compiler{}
	}

	if bp.sources == nil {
		bp.sources = map[string]*ninja.Build{}
	}

	if cfg.RequiresCC() && bp.Compilers["cc"] == nil {
		bp.addCompiler("cc", compilers.CC(cfg.Compilers.CC, cfg.Debug, !cfg.NoDeps))
		bp.addCompiler("ld", compilers.LD(cfg.Compilers.LD))
	}

	if cfg.RequiresCXX() && bp.Compilers["cxx"] == nil {
		bp.addCompiler("cxx", compilers.CXX(cfg.Compilers.CXX, cfg.Debug, !cfg.NoDeps))
		bp.addCompiler("ld", compilers.LD(cfg.Compilers.LD))

		err := compilers.CXXScanTargets(cfg)
		die.If(err)
	}

	bp.addTargetsFor(cfg.Targets.CC, "cc")
	bp.addTargetsFor(cfg.Targets.CXX, "cxx")
	return nil
}

func (bp *Plan) Emit(w io.Writer) error {
	for name, compiler := range bp.Compilers {
		err := compiler.Emit(w, name)
		if err != nil {
			return err
		}
	}

	_, err := fmt.Fprintln(w, "\n# ng build rules")
	if err != nil {
		return err
	}

	for _, target := range bp.Targets {
		_, err = fmt.Fprintln(w, target)
		if err != nil {
			return err
		}
	}

	return err
}
