package example

import (
	"fans-lucky-draw/logic"
	"testing"
)

var (
	csdnTotal   = int64(1022)
	weixinTotal = int64(392)
	circle      = "April.2022"
)

func TestDrawFunc(t *testing.T) {
	for i := 0; i < 100; i++ {
		logic.Execute(csdnTotal, weixinTotal, circle)
	}
}

