package cli

import "strconv"

type Bool bool

func (c *command) BoolVar(p *bool, name string, val bool, usage string) {
	c.Var(name, newBoolVar(val, p), usage)
}

func (c *command) Bool(name string, val bool, usage string) *bool {
	p := new(bool)
	c.BoolVar(p, name, val, usage)
	return p
}

func newBoolVar(val bool, p *bool) *Bool {
	*p = val

	return (*Bool)(p)
}

func (i *Bool) Set(x string) error {
	i2, err := strconv.ParseBool(x)
	*i = ((Bool)(i2))
	return err
}

func (i *Bool) String() string {
	return strconv.FormatBool(bool(*i))
}

func (i *Bool) isBool() bool {
	return true
}
