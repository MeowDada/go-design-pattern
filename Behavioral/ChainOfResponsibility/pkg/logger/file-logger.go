package logger

import (
	"fmt"
)

type FileLogger struct{
	next Logger
	mask LogLevel
}

func NewFileLogger(mask LogLevel) FileLogger {
	return FileLogger{
		mask: mask,
	}
}

func (fl *FileLogger) SetNext(logger Logger) {
	fl.next = logger
}

func (fl FileLogger) Log(msg string, level LogLevel) {
	if mask(level, fl.mask) {
		out := fmt.Sprintf("[FILE]: %s", msg)
		fmt.Println(out)
	}
	if fl.next != nil {
		fl.next.Log(msg, level)
	}
}