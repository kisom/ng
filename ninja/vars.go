package ninja

// Vars provide support for declaring shorter, reusable names for
// strings.
type Var struct {
	name  string
	value string
}

func (v *Var) String() string {
	return v.name + " = " + v.value
}

func NewVar(name, value string) *Var {
	return &Var{
		name:  name,
		value: value,
	}
}

func (v *Var) Append(s string) {
	v.value += s
}
