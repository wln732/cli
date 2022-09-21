package cli

import "strconv"

type intVar int

func (c *command) IntVar(p *int, name string, val int, usage string) {
	c.Var(name, newIntVar(val, p), usage)
}

func (c *command) Int(name string, val int, usage string) *int {
	p := new(int)
	c.IntVar(p, name, val, usage)
	return p
}

func newIntVar(val int, p *int) *intVar {
	*p = val

	return (*intVar)(p)
}

func (i *intVar) Set(x string) error {
	i2, err := strconv.Atoi(x)
	*i = ((intVar)(i2))
	return err
}

func (i *intVar) String() string {
	return strconv.Itoa(int(*i))
}

func (i *intVar) Get() int {
	return (int(*i))
}
