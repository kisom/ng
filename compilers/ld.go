package compilers

import (
	"os"

	"github.com/kisom/ng/ninja"
)

func LD(ld string) *Compiler {
	var (
		LDFlags *ninja.Var
		LDRule  *ninja.Rule
	)

	LDFlags = ninja.NewVar("ldflags", "")
	envFlags := os.Getenv("LDFLAGS")
	if len(envFlags) > 0 {
		LDFlags.Append(" " + envFlags)
	}

	LDRule = ninja.NewRule("ld", ld+" $ldflags -o $out $in")
	return &Compiler{
		Vars:  []*ninja.Var{LDFlags},
		Rules: []*ninja.Rule{LDRule},
	}
}
