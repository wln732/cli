package cli

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
)

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

	cmd.Action = func(args []Arg) error {
		fmt.Println(args)
		return nil
	}

	err := cmd.Run(args[1:])
	if err != nil {
		t.Errorf("err=%v\n", err)
	}

}

func TestStdCommand(t *testing.T) {

	cmd := flag.NewFlagSet("echo", flag.ContinueOnError)
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

func TestChildCmd(t *testing.T) {

	c1 := NewCommand("c1")
	c2 := NewCommand("c2")
	c3 := NewCommand("c3")

	c1.AddCommand(c2.name, c2)
	c2.AddCommand(c3.name, c3)

	c1.Help = `
	我是c11111111子命令
	`

	c2.Help = `
	我是c2子命令
	`

	c3.Help = `
	我是c3子命令
	`

	var a = new(int64)

	c3.Int64Var(a, "c3a", 0xc3a, "我是c3——a")
	var c1s string
	c1.StringVar(&c1s, "c1s", "我是cl命令", "cl flag")

	c1.Action = func(args []Arg) error {
		fmt.Println("我是cmd1", c1s)
		return nil
	}

	c2.Action = func(args []Arg) error {
		fmt.Println("我是cmd2")
		return nil
	}

	c3.Action = func(args []Arg) error {
		fmt.Println("我是cmd3")
		return nil
	}

	err := c1.Run([]string{
		"c1", "-c1s", "wln",
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("c3a=%d\n", *a)
}

func Test_command_Usage(t *testing.T) {

	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 5, 4, 0, '-', 0)
	fmt.Fprintln(w, "adfd\t你a\t刘锣哥二人\td\t.")

	w.Flush()

}
