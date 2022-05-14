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
	WebhookUrl        string `form:"webhookUrl"`
	MessageTitle      string `form:"messageTitle"`
	MessageTitleColor string `form:"messageTitleColor"`
	MessagePassword   string `form:"messagePassword"`
	MessageContent    string `form:"messageContent"`
	CronTime          string `form:"cronTime"`
}

func saveConfig(c *gin.Context) {
	var configDto ConfigDto
	if err := c.ShouldBind(&configDto); err == nil {
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
			c.JSON(http.StatusOK, Success())
		} else {
			c.JSON(http.StatusOK, Fail(writeErr.Error()))
		}

	} else {
		c.JSON(http.StatusOK, Fail(err.Error()))
	}
}

func getConfig(c *gin.Context) {
	conf, err := pkg.GetConfig()
	if err != nil {
		c.JSON(http.StatusOK, SuccessData(conf))
	} else {
		c.JSON(http.StatusOK, Fail(err.Error()))
	}
}

func GinStart(FS *embed.FS) {
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
