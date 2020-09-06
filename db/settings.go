package db

import (
	"time"

	config "github.com/delineateio/mimas/config"
)

const maxIdle = 20
const maxOpen = 50
const maxLifetime = 60 * time.Second
const attemptsInt = 3
const attempts = uint(attemptsInt)
const delay = 500 * time.Millisecond

// Settings for the database connection
type Settings struct {
	MaxIdle     int
	MaxOpen     int
	MaxLifetime time.Duration
	Attempts    uint
	Delay       time.Duration
}

// NewSettings creates the settings
func NewSettings(name string) *Settings {
	return &Settings{
		MaxIdle:     config.GetInt("db."+name+".limits.maxIdle", maxIdle),
		MaxOpen:     config.GetInt("db."+name+".limits.maxOpen", maxOpen),
		MaxLifetime: config.GetDuration("db."+name+".limits.maxLifetime", maxLifetime),
		Attempts:    config.GetUint("db."+name+".retries.attempts", attempts),
		Delay:       config.GetDuration("db."+name+".retries.delay", delay),
	}
}
