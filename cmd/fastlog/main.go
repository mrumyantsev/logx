package main

import (
	"errors"
	"fmt"

	log "github.com/mrumyantsev/fastlog"
)

const (
	WRITTER_TEMPLATE string = "will be written to %s:\ntime: %s, type: %s, msg: %s"
)

type FileController struct {
}

func (f *FileController) WriteLog(datetime string, messageType string, message string) error {
	const (
		DESTINATION string = "file"
	)

	fmt.Println(fmt.Sprintf(
		WRITTER_TEMPLATE, DESTINATION, datetime, messageType, message))

	return nil
}

type DatabaseController struct {
}

func (d *DatabaseController) WriteLog(datetime string, messageType string, message string) error {
	const (
		DESTINATION string = "database"
	)

	fmt.Println(fmt.Sprintf(
		WRITTER_TEMPLATE, DESTINATION, datetime, messageType, message))

	return nil
}

func main() {
	fileController := &FileController{}
	databaseController := &DatabaseController{}

	log.SetFileLogWriter(fileController)
	log.SetDatabaseLogWriter(databaseController)

	log.Info("info message")
	log.Debug("debug message")
	log.Error("error description", errors.New("errors happens"))
	log.Fatal("fatal error description", errors.New("fatal errors happens")).
		WriteLogToFile().
		WriteLogToDatabase()
}
