package cli

import "strconv"

type Int int

func (c *command) IntVar(p *int, name string, val int, usage string) {
	c.Var(name, newIntVar(val, p), usage)
}

func (c *command) Int(name string, val int, usage string) *int {
	p := new(int)
	c.IntVar(p, name, val, usage)
	return p
}

func newIntVar(val int, p *int) *Int {
	*p = val

	return (*Int)(p)
}

func (i *Int) Set(x string) error {
	i2, err := strconv.Atoi(x)
	*i = ((Int)(i2))
	return err
}

func (i *Int) String() string {
	return strconv.Itoa(int(*i))
}

func (i *Int) Get() int {
	return (int(*i))
}
