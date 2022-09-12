package cli

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

type Interface interface {
	Run(args []string) error
}

type TestCommond struct {
	I    bool   `flag:"i"  usage:"i是个flag"`
	Dir  string `flag:"dir" usage:"dir是个flag"`
	Name string `flag:"name" usage:"name是个flag"`
	N    int    `flag:"n" usage:"n是个flag"`
}

func (t *TestCommond) Run(args []string) error {
	fmt.Println("test commond 运行了", args)

	return nil
}

// 解析结构体
func parseStruct(i Interface) *command {
	if i == nil {
		return nil
	}

	val := reflect.ValueOf(i)
	if val.Kind() != reflect.Ptr {
		fmt.Println("必须传入指针类型")
		return nil
	}
	name := val.Elem().Type().Name()
	name = strings.ToLower(name)
	cmd := NewCommand(name)

	fmt.Println(val.Elem().NumField())

	for val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)

		if field.IsExported() {
			tag := field.Tag.Get("flag")
			if tag == "" {
				tag = strings.ToLower(field.Name)
			}
			flag_name := strings.Split(tag, ",")
			var defaultVal string
			if len(flag_name) == 2 {
				defaultVal = flag_name[1]
			}

			usage := field.Tag.Get("usage")

			switch val.Field(i).Interface().(type) {
			case string:
				var sVar String
				sVar.Set(defaultVal)

				cmd.Var(flag_name[0],
					newStringVar(
						sVar.String(), (*string)((unsafe.Pointer)(val.Field(i).Addr().Pointer()))),
					usage)
			case int:
				var intVar Int
				intVar.Set(defaultVal)

				cmd.Var(flag_name[0],
					newIntVar(intVar.Get(), (*int)((unsafe.Pointer)(val.Field(i).Addr().Pointer()))), usage)

			case int64:
				var intVar Int64
				intVar.Set(defaultVal)

				cmd.Var(flag_name[0],
					newInt64Var(intVar.Get(), (*int64)((unsafe.Pointer)(val.Field(i).Addr().Pointer()))), usage)
			case bool:
				var boolVar Bool
				boolVar.Set(defaultVal)

				cmd.Var(flag_name[0],
					newBoolVar(boolVar.Get(), (*bool)((unsafe.Pointer)(val.Field(i).Addr().Pointer()))), usage)

			case float64:
				var f64Var Float64
				f64Var.Set(defaultVal)

				cmd.Var(flag_name[0],
					newFloat64Var(f64Var.Get(), UintptrTo[float64](val.Field(i).Addr().Pointer())), usage)

			}
		}

	}

	cmd.Action = func(flags FlagSet, args []string) error {
		return i.Run(args)
	}
	return cmd
}



func UintptrTo[T any](p uintptr) *T {
	return (*T)(unsafe.Pointer(p))
}