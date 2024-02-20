package logx

import "time"

// A LogWriter interface, to create custom log writer, which can be
// used by LogX to call WriteLog method with log message parameters.
type LogWriter interface {
	WriteLog(time time.Time, level LogLevel, msg string) error
}
