package cli

import (
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

	err := cmd.Run(strings.Split("test -i -dir ./19_example -name lxq args2 -n 998 args1  args3", " "))

	if err != nil {
		t.Fatalf("%v\n", err)

	}

	t.Logf("%+v\n", tc)
}
