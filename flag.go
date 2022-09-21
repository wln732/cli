package cli

type Value interface {
	Set(string) error
	String() string
}

type BoolValue interface {
	Value
	isBool() bool
}

type FlagSet map[string]*Flag

type Flag struct {
	name  string
	value Value
	usage string
}

func NewFlag(name string, val Value, usage string) *Flag {
	return &Flag{
		name:  name,
		value: val,
		usage: usage,
	}
}

func (f FlagSet) Lookup(name string) *Flag {
	return f[name]
}
