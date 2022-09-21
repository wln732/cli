package cli

import (
	"os"
)

var (
	commandLine = NewCommand(os.Args[0])
)

func StringVar(p *string, name string, val string, usage string) {
	commandLine.Var(name, newStringVar(val, p), usage)
}

func String(name string, val string, usage string) *string {
	p := new(string)
	commandLine.StringVar(p, name, val, usage)
	return p
}

func Int64Var(p *int64, name string, val int64, usage string) {
	commandLine.Var(name, newInt64Var(val, p), usage)
}

func Int64(name string, val int64, usage string) *int64 {
	p := new(int64)
	commandLine.Int64Var(p, name, val, usage)
	return p
}

func IntVar(p *int, name string, val int, usage string) {
	commandLine.Var(name, newIntVar(val, p), usage)
}

func Int(name string, val int, usage string) *int {
	p := new(int)
	commandLine.IntVar(p, name, val, usage)
	return p
}

func Float64Var(p *float64, name string, val float64, usage string) {
	commandLine.Var(name, newFloat64Var(val, p), usage)
}

func Float64(name string, val float64, usage string) *float64 {
	p := new(float64)
	commandLine.Float64Var(p, name, val, usage)
	return p
}

func BoolVar(p *bool, name string, val bool, usage string) {
	commandLine.Var(name, newBoolVar(val, p), usage)
}

func Bool(name string, val bool, usage string) *bool {
	p := new(bool)
	commandLine.BoolVar(p, name, val, usage)
	return p
}

func Var(name string, val Value, usage string) {
	commandLine.Var(name, val, usage)
}

func Run(args []string) error {
	return commandLine.Run(args)
}
