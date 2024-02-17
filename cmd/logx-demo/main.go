package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/mrumyantsev/logx"
	"github.com/mrumyantsev/logx/log"
)

const (
	_WRITER_ACCEPTING_EXAMPLE string = "| %s | %s | %s | %s |"
)

func main() {
	logCfg := &logx.Config{
		TimeFormat:      time.Kitchen,
		IsDisableColors: false,
		OutputStream:    os.Stdout,
	}

	log.ApplyConfig(logCfg)

	file := &FileController{"logs.txt  ", nil}
	mySql := &DatabaseController{"MySQL     "}
	postgreSql := &DatabaseController{"PostgreSQL"}

	log.AddWriter(postgreSql)
	log.AddWriter(mySql)

	log.Info("message")

	log.AddWriter(file)

	log.Debug("message")

	log.RemoveWriter(mySql)
	log.RemoveWriter(file)

	log.Warn("message")

	// an error is put to a writer to
	// break it and show, how the logger
	// will handle that situation
	file.ProvokeError()

	log.RemoveWriter(postgreSql)
	log.AddWriter(mySql)
	log.AddWriter(file)

	log.Error("description", errors.New("some error occurred"))

	log.RemoveWriter(mySql)
	log.AddWriter(postgreSql)
	log.AddWriter(file)

	log.Fatal("description", errors.New("app crashed, as was planned"))
}

type FileController struct {
	Destination string
	err         error
}

func (f *FileController) WriteLog(datetime time.Time, levelId uint8, message string) error {
	if f.err != nil {
		return f.err
	}

	fmt.Println(fmt.Sprintf(
		_WRITER_ACCEPTING_EXAMPLE, f.Destination, datetime.Format(logx.TimeFormat), logx.GetLevelText(levelId), message))

	return nil
}

func (f *FileController) ProvokeError() {
	f.err = errors.New("i can't write this log to file")

	fmt.Println("*** An error is provoken to crash the file writer, ha-ha! ***")
}

type DatabaseController struct {
	Destination string
}

func (d *DatabaseController) WriteLog(datetime time.Time, levelId uint8, message string) error {
	fmt.Println(fmt.Sprintf(
		_WRITER_ACCEPTING_EXAMPLE, d.Destination, datetime.Format(logx.TimeFormat), logx.GetLevelText(levelId), message))

	return nil
}