package main

import (
	"github.com/OhYee/ait/server"
	_ "github.com/OhYee/ait/test/B/api"
	"github.com/OhYee/rainbow/color"
	"github.com/OhYee/rainbow/log"
)

var (
	infoLogger = log.New().SetSuffix(func(s string) string { return "\n" })
	debugLogger = log.New().SetColor(color.New().SetFrontYellow()).SetSuffix(func(s string) string { return "\n" })
	errLogger   = log.New().SetColor(color.New().SetFrontRed().SetFontBold())
)

func main() {
	server.SetLogCallback(func(format string, args ...interface{}) {
		infoLogger.Printf(format, args...)
	})
	server.SetDebugCallback(func(format string, args ...interface{}) {
		debugLogger.Printf(format, args...)
	})
	server.SetErrCallback(func(err error) {
		errLogger.Println(err)
	})

	server.Start("127.0.0.1:50001", 10)
	for {
	}
}
