package ninja

import (
	"fmt"
	"io"

	"github.com/kisom/ng/config"
)

// Compiler provides the support for a given compiler / language.
type Compiler struct {
	Vars  []*Var
	Rules []*Rule
}

func (c *Compiler) Emit(w io.Writer, name string) error {
	_, err := fmt.Fprintln(w, "# ng compiler section: "+name)
	if err != nil {
		return err
	}

	for i := range c.Vars {
		_, err = fmt.Fprintln(w, c.Vars[i])
		if err != nil {
			return err
		}
	}

	for i := range c.Rules {
		_, err = fmt.Fprintln(w, c.Rules[i])
		if err != nil {
			return err
		}
	}

	return err
}

// BuildPlans tell ng how to generate a build file.
type BuildPlan struct {
	Compilers map[string]*Compiler
	Targets   []*Build

	// Targets is used for ordering, sources is used to prevent
	// duplicates.
	sources map[string]*Build
}

func (bp *BuildPlan) getObjectFor(source, rule string) (string, *Build) {
	objFile := source + ".o"
	return objFile, NewBuild(objFile, rule, source)
}

func (bp *BuildPlan) generateBuildsFor(target, rule string, sources ...string) []*Build {
	objs := make([]string, 0, len(sources))
	builds := make([]*Build, 0, len(sources)+1)
	for _, source := range sources {
		o, b := bp.getObjectFor(source, rule)
		objs = append(objs, o)
		builds = append(builds, b)
	}

	builds = append(builds, NewBuild(target, rule, objs...))
	return builds
}

func (bp *BuildPlan) addTargetsFor(confTargets map[string][]string, rule string) {
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

func (bp *BuildPlan) Plan(cfg *config.Config) error {
	if bp.Compilers == nil {
		bp.Compilers = map[string]*Compiler{}
	}

	if bp.sources == nil {
		bp.sources = map[string]*Build{}
	}

	bp.addTargetsFor(cfg.Targets.CC, "cc")
	bp.addTargetsFor(cfg.Targets.CXX, "cxx")
	return nil
}

func (bp *BuildPlan) Emit(w io.Writer) error {
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
