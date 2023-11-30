package main

import (
	"errors"
	"fmt"

	log "github.com/mrumyantsev/multilog"
)

const (
	_WRITTEN_FILE_EXAMPLE     string = "dest: %s, time: %s, type: %s, msg: %s"
	_WRITTEN_DATABASE_EXAMPLE string = "| %s | %s | %s | %s |"
)

func main() {
	log.ExitStatusCodeWhenFatal = 123
	log.IsEnableDebugLogs = false

	file := &FileController{"./non-ordinary-logs.txt"}
	mySql := &DatabaseController{"MySQL"}
	postgreSql := &DatabaseController{"PostgreSQL"}

	log.RegisterWriter("db1", mySql)
	log.RegisterWriter("db2", postgreSql)

	log.Info("info message")

	log.Debug("debug message")

	log.RegisterWriter("file", file)

	log.Error("error description", errors.New("errors happens"))

	log.UnregisterWriter("db1")

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
