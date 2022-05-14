package main

import (
	"WxRobot/pkg"
	"fmt"
)

func main() {
	isSuccess := pkg.RunTask()
	if isSuccess {
		fmt.Println("程序启动成功")
		fmt.Println("不可以关闭本窗口。如果关闭本窗口，程序将停止运行")
	} else {
		fmt.Println("程序启动失败")
	}

	// 阻塞
	select {}
}
