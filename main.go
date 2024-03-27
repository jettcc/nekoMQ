package main

import (
	"flag"
	"sesm/global"
	"sesm/internal/load"
)

func main() {
	var logFilePath = flag.String("l", "log/service.log", "log file path")
	flag.Parse()

	load.Init(*logFilePath)
	defer load.Sync()

	global.Sugar.Info("Service start")
}
