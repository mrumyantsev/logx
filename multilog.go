package multilog

import (
	"fmt"
	"os"
	"time"
)

type LogWriter interface {
	WriteLog(datetime *string, messageType *string, message *string) error
}

const (
	errorWord string = ". error: "
)

var (
	logWriters   map[int]LogWriter = nil
	logWriterErr error

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
	InfoMessageType     string   = "INF"
	DebugMessageType    string   = "DBG"
	WarnMessageType     string   = "WRN"
	ErrorMessageType    string   = "ERR"
	FatalMessageType    string   = "FTL"
	TimeFormat          string   = "2006-01-02T15:04:05-07:00"
)

func RegisterWriter(id int, writer LogWriter) {
	if logWriters == nil {
		logWriters = map[int]LogWriter{}
	}

	logWriters[id] = writer
}

func UnregisterWriter(id int) {
	delete(logWriters, id)
}

func Info(msg string) {
	datetime := time.Now().Format(TimeFormat)

	writeToStream(
		&datetime,
		&InfoMessageType,
		&msg,
		InfoOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&InfoMessageType,
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

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&DebugMessageType,
			&msg,
		)
	}
}

func Warn(msg string) {
	if !IsEnableWarnLogs {
		return
	}

	datetime := time.Now().Format(TimeFormat)

	writeToStream(
		&datetime,
		&WarnMessageType,
		&msg,
		WarnOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&WarnMessageType,
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

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&ErrorMessageType,
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

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&FatalMessageType,
			&desc,
		)
	}

	os.Exit(FatalExitStatusCode)
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
	for id, writer := range logWriters {
		logWriterErr = writer.WriteLog(
			datetime,
			messageType,
			message,
		)

		if logWriterErr != nil {
			Error(
				fmt.Sprintf("could not write to log writer with id=%d", id),
				logWriterErr,
			)
		}
	}
}
