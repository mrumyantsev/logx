package defaults

import (
	"os"
)

// Constants, that defines the presentation of logs.
const (
	TIME_FORMAT string = "2006-01-02T15:04:05-07:00"

	ITEM_SEPARATOR_TEXT string = " "
	LINE_ENDING_TEXT    string = "\n"
	INFO_LEVEL_TEXT     string = "INF"
	DEBUG_LEVEL_TEXT    string = "DBG"
	WARN_LEVEL_TEXT     string = "WRN"
	ERROR_LEVEL_TEXT    string = "ERR"
	FATAL_LEVEL_TEXT    string = "FTL"

	// only participates in
	// comparisons
	EMPTY_STRING string = ""
)

// Stores output stream settings.
type OutputStreams struct {
	InfoOutputStream  *os.File
	DebugOutputStream *os.File
	WarnOutputStream  *os.File
	ErrorOutputStream *os.File
	FatalOutputStream *os.File
}

// Get output stream settings, filled with default values.
func GetOutputStreams() *OutputStreams {
	streams := &OutputStreams{}

	streams.InfoOutputStream = os.Stderr
	streams.DebugOutputStream = os.Stderr
	streams.WarnOutputStream = os.Stderr
	streams.ErrorOutputStream = os.Stderr
	streams.FatalOutputStream = os.Stderr

	return streams
}

// Stores text colors for output streams.
type Colors struct {
	DatetimeColor   string
	InfoLevelColor  string
	DebugLevelColor string
	WarnLevelColor  string
	ErrorLevelColor string
	FatalLevelColor string
	MessageColor    string
}

// Get text colors, filled with default values.
func GetColors() *Colors {
	colors := &Colors{}

	colors.DatetimeColor = string([]byte{27, 91, 57, 48, 109})
	colors.InfoLevelColor = string([]byte{27, 91, 51, 50, 109})
	colors.DebugLevelColor = string([]byte{27, 91, 51, 51, 109})
	colors.WarnLevelColor = string([]byte{27, 91, 51, 49, 109})
	colors.ErrorLevelColor = string([]byte{27, 91, 49, 109, 27, 91, 51, 49, 109})
	colors.FatalLevelColor = colors.ErrorLevelColor
	colors.MessageColor = string([]byte{27, 91, 48, 109})

	return colors
}
