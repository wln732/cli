package cli

import "strconv"

type boolVar bool

func (c *command) BoolVar(p *bool, name string, val bool, usage string) {
	c.Var(name, newBoolVar(val, p), usage)
}

func (c *command) Bool(name string, val bool, usage string) *bool {
	p := new(bool)
	c.BoolVar(p, name, val, usage)
	return p
}

func newBoolVar(val bool, p *bool) *boolVar {
	*p = val

	return (*boolVar)(p)
}

func (i *boolVar) Set(x string) error {
	i2, err := strconv.ParseBool(x)
	*i = ((boolVar)(i2))
	return err
}

func (i *boolVar) Get() bool {
	return (bool(*i))
}

func (i *boolVar) String() string {
	return strconv.FormatBool(bool(*i))
}

func (i *boolVar) isBool() bool {
	return true
}
