package cli

type stringVar string

func (c *command) StringVar(p *string, name string, val string, usage string) {
	c.Var(name, newStringVar(val, p), usage)
}

func (c *command) String(name string, val string, usage string) *string {
	p := new(string)
	c.StringVar(p, name, val, usage)
	return p
}

func newStringVar(val string, p *string) *stringVar {
	*p = val
	return (*stringVar)(p)
}

func (i *stringVar) Set(x string) error {
	*i = ((stringVar)(x))
	return nil
}

func (i *stringVar) String() string {
	return string(*i)
}

func (i *stringVar) Get() string {
	return (string(*i))
}
