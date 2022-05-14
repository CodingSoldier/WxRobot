package pkg

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

// TODO 修改task
var entryID = 0

func InitTask() (cron.EntryID, error) {
	// 读取配置文件
	conf, err := GetConfig()
	if err != nil {
		fmt.Printf("\nconfig.yaml配置信息为 %+v\n", conf)
		return 0, err
	}
	task := cron.New(cron.WithSeconds())
	entryID, err := task.AddFunc(conf.CronTime, func() {
		msg := conf.MessageContent + conf.MessagePassword
		fmt.Println(time.Now(), "发送企业微信消息，内容是 ", msg)
		WxSend(conf)
	})
	if err != nil {
		fmt.Println("创建定时任务出错了，错误信息是 ", err.Error())
		return entryID, err
	}
	fmt.Println("定时任务创建成功")
	task.Start()
	return entryID, nil
}

func ChangeTask(config *Config) {

}
