package log

import (
	"fmt"
	"os"
	"time"

	"github.com/mrumyantsev/multilog"
	"github.com/mrumyantsev/multilog/defaults"
)

// Logger working constants.
const (
	// insert between the error and
	// its description. Used in error,
	// fatal and panic level logging
	// functions
	_ERROR_INSERT string = ". error: "
)

// Logger working variables.
var (
	// flags to disable logs execution
	isDisableDebugLogs bool = false
	isDisableWarnLogs  bool = false

	// utility text parts for stream logs
	itemSeparatorText string = defaults.ITEM_SEPARATOR_TEXT
	lineEndingText    string = defaults.LINE_ENDING_TEXT

	// level text for stream logs
	infoLevelText  string = defaults.INFO_LEVEL_TEXT
	debugLevelText string = defaults.DEBUG_LEVEL_TEXT
	warnLevelText  string = defaults.WARN_LEVEL_TEXT
	errorLevelText string = defaults.ERROR_LEVEL_TEXT
	fatalLevelText string = defaults.FATAL_LEVEL_TEXT
	panicLevelText string = defaults.PANIC_LEVEL_TEXT

	// colors of level text for stream logs
	datetimeColor   string = defaults.DATETIME_COLOR
	infoLevelColor  string = defaults.INFO_LEVEL_COLOR
	debugLevelColor string = defaults.DEBUG_LEVEL_COLOR
	warnLevelColor  string = defaults.WARN_LEVEL_COLOR
	errorLevelColor string = defaults.ERROR_LEVEL_COLOR
	fatalLevelColor string = defaults.FATAL_LEVEL_COLOR
	panicLevelColor string = defaults.PANIC_LEVEL_COLOR
	messageColor    string = defaults.MESSAGE_COLOR

	// datetime patten for stream logs
	timeFormat string = defaults.TIME_FORMAT

	// data stream pointer for stream logs
	outputStream *os.File = defaults.GetOutputStream()

	// log writer interface objects
	writers []Writer = nil

	// log writers total count
	writersCount int = 0

	// log writer error
	writerErr error = nil
)

// Apply the configuration, that was created by user, to logger.
func ApplyConfig(cfg *multilog.Config) {
	cfg.InitEmptyFields()

	if cfg.IsDisableColors {
		datetimeColor = defaults.EMPTY_STRING
		infoLevelColor = defaults.EMPTY_STRING
		debugLevelColor = defaults.EMPTY_STRING
		warnLevelColor = defaults.EMPTY_STRING
		errorLevelColor = defaults.EMPTY_STRING
		fatalLevelColor = defaults.EMPTY_STRING
		panicLevelColor = defaults.EMPTY_STRING
		messageColor = defaults.EMPTY_STRING
	} else {
		datetimeColor = defaults.DATETIME_COLOR
		infoLevelColor = defaults.INFO_LEVEL_COLOR
		debugLevelColor = defaults.DEBUG_LEVEL_COLOR
		warnLevelColor = defaults.WARN_LEVEL_COLOR
		errorLevelColor = defaults.ERROR_LEVEL_COLOR
		fatalLevelColor = defaults.FATAL_LEVEL_COLOR
		panicLevelColor = defaults.PANIC_LEVEL_COLOR
		messageColor = defaults.MESSAGE_COLOR
	}

	timeFormat = cfg.TimeFormat
	outputStream = cfg.OutputStream
}

// A log writer's interface, that need to implement by user's object,
// that able to write logs.
type Writer interface {
	WriteLog(datetime time.Time, level string, message string) error
}

// Add implemented log writer object.
func AddWriter(writer Writer) {
	if writers == nil {
		writers = []Writer{writer}
		writersCount = 1
		return
	}

	var i int = 0

	for ; i < writersCount; i++ {
		if writers[i] == writer {
			return
		}
	}

	i = 0

	for ; i < writersCount; i++ {
		if writers[i] == nil {
			writers[i] = writer
			return
		}
	}

	writers = append(writers, writer)
	writersCount++
}

// Remove implemented log writer object.
func RemoveWriter(writer Writer) {
	if writers == nil {
		return
	}

	var i int = 0

	for ; i < writersCount; i++ {
		if writers[i] == writer {
			writers[i] = nil
			return
		}
	}
}

// Write info level log to its own output stream. Then write it to the
// log writers (that exists and set to enabled).
func Info(msg string) {
	var datetime time.Time = time.Now()

	writeToStream(
		&datetime,
		&infoLevelText,
		&infoLevelColor,
		&msg,
		outputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&infoLevelText,
			&msg,
		)
	}
}

// Write debug level log to its own output stream. Then write it to the
// log writers (that exists and set to enabled).
func Debug(msg string) {
	if isDisableDebugLogs {
		return
	}

	var datetime time.Time = time.Now()

	writeToStream(
		&datetime,
		&debugLevelText,
		&debugLevelColor,
		&msg,
		outputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&debugLevelText,
			&msg,
		)
	}
}

// Write warn level log to its own output stream. Then write it to the
// log writers (that exists and set to enabled).
func Warn(msg string) {
	if isDisableWarnLogs {
		return
	}

	var datetime time.Time = time.Now()

	writeToStream(
		&datetime,
		&warnLevelText,
		&warnLevelColor,
		&msg,
		outputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&warnLevelText,
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
		&errorLevelText,
		&errorLevelColor,
		&desc,
		outputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&errorLevelText,
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
		&fatalLevelText,
		&fatalLevelColor,
		&desc,
		outputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&fatalLevelText,
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
		&fatalLevelText,
		&fatalLevelColor,
		&desc,
		outputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&fatalLevelText,
			&desc,
		)
	}

	os.Exit(exitCode)
}

// Write panic level log to its own output stream. Then write it to the
// log writers (that exists and set to enabled). Then call standard
// panic in the current goroutine.
func Panic(desc string, err error) {
	var datetime time.Time = time.Now()

	desc = desc + _ERROR_INSERT + err.Error()

	writeToStream(
		&datetime,
		&panicLevelText,
		&panicLevelColor,
		&desc,
		outputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&panicLevelText,
			&desc,
		)
	}

	panic(desc)
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
			datetimeColor +
				(*datetime).Format(timeFormat) +
				itemSeparatorText +
				*levelColor +
				*level +
				itemSeparatorText +
				messageColor +
				*message +
				lineEndingText,
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
		i      int    = 0
		writer Writer = nil
	)

	for ; i < writersCount; i++ {
		writer = writers[i]

		if writer == nil {
			continue
		}

		// writes current log to enabled log writer
		writerErr = writer.WriteLog(
			*datetime,
			*level,
			*message,
		)
		if writerErr != nil {
			// writes error level log to error output
			// stream, when the error occurs
			var desc string = fmt.Sprintf(
				"could not write to log writer=%T", writer) +
				_ERROR_INSERT + writerErr.Error()

			writeToStream(
				datetime,
				&errorLevelText,
				&errorLevelColor,
				&desc,
				outputStream,
			)
		}
	}
}
