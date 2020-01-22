package main

import (
	"fmt"
	"github.com/OhYee/ait/server"
	"github.com/OhYee/ait/server/dir"
	"github.com/OhYee/ait/test/B/api"
	"github.com/OhYee/rainbow/color"
	"github.com/OhYee/rainbow/log"
	"time"
)

var (
	infoLogger  = log.New().SetSuffix(func(s string) string { return "\n" })
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

	server.SetServerInfo(server.NewInfo("directory", "127.0.0.1:50000", 0, 0))
	server.Start("C", "127.0.0.1:50002", 10)
	go dir.StartHeartBeatThread()
	<-time.After(time.Second * 5)
	server.Debug("5+6=?")
	rep, err := api.Sum(api.SumRequest{
		A: 5,
		B: 6,
	})
	fmt.Println(rep, err)
}
