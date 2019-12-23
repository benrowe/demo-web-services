package main

import (
	"github.com/benrowe/demo-web-services/cmd/cli/cmd"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt)

	cmd.Execute()

	<-stop
	log.Info("Shutting down")
}
