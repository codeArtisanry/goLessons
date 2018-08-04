package tick

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

type Conf struct {
	IntervalMs time.Duration
	Allbytes   uint64
	Allrows    uint64
}

var bs uint64 = 10
var args = Conf{IntervalMs: 5, Allbytes: 10000000, Allrows: 15}

func TestTick(tt *testing.T) {
	t := time.Now()
	tick := time.NewTicker(time.Millisecond * time.Duration(args.IntervalMs))
	defer tick.Stop()
	go func() {
		for range tick.C {
			diff := time.Since(t).Seconds()
			allbytesMB := float64(atomic.LoadUint64(&args.Allbytes) / 1024 / 1024)
			allrows := atomic.LoadUint64(&args.Allrows)
			rates := allbytesMB / diff
			fmt.Println(allbytesMB, rates, allrows)
			// log.Info("dumping.allbytes[%vMB].allrows[%v].time[%.2fsec].rates[%.2fMB/sec]...", allbytesMB, allrows, diff, rates)
		}
	}()
	time.Sleep(1 * time.Second)

}
