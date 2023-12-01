package log

import (
	"fmt"
	"os"
	"time"

	"github.com/mrumyantsev/go-idmap"
	"github.com/mrumyantsev/multilog"
	"github.com/mrumyantsev/multilog/defaults"
)

const (
	_ERROR_WORD string = ". error: "
)

var (
	config *multilog.Config = &multilog.Config{
		InfoOutputStream:  os.Stderr,
		DebugOutputStream: os.Stderr,
		WarnOutputStream:  os.Stderr,
		ErrorOutputStream: os.Stderr,
		FatalOutputStream: os.Stderr,

		IsEnableDebugLogs: true,
		IsEnableWarnLogs:  true,

		FatalExitStatusCode: 1,

		TimeFormat: defaults.TIME_FORMAT,

		ItemSeparatorText: defaults.ITEM_SEPARATOR_TEXT,
		LineEndingText:    defaults.LINE_ENDING_TEXT,
		InfoLevelText:     defaults.INFO_LEVEL_TEXT,
		DebugLevelText:    defaults.DEBUG_LEVEL_TEXT,
		WarnLevelText:     defaults.WARN_LEVEL_TEXT,
		ErrorLevelText:    defaults.ERROR_LEVEL_TEXT,
		FatalLevelText:    defaults.FATAL_LEVEL_TEXT,
	}
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

func ApplyConfig(cfg *multilog.Config) {
	if cfg.InfoOutputStream == nil {
		cfg.InfoOutputStream = os.Stderr
	}

	if cfg.DebugOutputStream == nil {
		cfg.DebugOutputStream = os.Stderr
	}

	if cfg.WarnOutputStream == nil {
		cfg.WarnOutputStream = os.Stderr
	}

	if cfg.ErrorOutputStream == nil {
		cfg.ErrorOutputStream = os.Stderr
	}

	if cfg.FatalOutputStream == nil {
		cfg.FatalOutputStream = os.Stderr
	}

	config = cfg
}

func Info(msg string) {
	datetime := time.Now()

	writeToStream(
		&datetime,
		&config.InfoLevelText,
		&msg,
		config.InfoOutputStream,
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
	if !config.IsEnableDebugLogs {
		return
	}

	datetime := time.Now()

	writeToStream(
		&datetime,
		&config.DebugLevelText,
		&msg,
		config.DebugOutputStream,
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
	if !config.IsEnableWarnLogs {
		return
	}

	datetime := time.Now()

	writeToStream(
		&datetime,
		&config.WarnLevelText,
		&msg,
		config.WarnOutputStream,
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

	desc = desc + _ERROR_WORD + err.Error()

	writeToStream(
		&datetime,
		&config.ErrorLevelText,
		&desc,
		config.ErrorOutputStream,
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

	desc = desc + _ERROR_WORD + err.Error()

	writeToStream(
		&datetime,
		&config.FatalLevelText,
		&desc,
		config.FatalOutputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&config.FatalLevelText,
			&desc,
		)
	}

	os.Exit(config.FatalExitStatusCode)
}

func writeToStream(
	datetime *time.Time,
	level *string,
	message *string,
	stream *os.File,
) {
	stream.Write(
		[]byte(
			(*datetime).Format(config.TimeFormat) +
				config.ItemSeparatorText +
				*level +
				config.ItemSeparatorText +
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
