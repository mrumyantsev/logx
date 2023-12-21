package log

import (
	"fmt"
	"os"
	"time"

	"github.com/mrumyantsev/multilog"
	"github.com/mrumyantsev/multilog/defaults"
)

const (
	// insert between the error and
	// its description. Used in error,
	// fatal and panic level logging
	// functions
	_ERROR_INSERT string = ". error: "
)

var (
	// flags to disable logs execution
	isDisableDebugLogs bool = false
	isDisableWarnLogs  bool = false

	// utility text parts for the stream logs
	itemSeparator string = defaults.ITEM_SEPARATOR
	lineEnding    string = defaults.LINE_ENDING

	// level text for the stream logs
	infoLevel  string = defaults.INFO_LEVEL
	debugLevel string = defaults.DEBUG_LEVEL
	warnLevel  string = defaults.WARN_LEVEL
	errorLevel string = defaults.ERROR_LEVEL
	fatalLevel string = defaults.FATAL_LEVEL
	panicLevel string = defaults.PANIC_LEVEL

	// colors of level text for the stream logs
	datetimeColor   string = defaults.DATETIME_COLOR
	infoLevelColor  string = defaults.INFO_LEVEL_COLOR
	debugLevelColor string = defaults.DEBUG_LEVEL_COLOR
	warnLevelColor  string = defaults.WARN_LEVEL_COLOR
	errorLevelColor string = defaults.ERROR_LEVEL_COLOR
	fatalLevelColor string = defaults.FATAL_LEVEL_COLOR
	panicLevelColor string = defaults.PANIC_LEVEL_COLOR
	messageColor    string = defaults.MESSAGE_COLOR

	// datetime patten for the stream logs
	timeFormat string = defaults.TIME_FORMAT

	// data stream pointer for the stream logs
	outputStream *os.File = defaults.GetOutputStream()

	// log writer interface objects
	writers []Writer = nil

	// log writers slice length
	writersLength int = 0

	// log writer error
	writerErr error = nil
)

// Apply new configuration to the logger.
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

// Log writer interface. Any implemented objects are assumed to be
// supplemental log writers to the logger. Implement this interface
// with your custom writer and add it to logger by calling the
// AddWriter() method, and the logger will send logs to it.
type Writer interface {
	WriteLog(datetime time.Time, level string, message string) error
}

// Add log writer, so the logger can call it to do logs with it.
func AddWriter(w Writer) {
	if writers == nil {
		writers = []Writer{w}
		writersLength = 1
		return
	}

	var i int = 0

	for ; i < writersLength; i++ {
		if writers[i] == w {
			return
		}
	}

	i = 0

	for ; i < writersLength; i++ {
		if writers[i] == nil {
			writers[i] = w
			return
		}
	}

	writers = append(writers, w)
	writersLength++
}

// Remove log writer, so the logger can not call it any more to do logs.
func RemoveWriter(w Writer) {
	if writers == nil {
		return
	}

	var i int = 0

	for ; i < writersLength; i++ {
		if writers[i] == w {
			writers[i] = nil
			return
		}
	}
}

// Write the info level log to stream, then write it by all of the log
// writers (if they were added).
func Info(msg string) {
	var datetime time.Time = time.Now()

	writeToStream(
		&datetime,
		&infoLevel,
		&infoLevelColor,
		&msg,
		outputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&infoLevel,
			&msg,
		)
	}
}

// Write the debug level log to stream, then write it by all of the log
// writers (if they were added).
func Debug(msg string) {
	if isDisableDebugLogs {
		return
	}

	var datetime time.Time = time.Now()

	writeToStream(
		&datetime,
		&debugLevel,
		&debugLevelColor,
		&msg,
		outputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&debugLevel,
			&msg,
		)
	}
}

// Write the warning level log to stream, then write it by all of the log
// writers (if they were added).
func Warn(msg string) {
	if isDisableWarnLogs {
		return
	}

	var datetime time.Time = time.Now()

	writeToStream(
		&datetime,
		&warnLevel,
		&warnLevelColor,
		&msg,
		outputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&warnLevel,
			&msg,
		)
	}
}

// Write the error level log to stream, then write it by all of the log
// writers (if they were added).
func Error(desc string, err error) {
	var datetime time.Time = time.Now()

	desc = desc + _ERROR_INSERT + err.Error()

	writeToStream(
		&datetime,
		&errorLevel,
		&errorLevelColor,
		&desc,
		outputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&errorLevel,
			&desc,
		)
	}
}

// Write the fatal level log to stream, then write it by all of the log
// writers (if they were added). Then exit the program at the end with
// the exit code 1.
func Fatal(desc string, err error) {
	var datetime time.Time = time.Now()

	desc = desc + _ERROR_INSERT + err.Error()

	writeToStream(
		&datetime,
		&fatalLevel,
		&fatalLevelColor,
		&desc,
		outputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&fatalLevel,
			&desc,
		)
	}

	os.Exit(1)
}

// Write the fatal level log to stream, then write it by all of the log
// writers (if they were added). Then exit the program at the end with
// the exit code that set by the exitCode argument.
func FatalWithCode(desc string, err error, exitCode int) {
	var datetime time.Time = time.Now()

	desc = desc + _ERROR_INSERT + err.Error()

	writeToStream(
		&datetime,
		&fatalLevel,
		&fatalLevelColor,
		&desc,
		outputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&fatalLevel,
			&desc,
		)
	}

	os.Exit(exitCode)
}

// Write the panic level log to stream, then write it by all of the log
// writers (if they were added). Then call a standard panic in the
// current goroutine.
func Panic(desc string, err error) {
	var datetime time.Time = time.Now()

	desc = desc + _ERROR_INSERT + err.Error()

	writeToStream(
		&datetime,
		&panicLevel,
		&panicLevelColor,
		&desc,
		outputStream,
	)

	if writers != nil {
		writeToWriters(
			&datetime,
			&panicLevel,
			&desc,
		)
	}

	panic(desc)
}

// Write the log to the stream.
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
				itemSeparator +
				*levelColor +
				*level +
				itemSeparator +
				messageColor +
				*message +
				lineEnding,
		),
	)
}

// Write the logs by the log writers, that been added to the logger.
func writeToWriters(
	datetime *time.Time,
	level *string,
	message *string,
) {
	var (
		i      int    = 0
		writer Writer = nil
	)

	for ; i < writersLength; i++ {
		writer = writers[i]

		if writer == nil {
			continue
		}

		// write current log by the log writer
		writerErr = writer.WriteLog(
			*datetime,
			*level,
			*message,
		)
		if writerErr != nil {
			// write the error level log to the
			// stream, if the error occurs
			var desc string = fmt.Sprintf(
				"could not write to log writer=%T", writer) +
				_ERROR_INSERT + writerErr.Error()

			writeToStream(
				datetime,
				&errorLevel,
				&errorLevelColor,
				&desc,
				outputStream,
			)
		}
	}
}
