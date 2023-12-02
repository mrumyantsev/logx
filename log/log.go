package log

import (
	"fmt"
	"os"
	"time"

	"github.com/mrumyantsev/go-idmap"
	"github.com/mrumyantsev/multilog"
)

const (
	_ERROR_INSERT string = ". error: "
)

var (
	config    *multilog.Config = multilog.NewConfig()
	writers   *idmap.IdMap     = nil
	writerErr error            = nil
)

func ApplyConfig(cfg *multilog.Config) {
	cfg.InitEmptyFields()
	config = cfg
}

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
		&config.InfoLevelText,
		&msg,
		config.InfoOutputStream,
		&config.InfoLevelColor,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&config.InfoLevelText,
			&msg,
		)
	}
}

func Debug(msg string) {
	if config.IsDisableDebugLogs {
		return
	}

	datetime := time.Now()

	writeToStream(
		&datetime,
		&config.DebugLevelText,
		&msg,
		config.DebugOutputStream,
		&config.DebugLevelColor,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&config.DebugLevelText,
			&msg,
		)
	}
}

func Warn(msg string) {
	if config.IsDisableWarnLogs {
		return
	}

	datetime := time.Now()

	writeToStream(
		&datetime,
		&config.WarnLevelText,
		&msg,
		config.WarnOutputStream,
		&config.WarnLevelColor,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&config.WarnLevelText,
			&msg,
		)
	}
}

func Error(desc string, err error) {
	datetime := time.Now()

	desc = desc + _ERROR_INSERT + err.Error()

	writeToStream(
		&datetime,
		&config.ErrorLevelText,
		&desc,
		config.ErrorOutputStream,
		&config.ErrorLevelColor,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&config.ErrorLevelText,
			&desc,
		)
	}
}

func Fatal(desc string, err error) {
	datetime := time.Now()

	desc = desc + _ERROR_INSERT + err.Error()

	writeToStream(
		&datetime,
		&config.FatalLevelText,
		&desc,
		config.FatalOutputStream,
		&config.FatalLevelColor,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&config.FatalLevelText,
			&desc,
		)
	}

	os.Exit(1)
}

func FatalWithCode(desc string, err error, exitCode int) {
	datetime := time.Now()

	desc = desc + _ERROR_INSERT + err.Error()

	writeToStream(
		&datetime,
		&config.FatalLevelText,
		&desc,
		config.FatalOutputStream,
		&config.FatalLevelColor,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&config.FatalLevelText,
			&desc,
		)
	}

	os.Exit(exitCode)
}

func writeToStream(
	datetime *time.Time,
	level *string,
	message *string,
	stream *os.File,
	levelColor *string,
) {
	stream.Write(
		[]byte(
			config.DatetimeColor +
				(*datetime).Format(config.TimeFormat) +
				config.ItemSeparatorText +
				*levelColor +
				*level +
				config.ItemSeparatorText +
				config.MessageColor +
				*message +
				config.LineEndingText,
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
