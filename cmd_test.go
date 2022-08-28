package cli

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"testing"
)

func TestGenerateVar(t *testing.T) {

	type T struct {
		Type  string
		_type string
	}

	var data = []T{
		{"Float32", "float32"},
		{"Float64", "float64"},
		{"Bool", "bool"},
		{"String", "string"},
		{"Int64", "int64"},
	}

	temp, err := template.ParseFiles("./int_go.gohtml")
	if err != nil {
		t.Fatalf("ParseFiles err:%b\n", err)
	}

	for i := 0; i < len(data); i++ {
		f2, err3 := os.Create("./var/" + data[i]._type + ".go")
		if err3 != nil {
			t.Fatalf("create err:%v\n", err)
		}
		defer f2.Close()

		err = temp.Execute(f2, data[i])
		if err3 != nil {
			t.Fatalf("Execute err:%v\n", err)
		}
	}

}

func TestNewCommand(t *testing.T) {
	args := []string{
		"echo",
		"-a",
		"1",
		"-d",
		"-c=123",
		"golang",
		"php",
		"-b",
		"'hello world'",
	}

	cmd := NewCommand("echo")

	var a = new(int64)
	var b = new(string)
	var c = new(int)
	var d = new(bool)
	var e = new(float64)
	cmd.Int64Var(a, "a", 100, "打印a")
	cmd.StringVar(b, "b", "789456123", "打印b")
	cmd.IntVar(c, "c", 999, "打印c")
	cmd.BoolVar(d, "d", false, "打印d")
	cmd.Float64Var(e, "e", 123456789.987654321, "打印e")
	err := cmd.parseArgs(args)
	if err != nil {
		t.Fatalf("err=%v\n", err)
	}

	t.Logf("a=%v \n", *a)
	t.Logf("b=%v \n", *b)
	t.Logf("c=%v \n", *c)
	t.Logf("d=%v \n", *d)
	t.Logf("e=%v \n", *e)
	t.Logf("%v\n", cmd.args)
}

func TestNewCommand2(t *testing.T) {

	cmd := NewCommand("cli")
	args := []string{
		"echo",
		"-a",
		"1",
		"-d",
		"-c=123",
		"golang",
		"php",
		"-b",
		"'hello world'",
	}

	var a = new(int64)
	var b = new(string)
	var c = new(int)
	var d = new(bool)
	var e = new(float64)
	cmd.Int64Var(a, "a", 100, "打印a")
	cmd.StringVar(b, "b", "789456123", "打印b")
	cmd.IntVar(c, "c", 999, "打印c")
	cmd.BoolVar(d, "d", false, "打印d")
	cmd.Float64Var(e, "e", 123456789.987654321, "打印e")

	cmd.Action = func(flags FlagSet, args []string) error {
		fmt.Println(args)
		return nil
	}

	err := cmd.Run(args[1:])
	if err != nil {
		t.Errorf("err=%v\n", err)
	}

}

func TestCommand(t *testing.T) {

	cmd := flag.NewFlagSet("echo", flag.PanicOnError)
	var a = new(int64)
	var b = new(string)
	var c = new(int)
	var d = new(bool)
	var e = new(float64)
	cmd.Int64Var(a, "a", 100, "打印a")
	cmd.StringVar(b, "b", "789456123", "打印b")
	cmd.IntVar(c, "c", 999, "打印c")
	cmd.BoolVar(d, "d", false, "打印d")
	cmd.Float64Var(e, "e", 123456789.987654321, "打印e")

	cmd.Parse([]string{"-h"})

}
