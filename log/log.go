package log

import (
	"fmt"
	"os"
	"time"

	"github.com/mrumyantsev/logx"
	"github.com/mrumyantsev/logx/logger"
)

const (
	// insert between error description and error message
	errorInsert = ": "
)

var (
	// flags to disable writing of logs
	isDisableDebugLogs = false
	isDisableWarnLogs  = false

	// standard logger
	std = logger.New()

	// log writers linked list storage
	headWriter *writerElement
)

// ApplyConfig applies new configuration to LogX.
func ApplyConfig(cfg *logx.Config) {
	cfg.InitEmptyFields()

	isDisableDebugLogs = cfg.IsDisableDebugLogs
	isDisableWarnLogs = cfg.IsDisableWarnLogs

	// the next statements are for standard logger

	if cfg.IsDisableStandardLogger {
		std = nil

		return
	}

	if std == nil {
		std = logger.New()
	}

	if cfg.IsDisableColors {
		std.DisableColors()
	} else {
		std.EnableColors()
	}

	std.SetTimeFormat(cfg.TimeFormat)
	std.SetOutput(cfg.Output)
}

// AddWriters adds log writers to LogX. Ignores passed nil values.
func AddWriters(w ...logx.LogWriter) {
	for _, writer := range w {
		if writer == nil {
			continue
		}

		addWriter(writer)
	}
}

// RemoveWriters removes log writers from LogX. Ignores passed nil
// values.
func RemoveWriters(w ...logx.LogWriter) {
	if headWriter == nil {
		return
	}

	for _, writer := range w {
		if writer == nil {
			continue
		}

		removeWriter(writer)
	}
}

// Info writes info level log by standard LogX logger, then writes it
// by log writers, if they were added.
func Info(msg string) {
	time := time.Now()

	if std != nil {
		std.WriteLog(time, logx.InfoLevel, msg)
	}

	if headWriter != nil {
		writeByWriters(time, logx.InfoLevel, msg)
	}
}

// Debug writes debug level log by standard LogX logger, then writes it
// by log writers, if they were added. Due to configuration the
// function can be skipped.
func Debug(msg string) {
	if isDisableDebugLogs {
		return
	}

	time := time.Now()

	if std != nil {
		std.WriteLog(time, logx.DebugLevel, msg)
	}

	if headWriter != nil {
		writeByWriters(time, logx.DebugLevel, msg)
	}
}

// Warn writes warn level log by standard LogX logger, then writes it
// by log writers, if they were added. Due to configuration the
// function can be skipped.
func Warn(msg string) {
	if isDisableWarnLogs {
		return
	}

	time := time.Now()

	if std != nil {
		std.WriteLog(time, logx.WarnLevel, msg)
	}

	if headWriter != nil {
		writeByWriters(time, logx.WarnLevel, msg)
	}
}

// Error writes error level log by standard LogX logger, then writes it
// by log writers, if they were added. Value of err may be nil.
func Error(desc string, err error) {
	time := time.Now()

	if err != nil {
		desc += errorInsert + err.Error()
	}

	if std != nil {
		std.WriteLog(time, logx.ErrorLevel, desc)
	}

	if headWriter != nil {
		writeByWriters(time, logx.ErrorLevel, desc)
	}
}

// Fatal writes fatal level log by standard LogX logger, then writes it
// by log writers, if they were added, and ends the program with exit
// code 1. Value of err may be nil.
func Fatal(desc string, err error) {
	time := time.Now()

	if err != nil {
		desc += errorInsert + err.Error()
	}

	if std != nil {
		std.WriteLog(time, logx.FatalLevel, desc)
	}

	if headWriter != nil {
		writeByWriters(time, logx.FatalLevel, desc)
	}

	os.Exit(logx.FatalExitCode)
}

// FatalWithCode writes fatal level log by standard LogX logger, then
// writes it by log writers, if they were added, and ends the program
// with specified exit code. Value of err may be nil.
func FatalWithCode(desc string, err error, exitCode int) {
	time := time.Now()

	if err != nil {
		desc += errorInsert + err.Error()
	}

	if std != nil {
		std.WriteLog(time, logx.FatalLevel, desc)
	}

	if headWriter != nil {
		writeByWriters(time, logx.FatalLevel, desc)
	}

	os.Exit(exitCode)
}

// Panic writes panic level log by standard LogX logger, then writes it
// by log writers, if they were added, then calls the standard panic in
// current goroutine. Value of err may be nil.
func Panic(desc string, err error) {
	time := time.Now()

	if err != nil {
		desc += errorInsert + err.Error()
	}

	if std != nil {
		std.WriteLog(time, logx.PanicLevel, desc)
	}

	if headWriter != nil {
		writeByWriters(time, logx.PanicLevel, desc)
	}

	panic(desc)
}

// addWriter adds a single element to log writers storage.
func addWriter(w logx.LogWriter) {
	headWriter = &writerElement{
		writer: w,
		next:   headWriter,
	}
}

// removeWriter removes a single element from log writers storage.
func removeWriter(w logx.LogWriter) {
	var prev *writerElement

	for el := headWriter; el != nil; el = el.next {
		if el.writer == w {
			if el == headWriter {
				headWriter = headWriter.next
			} else {
				prev.next = el.next
			}

			return
		}

		prev = el
	}
}

// writeByWriters writes the log message by stored log writers.
func writeByWriters(time time.Time, level logx.LogLevel, msg string) {
	var err error

	for el := headWriter; el != nil; el = el.next {
		if err = el.writer.WriteLog(time, level, msg); err != nil {
			desc := fmt.Sprintf(
				"could not write to log writer=%T",
				el.writer,
			) + errorInsert + err.Error()

			_ = writeToStream(time, logx.ErrorLevel, desc)
		}
	}
}

// writeToStream writes the log message to default output data stream.
func writeToStream(time time.Time, level logx.LogLevel, msg string) error {
	var err error

	_, err = logx.Output().Write(
		[]byte(
			time.Format(logx.TimeFormat) +
				logx.Space +
				logx.LevelText(level) +
				logx.Space +
				msg +
				logx.EndOfLine,
		),
	)

	return err
}

// writerElement is an element of log writers linked list storage.
type writerElement struct {
	writer logx.LogWriter
	next   *writerElement
}
