package cli

import (
	"fmt"
	"strings"
	"testing"
)

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
