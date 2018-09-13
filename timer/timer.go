package timer

import (
	"fmt"
	"github.com/sanguohot/medichain/service"
	"time"
)

func initFileAddLogTimer()  {
	ticks := time.NewTicker(1 * time.Minute)
	tick := ticks.C
	go func() {
		for _ = range tick {
			err := service.SetFileAddLogList()
			if err != nil {
				fmt.Printf("timer: FileAddLogTimer error %s\n", err.Error())
			}
		}
	}()
	fmt.Printf("timer: FileAddLogTimer start %v.\n", time.Now())
}
func init() {
	initFileAddLogTimer()
}