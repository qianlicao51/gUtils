package cmd

import (
	"bytes"
	"os/exec"
)

// ExecuteCmd 执行cmd命令
//命令的所有参数都需要分开填写才行
func ExecuteCmd(cmdShell string, arg ...string) (err error) {
	var outInfo bytes.Buffer
	cmd := exec.Command(cmdShell, arg...)
	// 设置接收
	cmd.Stdout = &outInfo
	// 执行|Start 不等待结果，Run 等待结果
	err = cmd.Run()
	if err != nil {
		//log.Println("失败 ", err, arg)
		return err
	} else {
		//log.Println("执行结束", outInfo.String(), arg)
	}
	return err
}
