package cli

import "strconv"

type int64Var int64

func (c *command) Int64Var(p *int64, name string, val int64, usage string) {
	c.Var(name, newInt64Var(val, p), usage)
}

func (c *command) Int64(name string, val int64, usage string) *int64 {
	p := new(int64)
	c.Int64Var(p, name, val, usage)
	return p
}

func newInt64Var(val int64, p *int64) *int64Var {
	*p = val

	return (*int64Var)(p)
}

func (i *int64Var) Set(x string) error {
	i2, err := strconv.ParseInt(x, 10, 64)
	*i = ((int64Var)(i2))
	return err
}

func (i *int64Var) String() string {
	return strconv.FormatInt(int64(*i), 10)
}

func (i *int64Var) Get() int64 {
	return (int64(*i))
}
