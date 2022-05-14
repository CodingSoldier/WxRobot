package pkg

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

// 定时任务
func RunTask() bool {
	// 读取配置文件
	var config Config
	conf := config.GetConf()
	fmt.Printf("\nconfig.yaml配置信息为 %+v\n", conf)

	task := cron.New(cron.WithSeconds())
	_, err := task.AddFunc(conf.CronTime, func() {
		msg := conf.MessageContent + conf.MessagePassword
		fmt.Println(time.Now(), "发送企业微信消息，内容是 ", msg)
		WxSend(conf)
	})
	if err != nil {
		fmt.Println("创建定时任务出错了，错误信息是 ", err.Error())
		task.Start()
		return false
	}
	fmt.Println("定时任务创建成功")
	task.Start()
	return true
}
