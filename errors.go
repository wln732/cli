package cli

import "fmt"

var (
	E参数过少    = newError("参数过少，请重新输入")
	E语法错误    = newError("语法错误，请重新输入")
	E命令未找到   = newError("命令为找到")
	Eflag未注册 = newError("flag未注册")
)

func newError(msg string) error {
	return fmt.Errorf("%s", msg)
}
