package ninja

// Rules declare a short name for a command line.
type Rule struct {
	name        string
	command     string
	description string
	vars        map[string]string
}

func (r *Rule) String() string {
	s := "rule " + r.name + "\n"

	for k, v := range r.vars {
		s += "  " + k + " = " + v + "\n"
	}

	return s
}

func (r *Rule) Describe(desc string) {
	r.description = desc
}

func (r *Rule) SetVar(k, v string) {
	r.vars[k] = v
}

func NewRule(name, command string) *Rule {
	return &Rule{
		name: name,
		vars: map[string]string{
			"command": command,
		},
	}
}
