package multilog

import (
	"os"

	"github.com/mrumyantsev/multilog/defaults"
)

// Multilog configurational structure.
type Config struct {
	// Disables debug logs to be executed.
	// Default: false
	IsDisableDebugLogs bool

	// Disables warn logs to be executed.
	// Default: false
	IsDisableWarnLogs bool

	// Chooses output data stream for the stream logs.
	// Default: os.Stderr
	OutputStream *os.File

	// Disables colored text in the stream logs.
	// Default: false
	IsDisableColors bool

	// Defines datetime format in the stream logs.
	// Default: "2006-01-02T15:04:05-07:00"
	TimeFormat string

	// Defines text between parts in the stream logs.
	// Default: " "
	ItemSeparatorText string

	// Defines conluding text in the stream logs.
	// Default: "\n"
	LineEndingText string

	// Defines info level text in the stream logs.
	// Default: "INF"
	InfoLevelText string

	// Defines debug level text in the stream logs.
	// Default: "DBG"
	DebugLevelText string

	// Defines warn level text in the stream logs.
	// Default: "WRN"
	WarnLevelText string

	// Defines error level text in the stream logs.
	// Default: "ERR"
	ErrorLevelText string

	// Defines fatal level text in the stream logs.
	// Default: "FTL"
	FatalLevelText string

	// Defines datetime color in the stream logs.
	// Default: defaults.COLOR_PREFIX + "90m"
	DatetimeColor string

	// Defines info level color in the stream logs.
	// Default: defaults.COLOR_PREFIX + "32m"
	InfoLevelColor string

	// Defines debug level color in the stream logs.
	// Default: defaults.COLOR_PREFIX + "33m"
	DebugLevelColor string

	// Defines warn level color in the stream logs.
	// Default: defaults.COLOR_PREFIX + "31m"
	WarnLevelColor string

	// Defines error level color in the stream logs.
	// Default: defaults.COLOR_PREFIX + "1m" +
	//          defaults.COLOR_PREFIX + "31m"
	ErrorLevelColor string

	// Defines fatal level color in the stream logs.
	// Default: defaults.COLOR_PREFIX + "1m" +
	//          defaults.COLOR_PREFIX + "31m"
	FatalLevelColor string

	// Defines message color in the stream logs.
	// Default: defaults.COLOR_PREFIX + "0m"
	MessageColor string
}

// Get start configuration, fully initialized with default values.
func GetStartConfig() *Config {
	var cfg *Config = &Config{}

	cfg.InitEmptyFields()

	return cfg
}

// Initialize fields, that were not set by user, with its default values.
func (c *Config) InitEmptyFields() {
	if c.OutputStream == nil {
		c.OutputStream = defaults.GetOutputStream()
	}

	if c.TimeFormat == defaults.EMPTY_STRING {
		c.TimeFormat = defaults.TIME_FORMAT
	}

	if c.ItemSeparatorText == defaults.EMPTY_STRING {
		c.ItemSeparatorText = defaults.ITEM_SEPARATOR_TEXT
	}

	if c.LineEndingText == defaults.EMPTY_STRING {
		c.LineEndingText = defaults.LINE_ENDING_TEXT
	}

	if c.InfoLevelText == defaults.EMPTY_STRING {
		c.InfoLevelText = defaults.INFO_LEVEL_TEXT
	}

	if c.DebugLevelText == defaults.EMPTY_STRING {
		c.DebugLevelText = defaults.DEBUG_LEVEL_TEXT
	}

	if c.WarnLevelText == defaults.EMPTY_STRING {
		c.WarnLevelText = defaults.WARN_LEVEL_TEXT
	}

	if c.ErrorLevelText == defaults.EMPTY_STRING {
		c.ErrorLevelText = defaults.ERROR_LEVEL_TEXT
	}

	if c.FatalLevelText == defaults.EMPTY_STRING {
		c.FatalLevelText = defaults.FATAL_LEVEL_TEXT
	}

	if c.IsDisableColors {
		c.DatetimeColor = defaults.EMPTY_STRING
		c.InfoLevelColor = defaults.EMPTY_STRING
		c.DebugLevelColor = defaults.EMPTY_STRING
		c.WarnLevelColor = defaults.EMPTY_STRING
		c.ErrorLevelColor = defaults.EMPTY_STRING
		c.FatalLevelColor = defaults.EMPTY_STRING
		c.MessageColor = defaults.EMPTY_STRING
	} else {
		if c.DatetimeColor == defaults.EMPTY_STRING {
			c.DatetimeColor = defaults.DATETIME_COLOR
		}

		if c.InfoLevelColor == defaults.EMPTY_STRING {
			c.InfoLevelColor = defaults.INFO_LEVEL_COLOR
		}

		if c.DebugLevelColor == defaults.EMPTY_STRING {
			c.DebugLevelColor = defaults.DEBUG_LEVEL_COLOR
		}

		if c.WarnLevelColor == defaults.EMPTY_STRING {
			c.WarnLevelColor = defaults.WARN_LEVEL_COLOR
		}

		if c.ErrorLevelColor == defaults.EMPTY_STRING {
			c.ErrorLevelColor = defaults.ERROR_LEVEL_COLOR
		}

		if c.FatalLevelColor == defaults.EMPTY_STRING {
			c.FatalLevelColor = defaults.FATAL_LEVEL_COLOR
		}

		if c.MessageColor == defaults.EMPTY_STRING {
			c.MessageColor = defaults.MESSAGE_COLOR
		}
	}
}
