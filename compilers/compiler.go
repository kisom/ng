package compilers

import (
	"fmt"
	"io"

	"github.com/kisom/ng/ninja"
)

// Compiler provides the support for a given compiler / language.
type Compiler struct {
	Vars  []*ninja.Var
	Rules []*ninja.Rule
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
