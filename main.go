package main

import (
	"WxRobot/pkg"
	"WxRobot/web"
	"embed"
)

//go:embed static
var FS embed.FS

func main() {
	pkg.WxTask.InitTask()
	web.GinStart(FS)

}
