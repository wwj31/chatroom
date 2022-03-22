package main

import (
	"chatroom/server/common/log"
	"flag"
	"os"
	"os/signal"
	"syscall"
)

var (
	appId = flag.String("type", "server", "server or client")
)

func main() {
	flag.Parse()
	log.Init(0, "./", "serLog", true)
	sys := startup()

	osSignal := make(chan os.Signal)
	signal.Notify(osSignal, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-osSignal
	sys.Stop()
	<-sys.CStop
}
