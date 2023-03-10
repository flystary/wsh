package http

import (
	"io/fs"
	"net/http"
	"os"
	"strconv"
	"time"
	"wsh/http/controller"

	"github.com/gin-gonic/gin"
)



func init() {
	envVal, ok := os.LookupEnv("savePass")
	if ok {
		b, err := strconv.ParseBool(envVal)
		if err != nil {
			savePass = false
		} else {
			savePass = b
		}
	}
}

func terminal(router *gin.Engine) {
	router.GET("/term", func(c *gin.Context) {
		controller.TermWs(c, time.Duration(timeout)*time.Minute)
	})
	router.GET("/check", func(c *gin.Context) {
		responseBody := controller.CheckSSH(c)
		responseBody.Data = map[string]interface{}{
			"savePass": false,
		}
		c.JSON(200, responseBody)
	})
}

func static(router *gin.Engine) {
	if username != "" && password != "" {
		accountList := map[string]string{
			username: password,
		}
		authorized := router.Group("/", gin.BasicAuth(accountList))
		authorized.GET("", func(c *gin.Context) {
			indexHTML, _ := f.ReadFile("web/dist/" + "index.html")
			c.Writer.Write(indexHTML)
		})
	} else {
		router.GET("/", func(c *gin.Context) {
			indexHTML, err := f.ReadFile("web/dist/" + "index.html")
			if err != nil {
				panic(err)
			}
			c.Writer.Write(indexHTML)
		})
	}
	staticFs, _ := fs.Sub(f, "web/dist/static")
	router.StaticFS("/static", http.FS(staticFs))
}
