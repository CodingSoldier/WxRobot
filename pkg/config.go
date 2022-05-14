package pkg

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const configFile = "./static/config.yaml"

// config.yaml配置文件载体
type Config struct {
	WebhookUrl        string `yaml:"webhook-url"`
	MessageTitle      string `yaml:"message-title"`
	MessageTitleColor string `yaml:"message-title-color"`
	MessagePassword   string `yaml:"message-password"`
	MessageContent    string `yaml:"message-content"`
	CronTime          string `yaml:"cron-time"`
}

func GetConfig() (*Config, error) {
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("读取config.yaml文件错误，错误信息 ", err.Error())
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		fmt.Println("config.yaml文件填写错误，错误信息 ", err.Error())
		return config, err
	}
	return config, nil
}

func (c *Config) WriteConf(config *Config) error {
	// 第二个表示每行的前缀，这里不用，第三个是缩进符号，这里用tab
	data, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println("配置信息转换yaml格式失败，", err.Error())
		return err
	}
	err = ioutil.WriteFile(configFile, data, 0777)
	if err != nil {
		fmt.Println("保存配置失败，", err.Error())
		return err
	}
	return nil
}
