package cli

type String string

func StringVar(p *int64, name string, value int64, usage string) *Flag {
	return NewFlag(name, newInt64Var(value, p), usage)
}

func (c *command) StringVar(p *string, name string, val string, usage string) {
	c.Var(name, newStringVar(val, p), usage)
}

func (c *command) String(name string, val string, usage string) *string {
	p := new(string)
	c.StringVar(p, name, val, usage)
	return p
}

func newStringVar(val string, p *string) *String {
	*p = val
	return (*String)(p)
}

func (i *String) Set(x string) error {
	*i = ((String)(x))
	return nil
}

func (i *String) String() string {
	return string(*i)
}

func (i *String) Get() string {
	return (string(*i))
}
