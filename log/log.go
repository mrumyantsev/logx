package log

import (
	"fmt"
	"os"
	"time"

	"github.com/mrumyantsev/go-idmap"
	"github.com/mrumyantsev/multilog"
)

const (
	errorWord string = ". error: "
)

var (
	logWriters   *idmap.IdMap = nil
	logWriterErr error
)

type LogWriter interface {
	WriteLog(datetime time.Time, level string, message string) error
}

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
		&multilog.InfoLevelText,
		&msg,
		multilog.InfoOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&multilog.InfoLevelText,
			&msg,
		)
	}
}

func Debug(msg string) {
	if !multilog.IsEnableDebugLogs {
		return
	}

	datetime := time.Now()

	writeToStream(
		&datetime,
		&multilog.DebugLevelText,
		&msg,
		multilog.DebugOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&multilog.DebugLevelText,
			&msg,
		)
	}
}

func Warn(msg string) {
	if !multilog.IsEnableWarnLogs {
		return
	}

	datetime := time.Now()

	writeToStream(
		&datetime,
		&multilog.WarnLevelText,
		&msg,
		multilog.WarnOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&multilog.WarnLevelText,
			&msg,
		)
	}
}

func Error(desc string, err error) {
	datetime := time.Now()

	desc = desc + errorWord + err.Error()

	writeToStream(
		&datetime,
		&multilog.ErrorLevelText,
		&desc,
		multilog.ErrorOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&multilog.ErrorLevelText,
			&desc,
		)
	}
}

func Fatal(desc string, err error) {
	datetime := time.Now()

	desc = desc + errorWord + err.Error()

	writeToStream(
		&datetime,
		&multilog.FatalLevelText,
		&desc,
		multilog.FatalOutputStream,
	)

	if logWriters != nil {
		writeToLogWriters(
			&datetime,
			&multilog.FatalLevelText,
			&desc,
		)
	}

	os.Exit(multilog.FatalExitStatusCode)
}

func writeToStream(
	datetime *time.Time,
	level *string,
	message *string,
	stream *os.File,
) {
	stream.Write(
		[]byte(
			(*datetime).Format(multilog.TimeFormat) +
				multilog.ItemSeparatorText +
				*level +
				multilog.ItemSeparatorText +
				*message +
				multilog.LineEndingText,
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
			*datetime,
			*level,
			*message,
		)

		if logWriterErr != nil {
			Error(
				fmt.Sprintf("could not write to log writer with id=%d", id),
				logWriterErr,
			)
		}
	}
}
