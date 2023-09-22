package main

import (
	"github.com/mrumyantsev/fastlog"
)

func main() {
	fastlog.Info("info message")
	fastlog.Debug("debug message")
	fastlog.Error("error message")
	fastlog.Fatal("fatal message")
}
