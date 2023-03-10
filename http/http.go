package http

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"time"
	"wsh/g"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)


var (
	f          embed.FS
	username   string
	password   string
	savePass   bool

	timeout    int
	port       int64
	host       string
)

var (
	httpConfig *g.Http
        authConfig *g.Auth
)

func Init(fs embed.FS) {
	// 前端地址
	f = fs

	// Auth
	authConfig = g.Config().Auth
	username   = authConfig.Username
	password   = authConfig.Password

	httpConfig = g.Config().Http
	host    = httpConfig.HOST
	port    = httpConfig.PORT
	timeout = httpConfig.Timeout
}

func Start() {
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	// webssh
        admin(engine)
	static(engine)
	terminal(engine)
	file(engine)

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", host, port),
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatalln(server.ListenAndServe())
}
