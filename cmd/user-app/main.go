package main

import (
	"flag"
)

var path = flag.String("c", "configs/app.yaml", "set config file path")

func init() {
	flag.Parse()
}

// @title User API
// @version 1.0.0
// @description User API
//
// @contact.name Sean Zheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space
//
// @license.name GPL-3.0
// @license.url https://spdx.org/licenses/GPL-3.0-only.html
//
// @BasePath /api
//
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	app, err := CreateApp(*path)
	if err != nil {
		panic(err)
	}

	if err = app.Start(); err != nil {
		panic(err)
	}

	app.AwaitSignal()
}
