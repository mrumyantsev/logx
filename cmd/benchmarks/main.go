package main

import (
	"fmt"
	"log"
	"time"

	logx "github.com/mrumyantsev/logx/log"
	zlog "github.com/rs/zerolog/log"
)

const (
	msg   = "The quick brown fox jumps over the lazy dog."
	times = 200_000
)

func main() {
	tests := []struct {
		loggerName string
		testFunc   func()
		duration   time.Duration
	}{
		{"log", logTest, 0},
		{"zerolog", zlogTest, 0},
		{"logx", logxTest, 0},
	}

	for i := range tests {
		startTime := time.Now()
		tests[i].testFunc()
		tests[i].duration = time.Since(startTime)
	}

	fmt.Println("Time Results")

	for i := range tests {
		fmt.Printf("%7s: %v\n", tests[i].loggerName, tests[i].duration)
	}
}

func logTest() {
	for i := 0; i < times; i++ {
		log.Println(msg)
	}
}

func zlogTest() {
	for i := 0; i < times; i++ {
		zlog.Info().Msg(msg)
	}
}

func logxTest() {
	for i := 0; i < times; i++ {
		logx.Info(msg)
	}
}
