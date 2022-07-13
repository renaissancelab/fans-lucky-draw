package main

import (
	"fans-lucky-draw/logic"
	"github.com/scorpiotzh/mylog"
)

var (
	log         = mylog.NewLoggerDefault("main", mylog.LevelInfo, nil)
	csdnTotal   = int64(1667)
	weixinTotal = int64(390)
	circle      = "June.2022"
)

func main() {
	log.Info("Start draw", circle, "Just wait")
	logic.Execute(csdnTotal, weixinTotal, circle)
}
