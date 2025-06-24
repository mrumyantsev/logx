package logger

import (
	"os"
	"strings"
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
	var sb strings.Builder
	var err error

	if l.isDisableColors {
		sb.WriteString(time.Format(l.timeFormat))
		sb.WriteString(logx.Space)
		sb.WriteString(logx.LevelText(level))
		sb.WriteString(logx.Space)
		sb.WriteString(msg)
		sb.WriteString(logx.EndOfLine)

		_, err = l.output.WriteString(sb.String())

		return err
	}

	sb.WriteString(logx.GrayColor)
	sb.WriteString(time.Format(l.timeFormat))
	sb.WriteString(logx.Space)
	sb.WriteString(logx.LevelColor(level))
	sb.WriteString(logx.LevelText(level))
	sb.WriteString(logx.Space)
	sb.WriteString(logx.RegularColor)
	sb.WriteString(msg)
	sb.WriteString(logx.EndOfLine)

	_, err = l.output.WriteString(sb.String())

	return err
}
