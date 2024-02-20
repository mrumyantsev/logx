package logger

import (
	"os"
	"time"

	"github.com/mrumyantsev/logx"
)

// Standard LogX logger. Implements logx.LogWriter interface and writes
// logs to data stream file (stderr by default).
type Logger struct {
	isDisableColors bool
	timeFormat      string
	output          *os.File
}

func New() *Logger {
	return &Logger{
		timeFormat: logx.TimeFormat,
		output:     logx.Output(),
	}
}

func (l *Logger) SetTimeFormat(timeFormat string) {
	l.timeFormat = timeFormat
}

func (l *Logger) Output() *os.File {
	return l.output
}

func (l *Logger) SetOutput(f *os.File) {
	l.output = f
}

func (l *Logger) EnableColors() {
	l.isDisableColors = false
}

func (l *Logger) DisableColors() {
	l.isDisableColors = true
}

// WriteLog writes the log message to output data stream.
func (l *Logger) WriteLog(time time.Time, level logx.LogLevel, msg string) error {
	var err error

	if l.isDisableColors {
		_, err = l.output.Write(
			[]byte(
				time.Format(l.timeFormat) +
					logx.Space +
					logx.LevelText(level) +
					logx.Space +
					msg +
					logx.EndOfLine,
			),
		)

		return err
	}

	_, err = l.output.Write(
		[]byte(
			logx.GrayColor +
				time.Format(l.timeFormat) +
				logx.Space +
				logx.LevelColor(level) +
				logx.LevelText(level) +
				logx.Space +
				logx.RegularColor +
				msg +
				logx.EndOfLine,
		),
	)

	return err
}
