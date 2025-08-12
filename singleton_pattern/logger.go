package singletonpattern

import (
	"fmt"
	"os"
	"sync"
)

type Logger struct {
	file *os.File
}

var (
	instance *Logger
	once     sync.Once
)

func GetInstance() *Logger {
	once.Do(func() {
		file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Sprintf("Failed to open log file: %v", err))
		}
		instance = &Logger{file: file}
	})
	return instance
}

func (l *Logger) Log(message string) {
	_, _ = l.file.WriteString(message + "\n")
}
