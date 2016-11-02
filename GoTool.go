package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	proc, errMsg := call("go.exe")
	if errMsg != nil {
		fmt.Println(errMsg.Error())
	}

	_, procMsg := proc.Wait()
	if procMsg != nil {
		fmt.Println(procMsg.Error())
	}
}

// 创建新进程
// fileName:新进程的文件名
// 返回值：
// error:异常数据
func call(fileName string) (*os.Process, error) {
	flName, erMsg := exec.LookPath(fileName)
	if erMsg != nil {
		return nil, erMsg
	}

	callArgs := os.Args
	callArgs = callArgs[1:]

	env := getEnv()
	wd, _ := os.Getwd()
	procAtt := os.ProcAttr{
		Dir:   wd,
		Env:   env,
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Sys:   nil,
	}

	proc, procErr := os.StartProcess(flName, callArgs, &procAtt)
	if procErr != nil {
		return nil, procErr
	}

	return proc, nil
}

// 获取传递给子线程的环境变量值
// 返回获取到的环境变量值
func getEnv() []string {
	nowEnv := os.Environ()
	nowEnv = nowEnv[:]

	// 填充配置值
	nowEnv = fill(nowEnv)

	return nowEnv
}
