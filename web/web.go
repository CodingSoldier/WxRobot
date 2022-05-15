package web

import (
	"WxRobot/pkg"
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"net/http"
)

type ConfigDto struct {
	WebhookUrl        string `form:"WebhookUrl"`
	MessageTitle      string `form:"MessageTitle"`
	MessageTitleColor string `form:"MessageTitleColor"`
	MessagePassword   string `form:"MessagePassword"`
	MessageContent    string `form:"MessageContent"`
	CronTime          string `form:"CronTime"`
}

func saveConfig(c *gin.Context) {
	var configDto ConfigDto
	err := c.ShouldBind(&configDto)
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	}
	config := &pkg.Config{
		WebhookUrl:        configDto.WebhookUrl,
		MessageTitle:      configDto.MessageTitle,
		MessageTitleColor: configDto.MessageTitleColor,
		MessagePassword:   configDto.MessagePassword,
		MessageContent:    configDto.MessageContent,
		CronTime:          configDto.CronTime,
	}
	writeErr := config.WriteConf(config)
	if writeErr != nil {
		c.JSON(http.StatusOK, Fail(writeErr.Error()))
	}
	// 修改定时器
	taskErr := pkg.WxTask.ChangeTask(config)
	if taskErr != nil {
		c.JSON(http.StatusOK, Fail(taskErr.Error()))
	}
	c.JSON(http.StatusOK, Success())
}

func getConfig(c *gin.Context) {
	config, err := pkg.GetConfig()
	if err != nil {
		c.JSON(http.StatusOK, Fail(err.Error()))
	}
	dto := &ConfigDto{
		WebhookUrl:        config.WebhookUrl,
		MessageTitle:      config.MessageTitle,
		MessageTitleColor: config.MessageTitleColor,
		MessagePassword:   config.MessagePassword,
		MessageContent:    config.MessageContent,
		CronTime:          config.CronTime,
	}
	c.JSON(http.StatusOK, SuccessData(dto))
}

func GinStart(FS embed.FS) {
	router := gin.Default()
	html := template.Must(template.New("").ParseFS(FS, "static/*.html"))
	router.SetHTMLTemplate(html)

	fsSub, _ := fs.Sub(FS, "static")
	router.StaticFS("/static", http.FS(fsSub))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/api/config/detail", getConfig)

	router.POST("/api/config/save", saveConfig)

	router.Run(":10010")
}
