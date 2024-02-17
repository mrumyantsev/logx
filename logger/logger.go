package logger

import (
	"os"
	"time"

	"github.com/mrumyantsev/logx"
)

type Logger struct {
	isDisableColors bool
	timeFormat      string
	outputStream    *os.File
}

func New() *Logger {
	return &Logger{
		timeFormat:   logx.TimeFormat,
		outputStream: logx.GetOutputStream(),
	}
}

func (l *Logger) SetDisableColors(isDisableColors bool) {
	l.isDisableColors = isDisableColors
}

func (l *Logger) SetTimeFormat(timeFormat string) {
	l.timeFormat = timeFormat
}

func (l *Logger) SetOutputStream(outputStream *os.File) {
	l.outputStream = outputStream
}

func (l *Logger) WriteLog(datetime time.Time, levelId uint8, message string) error {
	var err error

	if l.isDisableColors {
		_, err = l.outputStream.Write(
			[]byte(
				datetime.Format(l.timeFormat) +
					logx.Space +
					logx.GetLevelText(levelId) +
					logx.Space +
					message +
					logx.EndOfLine,
			),
		)

		return err
	}

	_, err = l.outputStream.Write(
		[]byte(
			logx.GrayColor +
				datetime.Format(l.timeFormat) +
				logx.Space +
				logx.GetLevelColor(levelId) +
				logx.GetLevelText(levelId) +
				logx.Space +
				logx.RegularColor +
				message +
				logx.EndOfLine,
		),
	)

	return err
}
