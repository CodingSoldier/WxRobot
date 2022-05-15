package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// 发送企业微信消息
func WxSend(config *Config) error {
	body := make(map[string]interface{})
	body["msgtype"] = "markdown"
	var content = fmt.Sprintf("# <font color=\"%s\">%s</font> \n %s%s",
		config.MessageTitleColor, config.MessageTitle, config.MessageContent, config.MessagePassword)
	body["markdown"] = map[string]string{
		"content": content,
	}
	b, _ := json.Marshal(body)
	resp, err := http.DefaultClient.Post(config.WebhookUrl, "application/json", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("【异常】发送企业微信消息出错 ", err.Error())
		return err
	}
	defer resp.Body.Close()
	bodyC, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("【异常】读取企业微信返回结果出错 ", err.Error())
		return err
	}
	var jsonMap map[string]interface{}
	err = json.Unmarshal(bodyC, &jsonMap)
	if err != nil {
		fmt.Println("【异常】反序列化企业微信返回结果出错 ", err.Error())
		return err
	}
	fmt.Println(time.Now(), "发送企业微信消息成功，企业微信返回报文是 ", jsonMap)
	return nil
}
