package fastlog

import (
	"os"
	"time"
)

const (
	WHITESPACE string = " "
	TAB        string = "\t"
	NEW_LINE   string = "\n"
	ERROR_WORD string = ". error: "
)

var (
	stdoutFile        *os.File = os.Stdout
	stderrFile        *os.File = os.Stderr
	IsEnableDebugLogs bool     = true
	ItemSeparator     string   = WHITESPACE
	LineEnding        string   = NEW_LINE
	InfoItem          string   = "INF"
	DebugItem         string   = "DBG"
	ErrorItem         string   = "ERR"
	FatalItem         string   = "FTL"
	TimeFormat        string   = "2006-01-02T15:04:05-07:00"
)

func Info(msg string) {
	var (
		endMsg string = time.Now().Format(TimeFormat) +
			ItemSeparator +
			InfoItem +
			ItemSeparator +
			msg +
			LineEnding
	)

	stdoutFile.Write([]byte(endMsg))
}

func Debug(msg string) {
	var (
		endMsg string = time.Now().Format(TimeFormat) +
			ItemSeparator +
			DebugItem +
			ItemSeparator +
			msg +
			LineEnding
	)

	if IsEnableDebugLogs {
		stdoutFile.Write([]byte(endMsg))
	}
}

func Error(desc string, err error) {
	var (
		endMsg string = time.Now().Format(TimeFormat) +
			ItemSeparator +
			ErrorItem +
			ItemSeparator +
			desc +
			ERROR_WORD +
			err.Error() +
			LineEnding
	)

	stderrFile.Write([]byte(endMsg))
}

func Fatal(desc string, err error) {
	var (
		endMsg string = time.Now().Format(TimeFormat) +
			ItemSeparator +
			FatalItem +
			ItemSeparator +
			desc +
			ERROR_WORD +
			err.Error() +
			LineEnding
	)

	stderrFile.Write([]byte(endMsg))

	os.Exit(1)
}
