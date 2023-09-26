package fastlog

import (
	"os"
	"time"
)

type LogWriter interface {
	WriteLog(datetime string, messageType string, message string) error
}

type logMessage struct {
	datetime    string
	messageType *string
	message     *string
}

const (
	errorWord                 string = ". error: "
	exitProhibitingStatusCode int    = 1337
)

var (
	stdoutFile          *os.File    = os.Stdout
	stderrFile          *os.File    = os.Stderr
	fileLogWriter       LogWriter   = nil
	databaseLogWriter   LogWriter   = nil
	logMsg              *logMessage = &logMessage{}
	IsEnableDebugLogs   bool        = true
	ItemSeparator       string      = " "
	LineEnding          string      = "\n"
	InfoMessageType     string      = "INF"
	DebugMessageType    string      = "DBG"
	ErrorMessageType    string      = "ERR"
	FatalMessageType    string      = "FTL"
	TimeFormat          string      = "2006-01-02T15:04:05-07:00"
	FatalExitStatusCode int         = 1
)

func SetFileLogWriter(w LogWriter) {
	fileLogWriter = w
}

func SetDatabaseLogWriter(w LogWriter) {
	databaseLogWriter = w
}

func Info(msg string) *logMessage {
	logMsg.datetime = time.Now().Format(TimeFormat)
	logMsg.messageType = &InfoMessageType
	logMsg.message = &msg

	stdoutFile.Write(
		[]byte(logMsg.datetime +
			ItemSeparator +
			*logMsg.messageType +
			ItemSeparator +
			*logMsg.message +
			LineEnding,
		),
	)

	return logMsg
}

func Debug(msg string) *logMessage {
	logMsg.datetime = time.Now().Format(TimeFormat)
	logMsg.messageType = &DebugMessageType
	logMsg.message = &msg

	stdoutFile.Write(
		[]byte(logMsg.datetime +
			ItemSeparator +
			*logMsg.messageType +
			ItemSeparator +
			*logMsg.message +
			LineEnding,
		),
	)

	return logMsg
}

func Error(desc string, err error) *logMessage {
	var (
		msg string = desc + errorWord + err.Error()
	)

	logMsg.datetime = time.Now().Format(TimeFormat)
	logMsg.messageType = &ErrorMessageType
	logMsg.message = &msg

	stderrFile.Write(
		[]byte(logMsg.datetime +
			ItemSeparator +
			*logMsg.messageType +
			ItemSeparator +
			*logMsg.message +
			LineEnding,
		),
	)

	return logMsg
}

func Fatal(desc string, err error) *logMessage {
	var (
		msg string = desc + errorWord + err.Error()
	)

	logMsg.datetime = time.Now().Format(TimeFormat)
	logMsg.messageType = &FatalMessageType
	logMsg.message = &msg

	stderrFile.Write(
		[]byte(logMsg.datetime +
			ItemSeparator +
			*logMsg.messageType +
			ItemSeparator +
			*logMsg.message +
			LineEnding,
		),
	)

	if FatalExitStatusCode != exitProhibitingStatusCode {
		logMsg.Exit(FatalExitStatusCode)
	}

	return logMsg
}

func (l *logMessage) WriteLogToFile() *logMessage {
	fileLogWriter.WriteLog(l.datetime, *l.messageType, *l.message)

	return l
}

func (l *logMessage) WriteLogToDatabase() *logMessage {
	databaseLogWriter.WriteLog(l.datetime, *l.messageType, *l.message)

	return l
}

func (l *logMessage) Exit(statusCode int) {
	os.Exit(statusCode)
}
