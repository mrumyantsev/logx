package main

import (
	"errors"
	"fmt"

	log "github.com/mrumyantsev/fastlog"
)

const (
	WRITTEN_FILE_EXAMPLE     string = "writing to %s:\ntime: %s, type: %s, msg: %s"
	WRITTEN_DATABASE_EXAMPLE string = "writing to %s:\n| %s | %s | %s |"
)

type FileController struct {
	Destination string
}

func (f *FileController) WriteLog(datetime string, messageType string, message string) error {
	fmt.Println(fmt.Sprintf(
		WRITTEN_FILE_EXAMPLE, f.Destination, datetime, messageType, message))

	return nil
}

type DatabaseController struct {
	Destination string
}

func (d *DatabaseController) WriteLog(datetime string, messageType string, message string) error {
	fmt.Println(fmt.Sprintf(
		WRITTEN_DATABASE_EXAMPLE, d.Destination, datetime, messageType, message))

	return nil
}

func main() {
	file := &FileController{"./non-ordinary-logs.txt"}
	mySql := &DatabaseController{"MySQL"}
	postgreSql := &DatabaseController{"PostgreSQL"}

	log.RegisterWriter("file", file)
	log.RegisterWriter("db1", mySql)
	log.RegisterWriter("db2", postgreSql)

	log.Info("info message")
	log.Debug("debug message")
	log.Error("error description", errors.New("errors happens"))

	log.FatalExitStatusCode = 1337 // hack

	log.Fatal("fatal error description", errors.New("fatal errors happens")).
		WriteTo("file").
		WriteTo("db1").
		WriteTo("db2")
}
