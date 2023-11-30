package multilog

import (
	"os"
	"time"
)

type LogWriter interface {
	WriteLog(datetime string, messageType string, message string) error
}

const (
	errorWord string = ". error: "
)

var (
	writers           *map[string]LogWriter = nil
	writer            LogWriter             = nil
	writersBeforeExit int                   = 0
	isWriterExists    bool                  = false
	isFatalLog        bool                  = false

	InfoOutputStream        *os.File = os.Stderr
	DebugOutputStream       *os.File = os.Stderr
	ErrorOutputStream       *os.File = os.Stderr
	FatalOutputStream       *os.File = os.Stderr
	ExitStatusCodeWhenFatal int      = 1
	IsEnableDebugLogs       bool     = true
	ItemSeparator           string   = " "
	LineEnding              string   = "\n"
	InfoMessageType         string   = "INF"
	DebugMessageType        string   = "DBG"
	ErrorMessageType        string   = "ERR"
	FatalMessageType        string   = "FTL"
	TimeFormat              string   = "2006-01-02T15:04:05-07:00"
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

func Info(msg string) {
	datetime := time.Now().Format(TimeFormat)

	writeToStream(
		&datetime,
		&InfoMessageType,
		&msg,
		InfoOutputStream,
	)
}

func Debug(msg string) {
	if !IsEnableDebugLogs {
		return
	}

	datetime := time.Now().Format(TimeFormat)

	writeToStream(
		&datetime,
		&DebugMessageType,
		&msg,
		DebugOutputStream,
	)
}

func Error(desc string, err error) {
	datetime := time.Now().Format(TimeFormat)

	desc = desc + errorWord + err.Error()

	writeToStream(
		&datetime,
		&ErrorMessageType,
		&desc,
		ErrorOutputStream,
	)
}

func Fatal(desc string, err error) {
	datetime := time.Now().Format(TimeFormat)

	desc = desc + errorWord + err.Error()

	writeToStream(
		&datetime,
		&FatalMessageType,
		&desc,
		FatalOutputStream,
	)

	if writersBeforeExit <= 0 {
		os.Exit(ExitStatusCodeWhenFatal)
	}

	isFatalLog = true
}

func writeToStream(
	datetime *string,
	messageType *string,
	message *string,
	stream *os.File,
) {
	stream.Write(
		[]byte(
			*datetime +
				ItemSeparator +
				*messageType +
				ItemSeparator +
				*message +
				LineEnding,
		),
	)
}

func (l *logMessage) WriteTo(writerName string) *logMessage {
	if (l == nil) || (writers == nil) {
		return nil
	}

	writer, isWriterExists = (*writers)[writerName]
	if !isWriterExists {
		return l
	}

	writer.WriteLog(l.datetime, *l.messageType, *l.message)

	if isFatalLog {
		writersBeforeExit--

		if writersBeforeExit <= 0 {
			os.Exit(ExitStatusCodeWhenFatal)
		}
	}

	return l
}
