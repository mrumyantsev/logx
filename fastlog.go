package fastlog

import (
	"os"
	"time"
)

var (
	stdoutFile        *os.File = os.Stdout
	stderrFile        *os.File = os.Stderr
	IsEnableDebugLogs bool     = true
	InfoTitle         string   = " INF "
	DebugTitle        string   = " DBG "
	ErrorTitle        string   = " ERR "
	FatalTitle        string   = " FTL "
	TimeFormat        string   = "2006-01-02T15:04:05-07:00"
)

const (
	NEW_LINE   string = "\n"
	ERROR_WORD string = ". error: "
)

func Info(msg string) {
	stdoutFile.Write([]byte(time.Now().Format(TimeFormat) + InfoTitle + msg + NEW_LINE))
}

func Debug(msg string) {
	if IsEnableDebugLogs {
		stdoutFile.Write([]byte(time.Now().Format(TimeFormat) + DebugTitle + msg + NEW_LINE))
	}
}

func Error(desc string, err error) {
	stderrFile.Write([]byte(time.Now().Format(TimeFormat) + ErrorTitle + desc + ERROR_WORD + err.Error() + NEW_LINE))
}

func Fatal(desc string, err error) {
	stderrFile.Write([]byte(time.Now().Format(TimeFormat) + FatalTitle + desc + ERROR_WORD + err.Error() + NEW_LINE))
	os.Exit(1)
}
