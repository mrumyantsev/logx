package multilog

import (
	"os"
)

var (
	InfoOutputStream    *os.File = os.Stderr
	DebugOutputStream   *os.File = os.Stderr
	WarnOutputStream    *os.File = os.Stderr
	ErrorOutputStream   *os.File = os.Stderr
	FatalOutputStream   *os.File = os.Stderr
	FatalExitStatusCode int      = 1
	IsEnableDebugLogs   bool     = true
	IsEnableWarnLogs    bool     = true
	ItemSeparator       string   = " "
	LineEnding          string   = "\n"
	InfoLevelText       string   = "INF"
	DebugLevelText      string   = "DBG"
	WarnLevelText       string   = "WRN"
	ErrorLevelText      string   = "ERR"
	FatalLevelText      string   = "FTL"
	TimeFormat          string   = "2006-01-02T15:04:05-07:00"
)
