package compilers

import (
	"os"

	"github.com/kisom/ng/ninja"
)

var CCExts = []string{"c"}

func CC(cc string, debugMode, deps bool) *Compiler {
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

	depFlags := ""
	if deps {
		depFlags = "-MD -MF $out.d "
	}

	CCRule = ninja.NewRule("cc", cc+" $cflags "+depFlags+"-c $in -o $out")
	if deps {
		CCRule.SetVar("depfile", "$out.d")
		CCRule.SetVar("deps", "gcc")
	}

	return &Compiler{
		Vars:  []*ninja.Var{CFlags},
		Rules: []*ninja.Rule{CCRule},
	}
}
