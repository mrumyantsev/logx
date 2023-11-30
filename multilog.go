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
	writers map[string]LogWriter = nil

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

func RegisterWriter(name string, lw LogWriter) {
	if writers == nil {
		writers = map[string]LogWriter{}
	}

	writers[name] = lw
}

func UnregisterWriter(name string) {
	delete(writers, name)
}

func Info(msg string) {
	datetime := time.Now().Format(TimeFormat)

	writeToStream(
		&datetime,
		&InfoMessageType,
		&msg,
		InfoOutputStream,
	)

	if writers != nil {
		writeToLogWriters(
			&datetime,
			&FatalMessageType,
			&msg,
		)
	}
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

	if writers != nil {
		writeToLogWriters(
			&datetime,
			&FatalMessageType,
			&msg,
		)
	}
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

	if writers != nil {
		writeToLogWriters(
			&datetime,
			&FatalMessageType,
			&desc,
		)
	}
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

	if writers != nil {
		writeToLogWriters(
			&datetime,
			&FatalMessageType,
			&desc,
		)
	}

	os.Exit(ExitStatusCodeWhenFatal)
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

func writeToLogWriters(
	datetime *string,
	messageType *string,
	message *string,
) {
	for _, writer := range writers {
		writer.WriteLog(
			*datetime,
			*messageType,
			*message,
		)
	}
}
