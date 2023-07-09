package randomize

import (
	"math/rand"
	"time"
)

// RandomTimeSleep will sleep a random time between 0 and 1000 millisecond.
func RandomTimeSleep() {
	RandomTimeSleepFromMinToMax(0, 1000, time.Millisecond)
}

// RandomTimeSleepFromMinToMax will sleep a random time between min and max * unit times.
func RandomTimeSleepFromMinToMax(min, max int, unit time.Duration) {
	if min > max {
		min, max = max, min
	}
	if min == max {
		time.Sleep(time.Duration(min) * unit)
		return
	}
	randtime := rand.Intn(max-min) + min
	if randtime <= 0 {
		time.Sleep(0)
		return
	}
	time.Sleep(time.Duration(randtime) * unit)
}
