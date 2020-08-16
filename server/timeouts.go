package server

import (
	"time"

	c "github.com/delineateio/mimas/common"
	"github.com/fvbock/endless"
)

const defaultRead = 5
const defaultWrite = 5

// TimeOuts represents the server timeout config
type TimeOuts struct {
	Read   time.Duration
	Write  time.Duration
	Hammer time.Duration
}

func readTimeOuts() TimeOuts {
	return TimeOuts{
		Read:   c.GetDuration("server.timeouts.read", defaultRead*time.Second),
		Write:  c.GetDuration("server.timeouts.write", defaultWrite*time.Second),
		Hammer: c.GetDuration("server.timeouts.hammer", time.Minute),
	}
}

func updateTimeOuts(timeOuts TimeOuts) {
	endless.DefaultReadTimeOut = timeOuts.Read
	endless.DefaultWriteTimeOut = timeOuts.Write
	endless.DefaultHammerTime = timeOuts.Hammer
}

func setTimeOuts() {
	updateTimeOuts(readTimeOuts())
}
