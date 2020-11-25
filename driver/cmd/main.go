package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/domarcio/bexs/config"
	commonLog "github.com/domarcio/bexs/src/common/log"
)

func main() {
	log := commonLog.NewLogfile(config.Logfile, "[BEXS] ", log.LstdFlags|log.Lmicroseconds|log.Llongfile)

	log.Info("Running cli interface on `%s` environment", config.Env)

	// Waiting for CTRL+C
	sg := make(chan os.Signal, 1)
	signal.Notify(sg, os.Interrupt)

	<-sg
	log.Info("Finish cli")
}
