package cli

import (
	"strconv"
)

type Arg string

func (a Arg) Int() int {
	n, _ := a.IntE()

	return n
}

func (a Arg) IntE() (int, error) {
	n, err := strconv.Atoi(string(a))
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (a Arg) MustInt() int {
	n, err := a.IntE()
	if err != nil {
		panic(err)
	}
	return n
}

func (a Arg) Int64() int64 {
	n, _ := a.Int64E()

	return n
}

func (a Arg) Int64E() (int64, error) {
	n, err := strconv.ParseInt(string(a), 10, 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (a Arg) MustInt64() int64 {
	n, err := a.Int64E()
	if err != nil {
		panic(err)
	}
	return n
}

func (a Arg) Float64() float64 {
	n, _ := a.Float64E()

	return n
}

func (a Arg) Float64E() (float64, error) {
	n, err := strconv.ParseFloat(string(a), 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (a Arg) MustFloat64() float64 {
	n, err := a.Float64E()
	if err != nil {
		panic(err)
	}
	return n
}

func (a Arg) Bool() bool {
	n, _ := a.BoolE()

	return n
}

func (a Arg) BoolE() (bool, error) {
	n, err := strconv.ParseBool(string(a))
	if err != nil {
		return false, err
	}
	return n, nil
}

func (a Arg) MustBool() bool {
	n, err := a.BoolE()
	if err != nil {
		panic(err)
	}
	return n
}

func (a Arg) String() string {
	return string(a)
}
