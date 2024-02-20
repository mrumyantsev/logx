package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/mrumyantsev/logx"
	"github.com/mrumyantsev/logx/log"
)

const (
	receiveTmpl = "| %s | %s | %s | %s |"
)

func main() {
	// create log writers
	file := NewFileWriter("logs.txt")
	mySql := NewDatabaseWriter("MySQL")
	postgreSql := NewDatabaseWriter("PostgreSQL")

	log.AddWriters(postgreSql, mySql)

	log.Info("message")

	log.AddWriters(file)

	log.Debug("message")

	log.RemoveWriters(mySql, file)

	log.Warn("message")

	// make an error in the writer, to break it and to show, how the
	// logger will answer to terminal
	file.ProvokeError()

	log.RemoveWriters(postgreSql)
	log.AddWriters(mySql, file)

	log.Error("description", errors.New("some error occurred"))

	log.RemoveWriters(mySql)
	log.AddWriters(postgreSql, file)

	log.Fatal("description", errors.New("app crashed, as was planned"))
}

type FileWriter struct {
	Destination string
	err         error
}

func NewFileWriter(dest string) *FileWriter {
	return &FileWriter{Destination: dest}
}

func (f *FileWriter) WriteLog(time time.Time, level logx.LogLevel, msg string) error {
	if f.err != nil {
		return f.err
	}

	fmt.Println(fmt.Sprintf(
		receiveTmpl, f.Destination, time.Format(logx.TimeFormat), logx.LevelText(level), msg))

	return nil
}

func (f *FileWriter) ProvokeError() {
	f.err = errors.New("i can't write this log to file")

	fmt.Println("*** An error is provoken to crash the file writer, ha-ha! ***")
}

type DatabaseWriter struct {
	Destination string
}

func NewDatabaseWriter(dest string) *DatabaseWriter {
	return &DatabaseWriter{Destination: dest}
}

func (d *DatabaseWriter) WriteLog(time time.Time, level logx.LogLevel, msg string) error {
	fmt.Println(fmt.Sprintf(
		receiveTmpl, d.Destination, time.Format(logx.TimeFormat), logx.LevelText(level), msg))

	return nil
}
