package multilog

import (
	"os"
)

// Constants, that defines the presentation of logs.
const (
	EmptyString string = ""
	Space       string = " "
	EndOfLine   string = "\n"
	ColorPrefix string = string(27) + "["

	InfoLevel  string = "INF"
	DebugLevel string = "DBG"
	WarnLevel  string = "WRN"
	ErrorLevel string = "ERR"
	FatalLevel string = "FTL"
	PanicLevel string = "PNC"

	RegularColor string = ColorPrefix + "0m"
	GrayColor    string = ColorPrefix + "90m"

	GreenColor   string = ColorPrefix + "32m"
	YellowColor  string = ColorPrefix + "33m"
	RedColor     string = ColorPrefix + "31m"
	BoldRedColor string = ColorPrefix + "1m" + RedColor

	TimeFormat string = "2006-01-02T15:04:05-07:00"
)

const (
	InfoLevelId uint8 = iota
	DebugLevelId
	WarnLevelId
	ErrorLevelId
	FatalLevelId
	PanicLevelId
)

var (
	outputStream = os.Stderr
	levelTexts   = [6]string{
		"INF",
		"DBG",
		"WRN",
		"ERR",
		"FTL",
		"PNC",
	}
	levelColors = [6]string{
		GreenColor,
		YellowColor,
		RedColor,
		BoldRedColor,
		BoldRedColor,
		BoldRedColor,
	}
)

// Get output stream settings, filled with default values.
func GetOutputStream() *os.File {
	return outputStream
}

func GetLevelText(levelId uint8) string {
	return levelTexts[levelId]
}

func GetLevelColor(colorId uint8) string {
	return levelColors[colorId]
}

// Multilog configurational structure.
type Config struct {
	// Disables debug logs to be executed.
	// Default: false
	IsDisableDebugLogs bool

	// Disables warn logs to be executed.
	// Default: false
	IsDisableWarnLogs bool

	// Disables colored text in the stream logs.
	// Default: false
	IsDisableColors bool

	// Defines datetime format in the stream logs.
	// Default: "2006-01-02T15:04:05-07:00"
	TimeFormat string

	// Chooses output data stream for the stream logs.
	// Default: os.Stderr
	OutputStream *os.File
}

// Initialize fields, that were not set by user, with its default values.
func (c *Config) InitEmptyFields() {
	if c.TimeFormat == EmptyString {
		c.TimeFormat = TimeFormat
	}

	if c.OutputStream == nil {
		c.OutputStream = GetOutputStream()
	}
}
