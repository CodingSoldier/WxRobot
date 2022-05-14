package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// 发送企业微信消息
func WxSend(config *Config) {
	body := make(map[string]interface{})
	body["msgtype"] = "markdown"
	var content = fmt.Sprintf("# <font color=\"%s\">%s</font> \n %s%s",
		config.MessageTitleColor, config.MessageTitle, config.MessageContent, config.MessagePassword)
	body["markdown"] = map[string]string{
		"content": content,
	}
	b, _ := json.Marshal(body)
	http.DefaultClient.Post(config.WebhookUrl, "application/json", bytes.NewBuffer(b))
}
