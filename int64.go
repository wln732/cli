package cli

import "strconv"

type Int64 int64

func (c *command) Int64Var(p *int64, name string, val int64, usage string) {
	c.Var(name, newInt64Var(val, p), usage)
}

func (c *command) Int64(name string, val int64, usage string) *int64 {
	p := new(int64)
	c.Int64Var(p, name, val, usage)
	return p
}

func newInt64Var(val int64, p *int64) *Int64 {
	*p = val

	return (*Int64)(p)
}

func (i *Int64) Set(x string) error {
	i2, err := strconv.ParseInt(x, 10, 64)
	*i = ((Int64)(i2))
	return err
}

func (i *Int64) String() string {
	return strconv.FormatInt(int64(*i), 10)
}

func (i *Int64) Get() int64 {
	return (int64(*i))
}
