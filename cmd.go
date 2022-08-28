package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type command struct {
	// 命令的名字
	name string
	// 注册的flag
	flags FlagSet

	// 命令的参数
	args []string
	// 子命令
	child map[string]*command

	Action func(flags FlagSet, args []string) error
}

func NewDefaultCommand(args []string) *command {

	return NewCommand(filepath.Base(os.Args[0]))
}

func NewCommand(name string) *command {

	c := &command{
		name:  name,
		flags: make(map[string]*Flag),
		args:  []string{},
		child: make(map[string]*command),
	}

	return c
}

// 解析参数
func (c *command) parseArgs(args []string) error {
	// cli -a=1 -b 'a' -c 2

	if len(args) < 2 {
		c.Usage()
		return nil
	}

	if args[0][0] != '-' {
		c.args = append(c.args, args[0])
		args = args[1:]
	}

	for i := 0; i < len(args); i++ {
		var err error
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
				return Eflag未注册
			}

			if args[i][1:] == "h" {
				c.Usage()
				return nil
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
			c.args = append(c.args, args[i])
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *command) Run(args []string) error {

	err := c.parseArgs(args)
	if err != nil {
		return err
	}
	err = c.Action(c.flags, c.args)
	return err
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
	for _, flag := range c.flags {
		fmt.Printf("%-s %-7s %s\n", "-"+flag.name,
			reflect.TypeOf(flag.value).Elem().Name(),
			flag.usage)
	}
}
