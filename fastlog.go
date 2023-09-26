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
	errorWord string = ". error: "
)

var (
	stdoutFile        *os.File              = os.Stdout
	stderrFile        *os.File              = os.Stderr
	writers           *map[string]LogWriter = nil
	writer            LogWriter             = nil
	isWriterExists    bool                  = false
	writersBeforeExit int                   = 0
	logMsg            *logMessage           = &logMessage{}

	IsEnableDebugLogs       bool   = true
	ItemSeparator           string = " "
	LineEnding              string = "\n"
	InfoMessageType         string = "INF"
	DebugMessageType        string = "DBG"
	ErrorMessageType        string = "ERR"
	FatalMessageType        string = "FTL"
	TimeFormat              string = "2006-01-02T15:04:05-07:00"
	ExitStatusCodeWhenFatal int    = 1
)

func RegisterWriter(name string, w LogWriter) {
	if writers == nil {
		writers = &map[string]LogWriter{}
	}

	(*writers)[name] = w

	writersBeforeExit++
}

func UnregisterWriter(name string) {
	delete((*writers), name)

	writersBeforeExit--
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
	if !IsEnableDebugLogs {
		return nil
	}

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

	if writersBeforeExit <= 0 {
		os.Exit(ExitStatusCodeWhenFatal)
	}

	return logMsg
}

func (l *logMessage) WriteTo(writerName string) *logMessage {
	if (l == nil) || (writers == nil) {
		return nil
	}

	writer, isWriterExists = (*writers)[writerName]
	if !isWriterExists {
		return nil
	}

	writer.WriteLog(l.datetime, *l.messageType, *l.message)

	writersBeforeExit--

	if writersBeforeExit <= 0 {
		os.Exit(ExitStatusCodeWhenFatal)
	}

	return l
}
