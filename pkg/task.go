package pkg

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

// 定时任务
func RunTask() error {
	// 读取配置文件
	config := GetConf()
	fmt.Printf("\nconfig.yaml配置信息为 %+v\n", config)

	task := cron.New(cron.WithSeconds())
	_, err := task.AddFunc(config.CronTime, func() {
		msg := config.MessageContent + config.MessagePassword
		fmt.Println(time.Now(), "发送企业微信消息，内容是 ", msg)
		WxSend(config)
	})
	if err != nil {
		fmt.Println("【异常】创建定时任务出错了，错误信息是 ", err.Error())
		task.Start()
		return err
	}
	fmt.Println("定时任务创建成功")
	task.Start()
	return nil
}
