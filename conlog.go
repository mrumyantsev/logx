package multilog

import (
	"os"
	"time"
)

type ConLog struct {
	isDisableColors bool
	timeFormat      string
	outputStream    *os.File
}

func NewConLog() *ConLog {
	conLog := &ConLog{}

	conLog.isDisableColors = false
	conLog.timeFormat = TimeFormat
	conLog.outputStream = GetOutputStream()

	return conLog
}

func (c *ConLog) SetDisableColors(isDisableColors bool) {
	c.isDisableColors = isDisableColors
}

func (c *ConLog) SetTimeFormat(timeFormat string) {
	c.timeFormat = timeFormat
}

func (c *ConLog) SetOutputStream(outputStream *os.File) {
	c.outputStream = outputStream
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
