package util

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

// Command 执行命令并将执行结果 写入writer
func Command(CommandString string, writer io.Writer) error {
	name, arg := StringToCommand(CommandString)
	cmd := exec.Command(name, arg...)
	cmd.Stdout = writer
	cmd.Stderr = writer
	err := cmd.Run()
	if err != nil {
		return errors.New(fmt.Sprintf("%s: exit error:%s", CommandString, err.Error()))
	}
	return nil
}

// StringToCommand 将命令转换为exec.Command需要的参数
func StringToCommand(Command string) (string, []string) {
	arr := strings.Split(Command, " ")
	var name string
	var arg = make([]string, 0)
	for k, v := range arr {
		if k == 0 {
			name = v
		} else {
			arg = append(arg, v)
		}
	}
	return name, arg
}
