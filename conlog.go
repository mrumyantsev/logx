package multilog

import (
	"os"
	"time"
)

type ConLog struct {
	outputStream    *os.File
	isDisableColors bool
	timeFormat      string
}

func NewConLog() *ConLog {
	conLog := &ConLog{}

	conLog.outputStream = GetOutputStream()

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
					Space +
					GetLevelText(levelId) +
					Space +
					message +
					EndOfLine,
			),
		)

		return err
	}

	_, err = c.outputStream.Write(
		[]byte(
			GrayColor +
				datetime.Format(c.timeFormat) +
				Space +
				GetLevelColor(levelId) +
				GetLevelText(levelId) +
				Space +
				RegularColor +
				message +
				EndOfLine,
		),
	)

	return err
}
