package shell

import (
	"fmt"
	"os/exec"
)

/*
执行sheel 脚本文件
*/
func execShellFile(filPath string) {

	cmd := exec.Command("touch", filPath)

	err := cmd.Run()

	if err != nil {

		fmt.Errorf("exec sheel file faile ", err)

	}

	fmt.Print("Sheel exec Success")

}
