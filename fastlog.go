package fastlog

import (
	"os"
	"time"
)

type ILogWriter interface {
	WriteLog(datetime string, messageType string, message string) error
}

type logMessage struct {
	datetime    string
	messageType *string
	message     *string
}

const (
	WHITESPACE         string = " "
	NEW_LINE           string = "\n"
	ERROR_WORD         string = ". error: "
	INFO_MESSAGE_TYPE  string = "INF"
	DEBUG_MESSAGE_TYPE string = "DBG"
	ERROR_MESSAGE_TYPE string = "ERR"
	FATAL_MESSAGE_TYPE string = "FTL"
	TIME_FORMAT        string = "2006-01-02T15:04:05-07:00"
)

var (
	stdoutFile        *os.File    = os.Stdout
	stderrFile        *os.File    = os.Stderr
	fileLogWriter     ILogWriter  = nil
	databaseLogWriter ILogWriter  = nil
	IsEnableDebugLogs bool        = true
	ItemSeparator     string      = WHITESPACE
	LineEnding        string      = NEW_LINE
	InfoMessageType   string      = INFO_MESSAGE_TYPE
	DebugMessageType  string      = DEBUG_MESSAGE_TYPE
	ErrorMessageType  string      = ERROR_MESSAGE_TYPE
	FatalMessageType  string      = FATAL_MESSAGE_TYPE
	TimeFormat        string      = TIME_FORMAT
	logMsg            *logMessage = &logMessage{}
)

func SetFileLogWriter(w ILogWriter) {
	fileLogWriter = w
}

func SetDatabaseLogWriter(w ILogWriter) {
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
		msg string = desc + ERROR_WORD + err.Error()
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
		msg string = desc + ERROR_WORD + err.Error()
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
