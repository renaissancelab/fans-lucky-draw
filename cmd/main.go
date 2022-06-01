package main

import (
	"fans-lucky-draw/logic"
	"github.com/scorpiotzh/mylog"
)

var (
	log         = mylog.NewLoggerDefault("main", mylog.LevelInfo, nil)
	csdnTotal   = int64(1252)
	weixinTotal = int64(388)
	circle      = "May.2022"
)

func main() {
	log.Info("Start draw", circle, "Just wait")
	logic.Execute(csdnTotal, weixinTotal, circle)
}
