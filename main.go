package main

import (
	"flag"
	"sesm/global"
	"sesm/internal/load"
	"sesm/internal/utils"
)

var logg = global.Sugar

func main() {
	var logFilePath = flag.String("l", "log/service.log", "log file path")
	flag.Parse()

	load.Init(*logFilePath)
	defer load.Sync()
	global.Sugar.Info("Service start")
	go utils.UuidFactory()
}
