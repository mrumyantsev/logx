package defaults

const (
	TIME_FORMAT string = "2006-01-02T15:04:05-07:00"

	ITEM_SEPARATOR_TEXT string = " "
	LINE_ENDING_TEXT    string = "\n"
	INFO_LEVEL_TEXT     string = "INF"
	DEBUG_LEVEL_TEXT    string = "DBG"
	WARN_LEVEL_TEXT     string = "WRN"
	ERROR_LEVEL_TEXT    string = "ERR"
	FATAL_LEVEL_TEXT    string = "FTL"

	EMPTY_STRING string = ""
)

type Colors struct {
	DatetimeColor   string
	InfoLevelColor  string
	DebugLevelColor string
	WarnLevelColor  string
	ErrorLevelColor string
	FatalLevelColor string
	MessageColor    string
}

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
