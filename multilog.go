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
	InfoLevelText       string   = "INF"
	DebugLevelText      string   = "DBG"
	WarnLevelText       string   = "WRN"
	ErrorLevelText      string   = "ERR"
	FatalLevelText      string   = "FTL"
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
		&InfoLevelText,
		&msg,
		InfoOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&InfoLevelText,
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
		&DebugLevelText,
		&msg,
		DebugOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&DebugLevelText,
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
		&WarnLevelText,
		&msg,
		WarnOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&WarnLevelText,
			&msg,
		)
	}
}

func Error(desc string, err error) {
	datetime := time.Now()

	desc = desc + errorWord + err.Error()

	writeToStream(
		&datetime,
		&ErrorLevelText,
		&desc,
		ErrorOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&ErrorLevelText,
			&desc,
		)
	}
}

func Fatal(desc string, err error) {
	datetime := time.Now()

	desc = desc + errorWord + err.Error()

	writeToStream(
		&datetime,
		&FatalLevelText,
		&desc,
		FatalOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&FatalLevelText,
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
