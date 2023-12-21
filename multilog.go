package multilog

import (
	"os"

	"github.com/mrumyantsev/multilog/defaults"
)

// Multilog configurational structure.
type Config struct {
	// Disables debug logs to be executed.
	// Default: false
	IsDisableDebugLogs bool

	// Disables warn logs to be executed.
	// Default: false
	IsDisableWarnLogs bool

	// Disables colored text in the stream logs.
	// Default: false
	IsDisableColors bool

	// Defines datetime format in the stream logs.
	// Default: "2006-01-02T15:04:05-07:00"
	TimeFormat string

	// Chooses output data stream for the stream logs.
	// Default: os.Stderr
	OutputStream *os.File
}

// Initialize fields, that were not set by user, with its default values.
func (c *Config) InitEmptyFields() {
	if c.TimeFormat == defaults.EMPTY_STRING {
		c.TimeFormat = defaults.TIME_FORMAT
	}

	if c.OutputStream == nil {
		c.OutputStream = defaults.GetOutputStream()
	}
}
