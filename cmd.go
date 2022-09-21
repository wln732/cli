package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/tabwriter"
)

type command struct {
	// 命令的名字
	name string
	// 注册的flag
	flags FlagSet

	// 命令的参数
	args []Arg
	// 子命令
	child map[string]*command

	Help string

	Action func(args []Arg) error
}

func NewDefaultCommand() *command {

	return NewCommand(filepath.Base(os.Args[0]))
}

func NewCommand(name string) *command {

	c := &command{
		name:  name,
		flags: make(FlagSet),
		args:  []Arg{},
		child: make(map[string]*command),
	}

	return c
}

// 解析参数
func (c *command) parseArgs(args []string) error {
	// cli -a=1 -b 'a' -c 2

	if args[0][0] != '-' {
		c.args = append(c.args, Arg(args[0]))
		args = args[1:]
	}

	for i := 0; i < len(args); i++ {
		var err error
		if len(args[i]) == 0 || args[i] == "" {
			continue
		}
		// 说明是个flag
		if args[i][0] == '-' {
			var flag *Flag
			index := strings.Index(args[i], "=")
			if index != -1 {
				flag = c.searchFlag(args[i][1:index])
			} else {
				flag = c.searchFlag(args[i][1:])
			}

			if flag == nil {
				c.Usage()
				if args[i][1:] == "h" {
					return E打印帮助信息
				}
				return Eflag未注册
			}

			if i+1 < len(args) {
				if index != -1 {
					err = flag.value.Set(args[i][index+1:])
				} else {

					// 如果下一个不是值，而是 -开头的flag
					if args[i+1][0] == '-' {
						// 如果是bool类型的flag
						b, ok := flag.value.(BoolValue)
						if ok {
							b.Set("true")
							continue
						}

					} else {
						err = flag.value.Set(args[i+1])
					}
					i++
				}

			} else {
				// 最后一个flag
				// 如果是  -a=123这类的
				// 否则是  -a 检查下是不是bool类型额flag
				if index != -1 {
					err = flag.value.Set(args[i][index+1:])
				} else {
					// 如果是bool类型的flag
					b, ok := flag.value.(BoolValue)
					if ok {
						b.Set("true")
						continue
					}

				}
			}

		} else {
			c.args = append(c.args, Arg(args[i]))
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *command) Run(args []string) error {

	if len(args) == 0 {
		c.Usage()
		return nil
	}

	// 如果查找到了子命令
	child := c.searchChildCmd(args[0])
	if child != nil {

		return child.Run(args[1:])

	} else {
		err := c.parseArgs(args[1:])
		if err != nil {
			if err == E打印帮助信息 {
				return nil
			}
			return err
		}
		err = c.Action(c.args)
		return err
	}

}

func (c *command) searchFlag(name string) *Flag {
	return c.flags[name]
}

func (c *command) Var(name string, val Value, usage string) {
	c.flags[name] = &Flag{
		name:  name,
		value: val,
		usage: usage,
	}
}

func (c *command) Usage() {
	maxFlagLength := 0
	maxFlagTypeLength := 0

	w := new(tabwriter.Writer)

	fmt.Println(c.Help)
	for _, flag := range c.flags {
		if len(flag.name) > maxFlagLength {
			maxFlagLength = len(flag.name)
		}
		type_name := reflect.TypeOf(flag.value).Elem().Name()
		if len(type_name) > maxFlagTypeLength {
			maxFlagTypeLength = len(type_name)
		}
	}
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, maxFlagTypeLength+4, 4, 0, '\t', 0)
	fmt.Println("Options:")
	for _, flag := range c.flags {
		fmt.Fprintf(w, "  -%s\t%s\t%s\n", flag.name,

			reflect.TypeOf(flag.value).Elem().Kind().String(),
			flag.usage)

		//fmt.Fprintf(w, "%-5s"+"%-"+strconv.Itoa(maxFlagLength+4)+"s"+
		//	"%-"+strconv.Itoa(maxFlagTypeLength+4)+"s"+"%s\n", "-"+flag.name,
		//
		//	reflect.TypeOf(flag.value).Elem().Name(),
		//	flag.usage)
	}

	w.Flush()
}

func (c *command) AddCommand(name string, child *command) {
	c.child[name] = child
}

func (c *command) searchChildCmd(name string) *command {

	return c.child[name]
}
