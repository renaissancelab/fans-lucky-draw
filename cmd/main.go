package main

import (
	"fans-lucky-draw/logic"
	"github.com/scorpiotzh/mylog"
)

var (
	log         = mylog.NewLoggerDefault("main", mylog.LevelInfo, nil)
	csdnTotal   = int64(1022)
	weixinTotal = int64(392)
	circle      = "April.2022"
)

func main() {
	log.Info("Start draw, Just wait")
	logic.Execute(csdnTotal, weixinTotal, circle)
}

