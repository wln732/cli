package cli

import (
	"fmt"
	"strings"
	"testing"
)

type TestCommond struct {
	I    bool   `flag:"i-treytrytrytr"  usage:"i是个flag"`
	Dir  string `flag:"dir-dirdirdirdirdirdir" usage:"dir是个flag"`
	Name string `flag:"name-namenamenamename" usage:"name是个flag"`
	N    int    `flag:"n-nnnnxxxxxxxx" usage:"n是个flag"`
}

func (t *TestCommond) Run(args []Arg) error {
	fmt.Println("test commond 运行了", args)

	return nil
}

func Test_parseStruct(t *testing.T) {

	tc := &TestCommond{
		I:    false,
		Dir:  "./18_regie",
		Name: "wln",
		N:    100,
	}

	cmd := parseStruct("test", tc)

	cmd.Action = func(args []Arg) error {
		fmt.Println("我是新添加的函数")
		return nil
	}

	err := cmd.Run(strings.Split("test -h", " "))

	if err != nil {
		t.Fatalf("%v\n", err)

	}

	t.Logf("%+v\n", tc)
}
