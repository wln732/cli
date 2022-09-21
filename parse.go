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
func parseStruct(name string,i any) *command {
	if i == nil {
		return nil
	}

	val := reflect.ValueOf(i)
	if val.Kind() != reflect.Ptr {
		fmt.Println("必须传入指针类型")
		return nil
	}

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

			usage := field.Tag.Get("usage")

			switch value:=val.Field(i).Interface().(type) {
			case string:


				cmd.Var(tag,
					newStringVar(
						value, UintptrTo[string](val.Field(i).Addr().Pointer())),
					usage)
			case int:

				cmd.Var(tag,
					newIntVar(value, UintptrTo[int](val.Field(i).Addr().Pointer())), usage)

			case int64:
	

				cmd.Var(tag,
					newInt64Var(value, UintptrTo[int64](val.Field(i).Addr().Pointer())), usage)
			case bool:
				
				cmd.Var(tag,
					newBoolVar(value, UintptrTo[bool](val.Field(i).Addr().Pointer())), usage)

			case float64:
				

				cmd.Var(tag,
					newFloat64Var(value, UintptrTo[float64](val.Field(i).Addr().Pointer())), usage)

			}
		}

	}

	if iface,ok:=i.(Interface);ok {
		cmd.Action = func(flags FlagSet, args []string) error {
			return iface.Run(args)
		}
	}
	
	return cmd
}



func UintptrTo[T any](p uintptr) *T {
	return (*T)(unsafe.Pointer(p))
}