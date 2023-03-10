package http

import (
	"wsh/http/controller"

	"github.com/gin-gonic/gin"
)


func file(router *gin.Engine) {

	file := router.Group("/file")
	{
		file.GET("/list", func(c *gin.Context) {
			c.JSON(200, controller.FileList(c))
		})
		file.GET("/download", func(c *gin.Context) {
			controller.DownloadFile(c)
		})
		file.POST("/upload", func(c *gin.Context) {
			c.JSON(200, controller.UploadFile(c))
		})
		file.GET("/progress", func(c *gin.Context) {
			controller.UploadProgressWs(c)
		})
	}

}
