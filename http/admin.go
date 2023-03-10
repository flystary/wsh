package http

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func admin(router *gin.Engine) {
	router.GET("/exit", func(c *gin.Context) {
	    c.Writer.Write([]byte("exiting"))
	    go func() {
			time.Sleep(time.Second)
			os.Exit(0)
	    }()
	})

}
