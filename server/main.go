package main

import (
	"QrCodeLogin/internal/core"
	"QrCodeLogin/pkg/log"
	"flag"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	confPathPtr := flag.String("conf", "./configs", "configure file path")
	flag.Parse()

	closeFunc, err := core.InitApp(*confPathPtr)
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	log.Info("QrCodeLogin 后台服务已启动")
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("QrCodeLogin 后台服务正在关闭中....")
			closeFunc()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
