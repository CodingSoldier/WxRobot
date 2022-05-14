package pkg

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// config.yaml配置文件载体
type Config struct {
	WebhookUrl        string `yaml:"webhook-url"`
	MessageTitle      string `yaml:"message-title"`
	MessageTitleColor string `yaml:"message-title-color"`
	MessagePassword   string `yaml:"message-password"`
	MessageContent    string `yaml:"message-content"`
	CronTime          string `yaml:"cron-time"`
}

func (c *Config) GetConf() *Config {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println("读取config.yaml文件错误，错误信息 ", err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println("config.yaml文件填写错误，错误信息 ", err.Error())
	}
	return c
}
