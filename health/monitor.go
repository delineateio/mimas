package health

import (
	"time"

	h "github.com/InVisionApp/go-health/v2"
	log "github.com/delineateio/mimas/log"
)

// NewMonitor access the monitors
func NewMonitor() *Monitor {
	return &Monitor{}
}

// Monitor is the wrapper around the health checks
type Monitor struct {
	health *h.Health
}

// MonitorStatus represents the current status of the dependencies
type MonitorStatus struct {
	IsMonitoring bool
	Failed       bool
}

// AddCheck adds checkers that implement IChecker dynamically to list
// TODO: This is bleeding the underlying interface
func (m *Monitor) AddCheck(name string, interval time.Duration, fatal bool, config h.Config) {
	// Creates the check if required
	if m.health == nil {
		m.health = h.New()
		m.health.DisableLogging()
	} else {
		err := m.health.Stop()
		if err != nil {
			log.Warn("healthcheck.stop.error", "There was an issue stopping the health checks but it's not terminal")
		}
	}

	err := m.health.AddCheck(&config)
	if err != nil {
		log.Error("healthcheck.add.error", err)
	}

	err = m.health.Start()
	if err != nil {
		log.Error("healthcheck.start.error", err)
	}
}

// GetStatus returns the status of the monitor
func (m *Monitor) GetStatus() MonitorStatus {
	return MonitorStatus{
		IsMonitoring: true,
		Failed:       false,
	}
}
