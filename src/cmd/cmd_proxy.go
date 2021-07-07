package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
)

//go build -ldflags -H=windowsgui  隐藏 go 自身的 Dos窗口
//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} 隐藏 windwos 的 Dos 窗口
func SetProxy() {
	sysTyoe := runtime.GOOS

	if sysTyoe == "linux" {

		fmt.Print("operating system is linux")
	}

	if sysTyoe == "windows" {
		fmt.Print("operating system is windows")

	}
}

var (
	proxy_url string = "http://127.0.0.1:7777"
)

func SetWinCmdProxy() {
	// set http_proxy_user=user
	// set http_proxy_pass=pass
	// set https_proxy_user=user
	// set https_proxy_pass=pass
	c := exec.Command("cmd", "set", "http_proxy="+proxy_url)
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
	c = exec.Command("cmd", "set", "https_proxy="+proxy_url)
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}

func ClearWinCmdProxy() {
	c := exec.Command("cmd", "set", "http_proxy=")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
	c = exec.Command("cmd", "set", "https_proxy=")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
