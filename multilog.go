package multilog

import (
	"os"
)

// Multilog configurational structure.
type Config struct {
	InfoOutputStream  *os.File
	DebugOutputStream *os.File
	WarnOutputStream  *os.File
	ErrorOutputStream *os.File
	FatalOutputStream *os.File

	IsDisableDebugLogs bool
	IsDisableWarnLogs  bool

	FatalExitStatusCode int

	TimeFormat string

	ItemSeparatorText string
	LineEndingText    string
	InfoLevelText     string
	DebugLevelText    string
	WarnLevelText     string
	ErrorLevelText    string
	FatalLevelText    string
}
