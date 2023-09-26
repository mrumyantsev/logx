package fastlog

import (
	"os"
	"time"
)

type logMessage struct {
	datetime    string
	messageType *string
	message     *string
}

const (
	WHITESPACE         string = " "
	DOT                string = "."
	COMMA              string = ","
	SEMICOLON          string = ";"
	COLON              string = ":"
	TAB                string = "\t"
	NEW_LINE           string = "\n"
	ERROR_WORD         string = ". error: "
	INFO_MESSAGE_TYPE  string = "INF"
	DEBUG_MESSAGE_TYPE string = "DBG"
	ERROR_MESSAGE_TYPE string = "ERR"
	FATAL_MESSAGE_TYPE string = "FTL"
	TIME_FORMAT        string = "2006-01-02T15:04:05-07:00"
)

var (
	stdoutFile        *os.File    = os.Stdout
	stderrFile        *os.File    = os.Stderr
	IsEnableDebugLogs bool        = true
	ItemSeparator     string      = WHITESPACE
	LineEnding        string      = NEW_LINE
	InfoMessageType   string      = INFO_MESSAGE_TYPE
	DebugMessageType  string      = DEBUG_MESSAGE_TYPE
	ErrorMessageType  string      = ERROR_MESSAGE_TYPE
	FatalMessageType  string      = FATAL_MESSAGE_TYPE
	TimeFormat        string      = TIME_FORMAT
	logMsg            *logMessage = &logMessage{}
)

func Info(msg string) {
	logMsg.datetime = time.Now().Format(TimeFormat)
	logMsg.messageType = &InfoMessageType
	logMsg.message = &msg

	stdoutFile.Write(
		[]byte(logMsg.datetime +
			ItemSeparator +
			*logMsg.messageType +
			ItemSeparator +
			*logMsg.message +
			LineEnding,
		),
	)
}

func Debug(msg string) {
	logMsg.datetime = time.Now().Format(TimeFormat)
	logMsg.messageType = &DebugMessageType
	logMsg.message = &msg

	stdoutFile.Write(
		[]byte(logMsg.datetime +
			ItemSeparator +
			*logMsg.messageType +
			ItemSeparator +
			*logMsg.message +
			LineEnding,
		),
	)
}

func Error(desc string, err error) {
	var (
		msg string = desc + ERROR_WORD + err.Error()
	)

	logMsg.datetime = time.Now().Format(TimeFormat)
	logMsg.messageType = &ErrorMessageType
	logMsg.message = &msg

	stderrFile.Write(
		[]byte(logMsg.datetime +
			ItemSeparator +
			*logMsg.messageType +
			ItemSeparator +
			*logMsg.message +
			LineEnding,
		),
	)
}

func Fatal(desc string, err error) {
	var (
		msg string = desc + ERROR_WORD + err.Error()
	)

	logMsg.datetime = time.Now().Format(TimeFormat)
	logMsg.messageType = &FatalMessageType
	logMsg.message = &msg

	stderrFile.Write(
		[]byte(logMsg.datetime +
			ItemSeparator +
			*logMsg.messageType +
			ItemSeparator +
			*logMsg.message +
			LineEnding,
		),
	)

	os.Exit(1)
}
