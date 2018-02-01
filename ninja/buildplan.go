package ninja

// Compiler provides the support for a given compiler / language.
type Compiler struct {
	Vars  []*Var
	Rules []*Rule
}

// BuildPlans tell ng how to generate a build file.
type BuildPlan struct {
	Compilers []*Compiler
}
