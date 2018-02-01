package compilers

import (
	"os"

	"github.com/kisom/ng/ninja"
)

func CC(cc string) *ninja.Rule {
	return ninja.NewRule("cc", cc+" $cflags -c $in -o $out")
}

var CFlags *ninja.Var

func EnableCC(debugMode bool) {
	var debug string
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
}
