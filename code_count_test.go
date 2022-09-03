package cli

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test代码统计行数(t *testing.T) {

	type CodeCount struct {
		name string
		line int
	}

	cmd := NewCommand("cc")
	cmd.Help = `
	cc [options] dir1,dir2 | file1,file2, ...
统计go语言代码
	`

	var dir string
	var exclude string
	var exclude_dir string

	cmd.StringVar(&dir, "dir", "./", "输入被统计go语言代码的目录")
	cmd.StringVar(&exclude, "exclude", "./", "排除一些文件")
	cmd.StringVar(&exclude_dir, "exclude-dir", "./", "排除一些目录")
	var tt string

	cmd.StringVar(&tt, "exclude-dir111111111111111", "./", "tttttttttttttt")

	cmd.Action = func(flags FlagSet, args []string) error {
		t.Logf("dir=%s\n", dir)
		fi, err := ioutil.ReadDir(dir)
		if err != nil {
			t.Fatalf("readdir err=%v\n", err)
		}

		filenames := make([]string, 0, 16)

		for i := 0; i < len(fi); i++ {
			// 如果是目录的话
			if fi[i].IsDir() {

				filenames = append(filenames,
					获取目录下所有文件名(t, filepath.Join(dir, fi[i].Name()))...)

			} else if fi[i].IsDir() == false {

				if strings.HasSuffix(fi[i].Name(), ".go") {
					filenames = append(filenames, filepath.Join(dir, fi[i].Name()))
				}

			}
		}

		cc := []CodeCount{}

		for i := 0; i < len(filenames); i++ {
			cc = append(cc, CodeCount{
				name: filenames[i],
				line: 统计代码行数(t, filenames[i]),
			})
		}

		var total = 0
		for i := 0; i < len(cc); i++ {
			t.Logf("name=%s, code=%d\n", cc[i].name, cc[i].line)
			total += cc[i].line
		}

		t.Logf("total=%d\n", total)

		return nil

	}

	args := []string{
		"-h",
	}

	err := cmd.Run(args)
	if err != nil {
		t.Fatalf("err=%v\n", err)
	}
}

func 获取目录下所有文件名(t *testing.T, dir string) []string {
	filenames := make([]string, 0, 16)
	f, err := os.Open(dir)
	if err != nil {
		t.Logf("打开目录:%s 错误:%v\n", dir, err)
		return nil
	}
	defer f.Close()

	finfo, err := f.ReadDir(-1)
	if err != nil {
		t.Logf("获取目录:%s 错误:%v\n", dir, err)
		return nil
	}

	for i := 0; i < len(finfo); i++ {
		if finfo[i].IsDir() == false {
			if strings.HasSuffix(finfo[i].Name(), ".go") {
				filenames = append(filenames, filepath.Join(dir, finfo[i].Name()))
			}
		} else {

			filenames = append(
				filenames,
				获取目录下所有文件名(t, filepath.Join(dir, finfo[i].Name()))...)

		}
	}

	return filenames
}

func 统计代码行数(t *testing.T, name string) int {

	defer func() {
		err := recover()
		if err != nil {
			panic(fmt.Errorf("file=%s, 统计代码行数err:%v\n", name, err))
		}
	}()

	if strings.HasSuffix(name, ".go") == false {
		return 0
	}
	f, err := os.Open(name)
	if err != nil {
		t.Logf("打开文件:%s 错误:%v\n", name, err)
	}

	defer f.Close()
	const KB = 1024
	reader := bufio.NewReaderSize(f, 16*KB)

	code := 0

	line_coment := 0

	for {
		buf, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Logf("打开文件:%s 错误:%v\n", name, err)
		}

		buf = bytes.TrimLeft(buf, "\t")
		buf = bytes.TrimLeft(buf, " ")

		if len(buf) > 1 {
			if string(buf[:2]) == "//" {

				continue
			}
			/*
				1
				2
				3
				4
			*/

			/*
				1
				2
				3
				4*/

			if string(buf[:2]) == "/*" {
				line_coment = 1
			}

		}
		// 如果多行注释已经有了开头
		// 判断结尾，
		// 有两种情况
		// 一种是 单独的行以 */ 结束
		// 一种是 跟在注释的最后面两个字符结束
		if line_coment > 0 {
			if len(buf) > 1 {
				if string(buf[:2]) == "*/" {
					line_coment = 0
					continue
				}

				if string(buf[len(buf)-2:]) == "*/" {
					line_coment = 0
					continue
				}
			}

		}
		if line_coment == 0 && len(buf) != 0 {
			code++
		}

	}

	return code
}
