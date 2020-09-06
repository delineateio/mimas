package server

import (
	"time"

	config "github.com/delineateio/mimas/config"
)

const defaultRead = 5 * time.Second
const defaultWrite = 5 * time.Second
const defaultHammer = time.Minute

// TimeOuts represents the server timeout config
type timeOuts struct {
	read   time.Duration
	write  time.Duration
	hammer time.Duration
}

func newTimeOuts() timeOuts {
	return timeOuts{
		read:   config.GetDuration("server.timeouts.read", defaultRead),
		write:  config.GetDuration("server.timeouts.write", defaultWrite),
		hammer: config.GetDuration("server.timeouts.hammer", defaultHammer),
	}
}
