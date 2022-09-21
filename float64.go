package cli

import "strconv"

type float64Var float64

func (c *command) Float64Var(p *float64, name string, val float64, usage string) {
	c.Var(name, newFloat64Var(val, p), usage)
}

func (c *command) Float64(name string, val float64, usage string) *float64 {
	p := new(float64)
	c.Float64Var(p, name, val, usage)
	return p
}

func newFloat64Var(val float64, p *float64) *float64Var {
	*p = val

	return (*float64Var)(p)
}

func (i *float64Var) Set(x string) error {
	i2, err := strconv.Atoi(x)
	*i = ((float64Var)(i2))
	return err
}

func (i *float64Var) String() string {
	return strconv.FormatFloat(float64(*i), 'f', -1, 64)
}

func (i *float64Var) Get() float64 {
	return (float64(*i))
}
