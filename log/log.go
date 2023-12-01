package log

import (
	"fmt"
	"os"
	"time"

	"github.com/mrumyantsev/go-idmap"
	"github.com/mrumyantsev/multilog"
)

const (
	_ERROR_WORD string = ". error: "
)

var (
	writers   *idmap.IdMap = nil
	writerErr error        = nil
)

type Writer interface {
	WriteLog(datetime time.Time, level string, message string) error
}

func AddWriter(id int, writer Writer) {
	if writers == nil {
		writers = idmap.New()
	}

	writers.SetValue(id, writer)
}

func RemoveWriter(id int) {
	writers.DeleteValue(id)
}

func EnableWriter(id int) {
	writers.Enable(id)
}

func DisableWriter(id int) {
	writers.Disable(id)
}

func Info(msg string) {
	datetime := time.Now()

	writeToStream(
		&datetime,
		&multilog.InfoLevelText,
		&msg,
		multilog.InfoOutputStream,
	)

	if writers != nil {
		writeToWriters(
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

	if writers != nil {
		writeToWriters(
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

	if writers != nil {
		writeToWriters(
			&datetime,
			&multilog.WarnLevelText,
			&msg,
		)
	}
}

func Error(desc string, err error) {
	datetime := time.Now()

	desc = desc + _ERROR_WORD + err.Error()

	writeToStream(
		&datetime,
		&multilog.ErrorLevelText,
		&desc,
		multilog.ErrorOutputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&multilog.ErrorLevelText,
			&desc,
		)
	}
}

func Fatal(desc string, err error) {
	datetime := time.Now()

	desc = desc + _ERROR_WORD + err.Error()

	writeToStream(
		&datetime,
		&multilog.FatalLevelText,
		&desc,
		multilog.FatalOutputStream,
	)

	if writers != nil {
		writeToWriters(
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

func writeToWriters(
	datetime *time.Time,
	level *string,
	message *string,
) {
	var (
		length    = writers.GetLength()
		writer    interface{}
		isEnabled bool
	)

	for id := 0; id < length; id++ {
		writer, isEnabled = writers.GetValue(id)

		if !isEnabled {
			continue
		}

		writerErr = writer.(Writer).WriteLog(
			*datetime,
			*level,
			*message,
		)

		if writerErr != nil {
			Error(
				fmt.Sprintf("could not write to log writer with id=%d", id),
				writerErr,
			)
		}
	}
}
