package log

import (
	"fmt"
	"os"
	"time"

	"github.com/mrumyantsev/logx"
	"github.com/mrumyantsev/logx/logger"
)

// New creates a new [Logger]. The out variable sets the destination to
// which log data will be written. The prefix appears at the beginning
// of each generated log line, or after the log header if the
// [Lmsgprefix] flag is provided. The flag argument defines the logging
// properties.
//
// prefix and flag parameters are not implemented.
func New(out *os.File, prefix string, flag int) *logger.Logger {
	logger := logger.New()

	logger.SetOutput(out)

	return logger
}

// Default returns the standard logger used by the package-level output
// functions.
func Default() *logger.Logger {
	return std
}

// SetOutput sets the output destination for the standard logger.
func SetOutput(f *os.File) {
	if std == nil {
		return
	}

	std.SetOutput(f)
}

// Writer returns the output destination for the standard logger.
func Writer() *os.File {
	if std == nil {
		return nil
	}

	return std.Output()
}

// Print calls Output to print to the standard logger. Arguments are
// handled in the manner of fmt.Print.
//
// Alias for log.Info()
func Print(v ...interface{}) {
	time := time.Now()
	msg := fmt.Sprint(v...)

	if std != nil {
		std.WriteLog(time, logx.InfoLevel, msg)
	}

	if headWriter != nil {
		writeByWriters(time, logx.InfoLevel, msg)
	}
}

// Printf calls Output to print to the standard logger. Arguments are
// handled in the manner of fmt.Printf.
//
// Alias for log.Info()
func Printf(format string, v ...interface{}) {
	time := time.Now()
	msg := fmt.Sprintf(format, v...)

	if std != nil {
		std.WriteLog(time, logx.InfoLevel, msg)
	}

	if headWriter != nil {
		writeByWriters(time, logx.InfoLevel, msg)
	}
}

// Println calls Output to print to the standard logger. Arguments are
// handled in the manner of fmt.Println.
//
// Alias for log.Info()
func Println(v ...interface{}) {
	time := time.Now()
	msg := fmt.Sprintln(v...)
	msg = msg[:len(msg)-1]

	if std != nil {
		std.WriteLog(time, logx.InfoLevel, msg)
	}

	if headWriter != nil {
		writeByWriters(time, logx.InfoLevel, msg)
	}
}

// Fatalf is equivalent to [Printf] followed by a call to os.Exit(logx.FatalExitCode).
//
// Alias for log.Fatal()
func Fatalf(format string, v ...interface{}) {
	time := time.Now()
	msg := fmt.Sprintf(format, v...)

	if std != nil {
		std.WriteLog(time, logx.FatalLevel, msg)
	}

	if headWriter != nil {
		writeByWriters(time, logx.FatalLevel, msg)
	}
}

// Fatalln is equivalent to [Println] followed by a call to os.Exit(logx.FatalExitCode).
//
// Alias for log.Fatal()
func Fatalln(v ...interface{}) {
	time := time.Now()
	msg := fmt.Sprintln(v...)
	msg = msg[:len(msg)-1]

	if std != nil {
		std.WriteLog(time, logx.FatalLevel, msg)
	}

	if headWriter != nil {
		writeByWriters(time, logx.FatalLevel, msg)
	}
}

// Panicf is equivalent to [Printf] followed by a call to panic().
//
// Alias for log.Panic()
func Panicf(format string, v ...interface{}) {
	time := time.Now()
	msg := fmt.Sprintf(format, v...)

	if std != nil {
		std.WriteLog(time, logx.PanicLevel, msg)
	}

	if headWriter != nil {
		writeByWriters(time, logx.PanicLevel, msg)
	}
}

// Panicln is equivalent to [Println] followed by a call to panic().
//
// Alias for log.Panic()
func Panicln(v ...interface{}) {
	time := time.Now()
	msg := fmt.Sprintln(v...)
	msg = msg[:len(msg)-1]

	if std != nil {
		std.WriteLog(time, logx.PanicLevel, msg)
	}

	if headWriter != nil {
		writeByWriters(time, logx.PanicLevel, msg)
	}
}
