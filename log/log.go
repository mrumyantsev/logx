package log

import (
	"container/list"
	"fmt"
	"os"
	"time"

	"github.com/mrumyantsev/logx"
	"github.com/mrumyantsev/logx/logger"
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

	// log writer interface objects
	writers *list.List = initWriters()
)

// initWriters returns list of writers with standard logger in it.
func initWriters() *list.List {
	list := list.New()

	stdLogger := logger.New()

	list.PushFront(stdLogger)

	return list
}

// Apply new configuration to the logger.
func ApplyConfig(cfg *logx.Config) {
	cfg.InitEmptyFields()

	isDisableDebugLogs = cfg.IsDisableDebugLogs
	isDisableWarnLogs = cfg.IsDisableWarnLogs

	// Next: apply parameters for standard logger

	// get head element (if not nil)
	headWriter := writers.Front()
	if headWriter == nil {
		return
	}

	// cast to standard logger
	logger, ok := headWriter.Value.(*logger.Logger)
	if !ok {
		return
	}

	// remove standard logger and exit (if config flag is set)
	if cfg.IsDisableDefaultConsoleLogger {
		writers.Remove(headWriter)
		return
	}

	// apply new configuration to standard logger
	logger.SetDisableColors(cfg.IsDisableColors)
	logger.SetTimeFormat(cfg.TimeFormat)
	logger.SetOutputStream(cfg.OutputStream)
}

// Log writer interface. Any implemented objects are assumed to be
// supplemental log writers to the logger. Implement this interface
// with your custom writer and add it to the logger by calling the
// AddWriter() method, and the logger will send logs through it.
type Writer interface {
	WriteLog(datetime time.Time, levelId uint8, message string) error
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

	if writers.Len() > 0 {
		writeToWriters(
			&datetime,
			logx.InfoLevelId,
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

	if writers.Len() > 0 {
		writeToWriters(
			&datetime,
			logx.DebugLevelId,
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

	if writers.Len() > 0 {
		writeToWriters(
			&datetime,
			logx.WarnLevelId,
			&msg,
		)
	}
}

// Write the error level log to stream, then write it by all of the log
// writers (if they were added).
func Error(desc string, err error) {
	var datetime time.Time = time.Now()

	desc = desc + errorInsert + err.Error()

	if writers.Len() > 0 {
		writeToWriters(
			&datetime,
			logx.ErrorLevelId,
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

	if writers.Len() > 0 {
		writeToWriters(
			&datetime,
			logx.FatalLevelId,
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

	if writers.Len() > 0 {
		writeToWriters(
			&datetime,
			logx.FatalLevelId,
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

	if writers.Len() > 0 {
		writeToWriters(
			&datetime,
			logx.PanicLevelId,
			&desc,
		)
	}

	panic(desc)
}

// Write the logs by the log writers, that been added to the logger.
func writeToWriters(
	datetime *time.Time,
	levelId uint8,
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
			levelId,
			*message,
		)
		if err != nil {
			// write the error level log to the
			// stream, if the error occurs
			var (
				desc string = fmt.Sprintf(
					"could not write to log writer=%T", writer) +
					errorInsert + err.Error()
				stream *os.File = logx.GetOutputStream()
			)

			stream.Write(
				[]byte(
					datetime.Format(logx.TimeFormat) +
						logx.Space +
						logx.GetLevelText(levelId) +
						logx.Space +
						desc +
						logx.EndOfLine,
				),
			)
		}
	}
}
