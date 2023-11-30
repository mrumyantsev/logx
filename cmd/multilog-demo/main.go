package main

import (
	"errors"
	"fmt"

	log "github.com/mrumyantsev/multilog"
)

const (
	_WRITTEN_FILE_EXAMPLE     string = "/ %s / %s / %s / %s /"
	_WRITTEN_DATABASE_EXAMPLE string = "| %s | %s | %s | %s |"

	_LWID_MYSQL int = iota
	_LWID_POSTGRESQL
	_LWID_FILE
)

func main() {
	log.FatalExitStatusCode = 123
	log.IsEnableDebugLogs = false

	file := &FileController{"./non-ordinary-logs.txt"}
	mySql := &DatabaseController{"MySQL"}
	postgreSql := &DatabaseController{"PostgreSQL"}

	log.RegisterWriter(_LWID_MYSQL, mySql)
	log.RegisterWriter(_LWID_POSTGRESQL, postgreSql)

	log.Info("info message")

	log.Debug("debug message")

	log.RegisterWriter(_LWID_FILE, file)

	log.Error("error description", errors.New("errors happens"))

	log.UnregisterWriter(_LWID_MYSQL)

	log.DisableWriter(_LWID_FILE)

	log.Fatal("fatal error description", errors.New("fatal errors happens"))
}

type FileController struct {
	Destination string
}

func (f *FileController) WriteLog(datetime *string, messageType *string, message *string) error {
	fmt.Println(fmt.Sprintf(
		_WRITTEN_FILE_EXAMPLE, f.Destination, *datetime, *messageType, *message))

	return nil
}

type DatabaseController struct {
	Destination string
}

func (d *DatabaseController) WriteLog(datetime *string, messageType *string, message *string) error {
	fmt.Println(fmt.Sprintf(
		_WRITTEN_DATABASE_EXAMPLE, d.Destination, *datetime, *messageType, *message))

	return nil
}
