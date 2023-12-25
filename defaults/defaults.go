package defaults

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

var (
	outputStream = os.Stderr
)

// Get output stream settings, filled with default values.
func GetOutputStream() *os.File {
	return outputStream
}
