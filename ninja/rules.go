package ninja

// Rules declare a short name for a command line.
type Rule struct {
	name        string
	command     string
	description string
}

func (r *Rule) String() string {
	s := "rule: " + r.name + "\n"
	s += "  command = " + r.command + "\n"
	if len(r.description) > 0 {
		s += "  description = " + r.description + "\n"
	}
	return s
}

func (r *Rule) Describe(desc string) {
	r.description = desc
}

func NewRule(name, command string) *Rule {
	return &Rule{
		name:    name,
		command: command,
	}
}
