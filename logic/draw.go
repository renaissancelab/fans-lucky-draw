package logic

import (
	"crypto/rand"
	"github.com/scorpiotzh/mylog"
	"math"
	"math/big"
)

var (
	log         = mylog.NewLoggerDefault("logic", mylog.LevelInfo, nil)
)

func Execute(csdnTotal, weixinTotal int64, circle string) {
	//Step1: random select platform source,0-csdn 1-weixin
	platformInt := RangeRand(1, weixinTotal+csdnTotal)
	platformIndex := 1
	platformStr := "weixin"
	if platformInt > weixinTotal {
		platformIndex = 0
		platformStr = "csdn"
	}
	log.Debug("Random platformInt:", platformInt, "platformIndex:", platformIndex, "platformStr:", platformStr)
	var drawTotal int64
	//Step2: random select platform source,0-csdn 1-weixin
	if platformIndex == 0 {
		drawTotal = csdnTotal
	} else {
		drawTotal = weixinTotal
	}
	seqInt := RangeRand(0, drawTotal)

	log.Info("Congratulations！User Number:", platformStr, "-", seqInt, "Get", circle, "Reward!")
}

func RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}
