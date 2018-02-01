package compilers

import (
	"os"

	"github.com/kisom/ng/ninja"
)

func CXX(cxx string) *ninja.Rule {
	return ninja.NewRule("cxx", cxx+" $cxxflags -c $in -o $out")
}

var CXXFlags *ninja.Var

func EnableCXX(debugMode bool) {
	var debug string
	if os.Getenv("DEBUG") != "" || debugMode {
		debug = "-O0 -g "
	}

	CXXFlags = ninja.NewVar("cxxflags", debug)
	CXXFlags.Append("-Wall -Wextra -pedantic -Wshadow -Wpointer-arith -Wcast-align")
	CXXFlags.Append(" -Wwrite-strings -Wmissing-declarations -Wno-long-long -Werror")
	CXXFlags.Append(" -Wunused-variable -std=c++17 -D_XOPEN_SOURCE-I.")
	CXXFlags.Append(" -fno-elide-constructors -Weffc++ -fPIC")
	envFlags := os.Getenv("CXXFLAGS")
	if len(envFlags) > 0 {
		CXXFlags.Append(" " + envFlags)
	}
}
