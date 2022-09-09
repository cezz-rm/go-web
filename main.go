package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"go-web/initialize"
)

func main() {
	initialize.InitLogger("./log", "go.log")

	initialize.InitConfig()

	router := initialize.InitRouter()
	if err := router.Run(":9999"); err != nil {
		fmt.Println(err)
		log.Error("start failed")
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Debug("quit")
}
