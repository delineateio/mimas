package containers

import (
	"time"

	config "github.com/delineateio/mimas/config"
)

const defaultRead = 5
const defaultWrite = 5

// TimeOuts represents the server timeout config
type timeOuts struct {
	read   time.Duration
	write  time.Duration
	hammer time.Duration
}

func newTimeOuts() timeOuts {
	return timeOuts{
		read:   config.GetDuration("server.timeouts.read", defaultRead*time.Second),
		write:  config.GetDuration("server.timeouts.write", defaultWrite*time.Second),
		hammer: config.GetDuration("server.timeouts.hammer", time.Minute),
	}
}
