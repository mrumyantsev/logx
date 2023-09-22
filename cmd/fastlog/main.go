package main

import (
	"errors"

	"github.com/mrumyantsev/fastlog"
)

func main() {
	fastlog.Info("info message")
	fastlog.Debug("debug message")
	fastlog.Error("error description", errors.New("errors happens"))
	fastlog.Fatal("fatal error description", errors.New("fatal errors happens"))
}
