package multilog

import (
	"os"

	"github.com/mrumyantsev/multilog/defaults"
)

// Multilog configurational structure.
type Config struct {
	OutputStream *os.File

	// affects only output stream logs
	// (showing in console/terminal)
	IsDisableColors bool

	IsDisableDebugLogs bool
	IsDisableWarnLogs  bool

	TimeFormat string

	ItemSeparatorText string
	LineEndingText    string
	InfoLevelText     string
	DebugLevelText    string
	WarnLevelText     string
	ErrorLevelText    string
	FatalLevelText    string

	DatetimeColor   string
	InfoLevelColor  string
	DebugLevelColor string
	WarnLevelColor  string
	ErrorLevelColor string
	FatalLevelColor string
	MessageColor    string
}

// Get initialized configuration instance.
func NewConfig() *Config {
	cfg := &Config{}

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
		colors := defaults.GetColors()

		if c.DatetimeColor == defaults.EMPTY_STRING {
			c.DatetimeColor = colors.DatetimeColor
		}
		if c.InfoLevelColor == defaults.EMPTY_STRING {
			c.InfoLevelColor = colors.InfoLevelColor
		}
		if c.DebugLevelColor == defaults.EMPTY_STRING {
			c.DebugLevelColor = colors.DebugLevelColor
		}
		if c.WarnLevelColor == defaults.EMPTY_STRING {
			c.WarnLevelColor = colors.WarnLevelColor
		}
		if c.ErrorLevelColor == defaults.EMPTY_STRING {
			c.ErrorLevelColor = colors.ErrorLevelColor
		}
		if c.FatalLevelColor == defaults.EMPTY_STRING {
			c.FatalLevelColor = colors.FatalLevelColor
		}
		if c.MessageColor == defaults.EMPTY_STRING {
			c.MessageColor = colors.MessageColor
		}
	}
}
