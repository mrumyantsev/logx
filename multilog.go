package multilog

import (
	"os"
)

// Multilog configurational variables.
var (
	InfoOutputStream  *os.File = os.Stderr
	DebugOutputStream *os.File = os.Stderr
	WarnOutputStream  *os.File = os.Stderr
	ErrorOutputStream *os.File = os.Stderr
	FatalOutputStream *os.File = os.Stderr

	IsEnableDebugLogs bool = true
	IsEnableWarnLogs  bool = true

	FatalExitStatusCode int = 1

	TimeFormat string = "2006-01-02T15:04:05-07:00"

	ItemSeparatorText string = " "
	LineEndingText    string = "\n"
	InfoLevelText     string = "INF"
	DebugLevelText    string = "DBG"
	WarnLevelText     string = "WRN"
	ErrorLevelText    string = "ERR"
	FatalLevelText    string = "FTL"
)
