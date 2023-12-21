package defaults

import (
	"os"
)

// Constants, that defines the presentation of logs.
const (
	EMPTY_STRING string = ""
	COLOR_PREFIX string = string(27) + "["

	TIME_FORMAT string = "2006-01-02T15:04:05-07:00"

	ITEM_SEPARATOR string = " "
	LINE_ENDING    string = "\n"
	INFO_LEVEL     string = "INF"
	DEBUG_LEVEL    string = "DBG"
	WARN_LEVEL     string = "WRN"
	ERROR_LEVEL    string = "ERR"
	FATAL_LEVEL    string = "FTL"
	PANIC_LEVEL    string = "PNC"

	GRAY_COLOR     string = COLOR_PREFIX + "90m"
	GREEN_COLOR    string = COLOR_PREFIX + "32m"
	YELLOW_COLOR   string = COLOR_PREFIX + "33m"
	RED_COLOR      string = COLOR_PREFIX + "31m"
	RED_BOLD_COLOR string = COLOR_PREFIX + "1m" + RED_COLOR
	REGULAR_COLOR  string = COLOR_PREFIX + "0m"
)

var (
	outputStream = os.Stderr
)

// Get output stream settings, filled with default values.
func GetOutputStream() *os.File {
	return outputStream
}
