package compilers

import (
	"os"

	"github.com/kisom/ng/ninja"
)

func CC(cc string) *ninja.Rule {
	return ninja.NewRule("cc", cc+" $cflags -c $in -o $out")
}

var CCExts = []string{"c"}

func EnableCC(bp *ninja.BuildPlan, cc string, debugMode bool) {
	var (
		CFlags *ninja.Var
		CCRule *ninja.Rule
		debug  string
	)

	if os.Getenv("DEBUG") != "" || debugMode {
		debug = "-O0 -g "
	}

	CFlags = ninja.NewVar("cflags", debug)
	CFlags.Append("-Wall -Wextra -pedantic -Wshadow -Wpointer-arith -Wcast-align")
	CFlags.Append(" -Wwrite-strings -Wmissing-prototypes -Wmissing-declarations")
	CFlags.Append(" -Wnested-externs -Winline -Wno-long-long -Wunused-variable")
	CFlags.Append(" -Wstrict-prototypes -Werror")
	envFlags := os.Getenv("CFLAGS")
	if len(envFlags) > 0 {
		CFlags.Append(" " + envFlags)
	}

	CCRule = ninja.NewRule("cc", cc+" -$cflags -c $in -o $out")
	compiler := &ninja.Compiler{
		Vars:  []*ninja.Var{CFlags},
		Rules: []*ninja.Rule{CCRule},
	}

	if bp.Compilers == nil {
		bp.Compilers = map[string]*ninja.Compiler{}
	}
	bp.Compilers[cc] = compiler
}
