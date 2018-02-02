package compilers

import (
	"errors"
	"os"

	"github.com/kisom/goutils/fileutil"
	"github.com/kisom/ng/config"
	"github.com/kisom/ng/ninja"
)

var CXXExts = []string{"cc", "cxx", "c++"}

func CXX(cxx string, debugMode bool) *Compiler {
	var (
		CXXFlags *ninja.Var
		CXXRule  *ninja.Rule
		debug    string
	)

	if os.Getenv("DEBUG") != "" || debugMode {
		debug = "-O0 -g "
	}

	CXXFlags = ninja.NewVar("cxxflags", debug)
	CXXFlags.Append("-Wall -Wextra -pedantic -Wshadow -Wpointer-arith -Wcast-align")
	CXXFlags.Append(" -Wwrite-strings -Wmissing-declarations -Wno-long-long -Werror")
	CXXFlags.Append(" -Wunused-variable -std=c++14 -D_XOPEN_SOURCE -I.")
	CXXFlags.Append(" -fno-elide-constructors -Weffc++ -fPIC")
	envFlags := os.Getenv("CXXFLAGS")
	if len(envFlags) > 0 {
		CXXFlags.Append(" " + envFlags)
	}

	CXXRule = ninja.NewRule("cxx", cxx+" $cxxflags -c $in -o $out")
	return &Compiler{
		Vars:  []*ninja.Var{CXXFlags},
		Rules: []*ninja.Rule{CXXRule},
	}
}

func CXXScanTargets(cfg *config.Config) error {
	for target, sources := range cfg.Targets.CXX {
		if len(sources) != 0 {
			continue
		}

		newSources := []string{}

		for _, ext := range CXXExts {
			source := target + "." + ext
			if fileutil.FileDoesExist(source) {
				newSources = append(newSources, source)
				break
			}
		}

		if len(newSources) > 0 {
			cfg.Targets.CXX[target] = newSources
		} else {
			return errors.New("couldn't find any valid sources for " + target)
		}
	}

	return nil
}
