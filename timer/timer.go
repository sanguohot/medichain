package timer

import (
	"github.com/sanguohot/medichain/service"
	"github.com/sanguohot/medichain/zap"
	"time"
)

func initFileAddLogTimer()  {
	ticks := time.NewTicker(1 * time.Minute)
	tick := ticks.C
	go func() {
		for _ = range tick {
			err := service.SetFileAddLogList()
			if err != nil {
				zap.Sugar.Error(err.Error())
			}
		}
	}()
}
func init() {
	initFileAddLogTimer()
}