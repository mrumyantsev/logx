package multilog

import (
	"fmt"
	"os"
	"time"

	"github.com/mrumyantsev/go-idmap"
)

type LogWriter interface {
	WriteLog(datetime *time.Time, level *string, message *string) error
}

const (
	errorWord string = ". error: "
)

var (
	logWriters   *idmap.IdMap = nil
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
	InfoLevel           string   = "INF"
	DebugLevel          string   = "DBG"
	WarnLevel           string   = "WRN"
	ErrorLevel          string   = "ERR"
	FatalLevel          string   = "FTL"
	TimeFormat          string   = "2006-01-02T15:04:05-07:00"
)

func AddLogWriter(id int, writer LogWriter) {
	if logWriters == nil {
		logWriters = idmap.New()
	}

	logWriters.SetValue(id, writer)
}

func RemoveLogWriter(id int) {
	logWriters.DeleteValue(id)
}

func EnableLogWriter(id int) {
	logWriters.Enable(id)
}

func DisableLogWriter(id int) {
	logWriters.Disable(id)
}

func Info(msg string) {
	datetime := time.Now()

	writeToStream(
		&datetime,
		&InfoLevel,
		&msg,
		InfoOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&InfoLevel,
			&msg,
		)
	}
}

func Debug(msg string) {
	if !IsEnableDebugLogs {
		return
	}

	datetime := time.Now()

	writeToStream(
		&datetime,
		&DebugLevel,
		&msg,
		DebugOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&DebugLevel,
			&msg,
		)
	}
}

func Warn(msg string) {
	if !IsEnableWarnLogs {
		return
	}

	datetime := time.Now()

	writeToStream(
		&datetime,
		&WarnLevel,
		&msg,
		WarnOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&WarnLevel,
			&msg,
		)
	}
}

func Error(desc string, err error) {
	datetime := time.Now()

	desc = desc + errorWord + err.Error()

	writeToStream(
		&datetime,
		&ErrorLevel,
		&desc,
		ErrorOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&ErrorLevel,
			&desc,
		)
	}
}

func Fatal(desc string, err error) {
	datetime := time.Now()

	desc = desc + errorWord + err.Error()

	writeToStream(
		&datetime,
		&FatalLevel,
		&desc,
		FatalOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&FatalLevel,
			&desc,
		)
	}

	os.Exit(FatalExitStatusCode)
}

func writeToStream(
	datetime *time.Time,
	level *string,
	message *string,
	stream *os.File,
) {
	stream.Write(
		[]byte(
			(*datetime).Format(TimeFormat) +
				ItemSeparator +
				*level +
				ItemSeparator +
				*message +
				LineEnding,
		),
	)
}

func writeToLogWriters(
	datetime *time.Time,
	level *string,
	message *string,
) {
	var (
		length    = logWriters.GetLength()
		writer    interface{}
		isEnabled bool
	)

	for id := 0; id < length; id++ {
		writer, isEnabled = logWriters.GetValue(id)

		if !isEnabled {
			continue
		}

		logWriterErr = writer.(LogWriter).WriteLog(
			datetime,
			level,
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
