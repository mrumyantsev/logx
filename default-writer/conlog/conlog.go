package conlog

import (
	"os"
	"time"

	"github.com/mrumyantsev/multilog"
)

type ConLog struct {
	outputStream    *os.File
	isDisableColors bool
	timeFormat      string
}

func New() *ConLog {
	conLog := &ConLog{}

	conLog.outputStream = multilog.GetOutputStream()

	conLog.isDisableColors = false

	conLog.timeFormat = "2006-01-02T15:04:05-07:00"

	return conLog
}

func (c *ConLog) WriteLog(datetime time.Time, levelId uint8, message string) error {
	var err error

	if c.isDisableColors {
		_, err = c.outputStream.Write(
			[]byte(
				datetime.Format(c.timeFormat) +
					multilog.Space +
					multilog.GetLevelText(levelId) +
					multilog.Space +
					message +
					multilog.EndOfLine,
			),
		)

		return err
	}

	_, err = c.outputStream.Write(
		[]byte(
			multilog.GrayColor +
				datetime.Format(c.timeFormat) +
				multilog.Space +
				multilog.GetLevelColor(levelId) +
				multilog.GetLevelText(levelId) +
				multilog.Space +
				multilog.RegularColor +
				message +
				multilog.EndOfLine,
		),
	)

	return err
}
