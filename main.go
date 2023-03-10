package main

import (
	"flag"
	"fmt"
        "embed"
	"os"
	"wsh/g"
	"wsh/http"
)
//go:embed web/dist/*
var F embed.FS


func main() {
	//cfg
	cfg := flag.String("c", "cfg.toml", "configuretion file")

	version := flag.Bool("v", false, "display Version")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	if g.Config().Debug {
		g.InitLog("debug")
	} else {
		g.InitLog("info")
	}

	// Root
	g.InitRootDir()

	// http
	http.Init(F)

	go http.Start()

	select {}
}
