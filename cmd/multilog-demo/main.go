package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/mrumyantsev/multilog"
	"github.com/mrumyantsev/multilog/defaults"
	"github.com/mrumyantsev/multilog/log"
)

const (
	// log writers IDs
	_WRITER_FILE     int = 0
	_WRITER_MYSQL    int = 7
	_WRITER_POSTGRES int = 15

	_WRITER_ACCEPTING_EXAMPLE string = "| %s | %s | %s | %s |"
)

func main() {
	logCfg := &multilog.Config{
		IsDisableDebugLogs: false,
		IsDisableWarnLogs:  false,
		InfoLevelText:      "info",
		DebugLevelText:     "debug",
		WarnLevelText:      "warn",
		ErrorLevelText:     "error",
		FatalLevelText:     "fatal",
		ItemSeparatorText:  "\t",
		LineEndingText:     " (for console)\n",
	}

	log.ApplyConfig(logCfg)

	file := &FileController{"logs.txt  ", nil}
	mySql := &DatabaseController{"MySQL     "}
	postgreSql := &DatabaseController{"PostgreSQL"}

	log.AddWriter(_WRITER_POSTGRES, postgreSql)
	log.AddWriter(_WRITER_MYSQL, mySql)

	log.Info("message")

	log.AddWriter(_WRITER_FILE, file)

	log.Debug("message")

	log.RemoveWriter(_WRITER_MYSQL)
	log.RemoveWriter(_WRITER_FILE)

	log.Warn("message")

	log.RemoveWriter(_WRITER_POSTGRES)
	log.AddWriter(_WRITER_FILE, file)

	file.ProvokeError() // look at this

	log.Error("description", errors.New("some error occurred"))

	log.RemoveWriter(_WRITER_MYSQL)
	log.AddWriter(_WRITER_POSTGRES, postgreSql)
	log.AddWriter(_WRITER_MYSQL, mySql)
	log.AddWriter(_WRITER_FILE, file)

	log.Fatal("description", errors.New("it crashed, as was planned"))
}

type FileController struct {
	Destination string
	err         error
}

func (f *FileController) WriteLog(datetime time.Time, level string, message string) error {
	if f.err != nil {
		return f.err
	}

	fmt.Println(fmt.Sprintf(
		_WRITER_ACCEPTING_EXAMPLE, f.Destination, datetime.Format(defaults.TIME_FORMAT), level, message))

	return nil
}

func (f *FileController) ProvokeError() {
	f.err = errors.New("i can't write this log to file")

	fmt.Println("*** An error is provoken to crash the file writer, ha-ha! ***")
}

type DatabaseController struct {
	Destination string
}

func (d *DatabaseController) WriteLog(datetime time.Time, level string, message string) error {
	fmt.Println(fmt.Sprintf(
		_WRITER_ACCEPTING_EXAMPLE, d.Destination, datetime.Format(defaults.TIME_FORMAT), level, message))

	return nil
}
