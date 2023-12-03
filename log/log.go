package log

import (
	"fmt"
	"os"
	"time"

	"github.com/mrumyantsev/go-idmap"
	"github.com/mrumyantsev/multilog"
)

// Logger working constants.
const (
	// inserts between the error and
	// its description. Used in error
	// and fatal level logging functions
	_ERROR_INSERT string = ". error: "
)

// Logger working variables.
var (
	// defines initial configuration,
	// when the application starts
	config *multilog.Config = multilog.NewConfig()
	// contains the log writer objects,
	// that should be implemented by user
	writers *idmap.IdMap = nil
	// to store the error, occurred by
	// user's log writer
	writerErr error = nil
)

// Apply the configuration, that was created by user, to logger.
func ApplyConfig(cfg *multilog.Config) {
	cfg.InitEmptyFields()
	config = cfg
}

// A log writer's interface, that need to implement by user's object,
// that able to write logs.
type Writer interface {
	WriteLog(datetime time.Time, level string, message string) error
}

// Add implemented log writer object with its ID.
func AddWriter(id int, writer Writer) {
	if writers == nil {
		writers = idmap.New()
	}

	writers.SetValue(id, writer)
}

// Remove implemented log writer object by its ID.
func RemoveWriter(id int) {
	writers.DeleteValue(id)
}

// Enable implemented log writer object by its ID.
func EnableWriter(id int) {
	writers.Enable(id)
}

// Disable implemented log writer object by its ID.
func DisableWriter(id int) {
	writers.Disable(id)
}

// Write info level log to its own output stream. Then write it to the
// log writers (that exists and set to enabled).
func Info(msg string) {
	var datetime time.Time = time.Now()

	writeToStream(
		&datetime,
		&config.InfoLevelText,
		&config.InfoLevelColor,
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

// Write debug level log to its own output stream. Then write it to the
// log writers (that exists and set to enabled).
func Debug(msg string) {
	if config.IsDisableDebugLogs {
		return
	}

	var datetime time.Time = time.Now()

	writeToStream(
		&datetime,
		&config.DebugLevelText,
		&config.DebugLevelColor,
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

// Write warn level log to its own output stream. Then write it to the
// log writers (that exists and set to enabled).
func Warn(msg string) {
	if config.IsDisableWarnLogs {
		return
	}

	var datetime time.Time = time.Now()

	writeToStream(
		&datetime,
		&config.WarnLevelText,
		&config.WarnLevelColor,
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

// Write error level log to its own output stream. Then write it to the
// log writers (that exists and set to enabled).
func Error(desc string, err error) {
	var datetime time.Time = time.Now()

	desc = desc + _ERROR_INSERT + err.Error()

	writeToStream(
		&datetime,
		&config.ErrorLevelText,
		&config.ErrorLevelColor,
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

// Write fatal level log to its own output stream. Then write it to the
// log writers (that exists and set to enabled). Exits the program at
// the end with the exit code 1.
func Fatal(desc string, err error) {
	var datetime time.Time = time.Now()

	desc = desc + _ERROR_INSERT + err.Error()

	writeToStream(
		&datetime,
		&config.FatalLevelText,
		&config.FatalLevelColor,
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

	os.Exit(1)
}

// Write fatal level log to its own output stream. Then write it to the
// log writers (that exists and set to enabled). Exits the program at
// the end with the exit code that set in argument.
func FatalWithCode(desc string, err error, exitCode int) {
	var datetime time.Time = time.Now()

	desc = desc + _ERROR_INSERT + err.Error()

	writeToStream(
		&datetime,
		&config.FatalLevelText,
		&config.FatalLevelColor,
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

	os.Exit(exitCode)
}

// Write to output stream function. The only function, that does not
// handle the error (if it will occur).
func writeToStream(
	datetime *time.Time,
	level *string,
	levelColor *string,
	message *string,
	stream *os.File,
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

// Write to log writers function. Does loop calling WriteLog() method
// for every log writer object, that stored in the logger, and also
// that has isEnabled=true flag.
func writeToWriters(
	datetime *time.Time,
	level *string,
	message *string,
) {
	var (
		length    int         = writers.GetLength()
		writer    interface{} = nil
		id        int         = 0
		isEnabled bool        = false
	)

	for ; id < length; id++ {
		// gets log writer by its ID
		writer, isEnabled = writers.GetValue(id)

		// checks, if it is enabled
		if !isEnabled {
			continue
		}

		// writes current log to enabled log writer
		writerErr = writer.(Writer).WriteLog(
			*datetime,
			*level,
			*message,
		)
		if writerErr != nil {
			// writes error level log to error output
			// stream, when the error occurs
			var desc string = fmt.Sprintf(
				"could not write to log writer with id=%d", id) +
				_ERROR_INSERT + writerErr.Error()

			writeToStream(
				datetime,
				&config.ErrorLevelText,
				&config.ErrorLevelColor,
				&desc,
				config.ErrorOutputStream,
			)
		}
	}
}
