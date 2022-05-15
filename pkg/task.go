package pkg

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

type Task struct {
	Cron    *cron.Cron
	EntryID cron.EntryID
}

func (t *Task) InitTask() error {
	// 读取配置文件
	config, err := GetConfig()
	if err != nil {
		fmt.Printf("\nconfig.yaml配置信息为 %+v\n", config)
		return err
	}
	wxCron := cron.New(cron.WithSeconds())
	entryID, err := wxCron.AddFunc(config.CronTime, func() {
		msg := config.MessageContent + config.MessagePassword
		fmt.Println(time.Now(), "发送企业微信消息，内容是 ", msg)
		WxSend(config)
	})
	if err != nil {
		fmt.Println("创建定时任务出错了，错误信息是 ", err.Error())
		return err
	}
	fmt.Println("定时任务创建成功")
	wxCron.Start()
	t.EntryID = entryID
	t.Cron = wxCron
	return nil
}

func (t *Task) ChangeTask(config *Config) error {
	entryID, err := t.Cron.AddFunc(config.CronTime, func() {
		msg := config.MessageContent + config.MessagePassword
		fmt.Println(time.Now(), "发送企业微信消息，内容是 ", msg)
		WxSend(config)
	})
	if err != nil {
		fmt.Println("创建定时任务出错了，错误信息是 ", err.Error())
		return err
	}
	t.Cron.Remove(t.EntryID)
	t.Cron.Entry(entryID)
	return nil
}
