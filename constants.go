package logx

import "os"

// Constants, used in LogX.
const (
	EmptyString = ""
	Space       = " "
	EndOfLine   = "\n"
	ColorPrefix = string(27) + "["

	InfoLevelText  = "INF"
	DebugLevelText = "DBG"
	WarnLevelText  = "WRN"
	ErrorLevelText = "ERR"
	FatalLevelText = "FTL"
	PanicLevelText = "PNC"

	RegularColor = ColorPrefix + "0m"
	GrayColor    = ColorPrefix + "90m"
	GreenColor   = ColorPrefix + "32m"
	YellowColor  = ColorPrefix + "33m"
	RedColor     = ColorPrefix + "31m"
	BoldRedColor = ColorPrefix + "1m" + RedColor

	TimeFormat = "2006-01-02 15:04:05 -0700"

	FatalExitCode = 1
)

// LogLevel is a type, to indicate current log message level, and to
// get level text and level color with its value.
type LogLevel uint8

const (
	InfoLevel LogLevel = iota
	DebugLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

var (
	levelTexts = [...]string{
		InfoLevelText,
		DebugLevelText,
		WarnLevelText,
		ErrorLevelText,
		FatalLevelText,
		PanicLevelText,
	}

	levelColors = [...]string{
		GreenColor,
		YellowColor,
		RedColor,
		BoldRedColor,
		BoldRedColor,
		BoldRedColor,
	}

	output *os.File = os.Stderr
)

// LevelText returns text, which indicates the level of log message.
// Like "INF", "ERR", "FTL", etc.
func LevelText(level LogLevel) string {
	return levelTexts[level]
}

// LevelColor returns text code to colorize text in UNIX-terminal.
func LevelColor(level LogLevel) string {
	return levelColors[level]
}

// Output returns stderr OS file, as default logger's output.
func Output() *os.File {
	return output
}
