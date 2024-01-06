package log

import (
	"container/list"
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
	errorInsert string = ". error: "
)

var (
	// flags to disable logs execution
	isDisableDebugLogs bool = false
	isDisableWarnLogs  bool = false

	// level text for the stream logs
	infoLevel  string = defaults.InfoLevel
	debugLevel string = defaults.DebugLevel
	warnLevel  string = defaults.WarnLevel
	errorLevel string = defaults.ErrorLevel
	fatalLevel string = defaults.FatalLevel
	panicLevel string = defaults.PanicLevel

	// colors of level text for the stream logs
	regularColor string = defaults.RegularColor
	grayColor    string = defaults.GrayColor
	greenColor   string = defaults.GreenColor
	yellowColor  string = defaults.YellowColor
	redColor     string = defaults.RedColor
	boldRedColor string = defaults.BoldRedColor

	// datetime patten for the stream logs
	timeFormat string = defaults.TimeFormat

	// data stream pointer for the stream logs
	outputStream *os.File = defaults.GetOutputStream()

	// log writer interface objects
	writers *list.List = list.New()
)

// Apply new configuration to the logger.
func ApplyConfig(cfg *multilog.Config) {
	cfg.InitEmptyFields()

	isDisableDebugLogs = cfg.IsDisableDebugLogs
	isDisableWarnLogs = cfg.IsDisableWarnLogs

	if cfg.IsDisableColors {
		regularColor = defaults.EmptyString
		grayColor = defaults.EmptyString
		greenColor = defaults.EmptyString
		yellowColor = defaults.EmptyString
		redColor = defaults.EmptyString
		boldRedColor = defaults.EmptyString
	} else {
		regularColor = defaults.RegularColor
		grayColor = defaults.GrayColor
		greenColor = defaults.GreenColor
		yellowColor = defaults.YellowColor
		redColor = defaults.RedColor
		boldRedColor = defaults.BoldRedColor
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
	writers.PushBack(w)
}

// Remove log writer, so the logger can not call it any more to do logs.
func RemoveWriter(w Writer) {
	for e := writers.Front(); e != nil; e = e.Next() {
		if e.Value.(Writer) == w {
			writers.Remove(e)
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
		&greenColor,
		&msg,
		outputStream,
	)

	if writers.Len() > 0 {
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
		&yellowColor,
		&msg,
		outputStream,
	)

	if writers.Len() > 0 {
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
		&redColor,
		&msg,
		outputStream,
	)

	if writers.Len() > 0 {
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

	desc = desc + errorInsert + err.Error()

	writeToStream(
		&datetime,
		&errorLevel,
		&boldRedColor,
		&desc,
		outputStream,
	)

	if writers.Len() > 0 {
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

	desc = desc + errorInsert + err.Error()

	writeToStream(
		&datetime,
		&fatalLevel,
		&boldRedColor,
		&desc,
		outputStream,
	)

	if writers.Len() > 0 {
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

	desc = desc + errorInsert + err.Error()

	writeToStream(
		&datetime,
		&fatalLevel,
		&boldRedColor,
		&desc,
		outputStream,
	)

	if writers.Len() > 0 {
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

	desc = desc + errorInsert + err.Error()

	writeToStream(
		&datetime,
		&panicLevel,
		&boldRedColor,
		&desc,
		outputStream,
	)

	if writers.Len() > 0 {
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
			grayColor +
				(*datetime).Format(timeFormat) +
				defaults.Space +
				*levelColor +
				*level +
				defaults.Space +
				regularColor +
				*message +
				defaults.EndOfLine,
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
		writer Writer = nil
		err    error  = nil
	)

	for e := writers.Front(); e != nil; e = e.Next() {
		writer = e.Value.(Writer)

		// write current log by the log writer
		err = writer.WriteLog(
			*datetime,
			*level,
			*message,
		)
		if err != nil {
			// write the error level log to the
			// stream, if the error occurs
			var desc string = fmt.Sprintf(
				"could not write to log writer=%T", writer) +
				errorInsert + err.Error()

			writeToStream(
				datetime,
				&errorLevel,
				&boldRedColor,
				&desc,
				outputStream,
			)
		}
	}
}
