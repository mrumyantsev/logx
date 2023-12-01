package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/mrumyantsev/multilog"
	"github.com/mrumyantsev/multilog/log"
)

const (
	_WRITTEN_FILE_EXAMPLE     string = "/ %s / %s / %s / %s /"
	_WRITTEN_DATABASE_EXAMPLE string = "| %s | %s | %s | %s |"

	_LWID_MYSQL int = iota
	_LWID_POSTGRESQL
	_LWID_FILE
)

func main() {
	multilog.FatalExitStatusCode = 123
	multilog.IsEnableDebugLogs = false

	file := &FileController{"./non-ordinary-logs.txt"}
	mySql := &DatabaseController{"MySQL"}
	postgreSql := &DatabaseController{"PostgreSQL"}

	log.AddLogWriter(_LWID_MYSQL, mySql)
	log.AddLogWriter(_LWID_POSTGRESQL, postgreSql)

	log.Info("info message")

	log.Debug("debug message")

	log.AddLogWriter(_LWID_FILE, file)

	log.Error("error description", errors.New("errors happens"))

	log.RemoveLogWriter(_LWID_MYSQL)

	log.DisableLogWriter(_LWID_FILE)

	log.Fatal("fatal error description", errors.New("fatal errors happens"))
}

type FileController struct {
	Destination string
}

func (f *FileController) WriteLog(datetime time.Time, level string, message string) error {
	fmt.Println(fmt.Sprintf(
		_WRITTEN_FILE_EXAMPLE, f.Destination, datetime, level, message))

	return nil
}

type DatabaseController struct {
	Destination string
}

func (d *DatabaseController) WriteLog(datetime time.Time, level string, message string) error {
	fmt.Println(fmt.Sprintf(
		_WRITTEN_DATABASE_EXAMPLE, d.Destination, datetime, level, message))

	return nil
}
