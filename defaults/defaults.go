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

	DATETIME_COLOR    string = COLOR_PREFIX + "90m"
	INFO_LEVEL_COLOR  string = COLOR_PREFIX + "32m"
	DEBUG_LEVEL_COLOR string = COLOR_PREFIX + "33m"
	WARN_LEVEL_COLOR  string = COLOR_PREFIX + "31m"
	ERROR_LEVEL_COLOR string = COLOR_PREFIX + "1m" + WARN_LEVEL_COLOR
	FATAL_LEVEL_COLOR string = ERROR_LEVEL_COLOR
	PANIC_LEVEL_COLOR string = ERROR_LEVEL_COLOR
	MESSAGE_COLOR     string = COLOR_PREFIX + "0m"
)

var (
	outputStream = os.Stderr
)

// Get output stream settings, filled with default values.
func GetOutputStream() *os.File {
	return outputStream
}
