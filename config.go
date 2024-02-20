package logx

import "os"

// Configuration for LogX.
type Config struct {
	// Disable debug logs from running.
	// Default: false
	IsDisableDebugLogs bool

	// Disable warn logs from running.
	// Default: false
	IsDisableWarnLogs bool

	// Disable standard logger from writing logs to data stream.
	// Default: false
	IsDisableStandardLogger bool

	// Disable colored text in standard logger messages.
	// Default: false
	IsDisableColors bool

	// Define time format in standard logger messages.
	// Default: "2006-01-02 15:04:05 -0700"
	TimeFormat string

	// Define output data stream for standard logger.
	// Default: os.Stderr
	Output *os.File
}

// Initialize fields, that were not redefined by user, with default
// values.
func (c *Config) InitEmptyFields() {
	if c.TimeFormat == EmptyString {
		c.TimeFormat = TimeFormat
	}

	if c.Output == nil {
		c.Output = Output()
	}
}
