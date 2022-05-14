package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"net/http"
)

//go:embed static
var FS embed.FS

func main() {
	router := gin.Default()

	html := template.Must(template.New("").ParseFS(FS, "static/*.html"))
	router.SetHTMLTemplate(html)

	fsSub, _ := fs.Sub(FS, "static")
	router.StaticFS("/static", http.FS(fsSub))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/test", func(c *gin.Context) {
		firstName := c.Query("firstName")
		lastName := c.DefaultQuery("lastName", "默认值")

		c.String(http.StatusOK, "%s%s", firstName, lastName)
	})

	router.Run(":8080")

}
