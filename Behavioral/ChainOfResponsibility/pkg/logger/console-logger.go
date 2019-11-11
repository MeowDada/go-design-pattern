package logger

import (
	"fmt"
)

type ConsoleLogger struct{
	next Logger
	mask LogLevel
}

func NewConsoleLogger(mask LogLevel) ConsoleLogger {
	return ConsoleLogger{
		mask: mask,
	}
}

func (cl *ConsoleLogger) SetNext(logger Logger) {
	cl.next = logger
}

func (cl ConsoleLogger) Log(msg string, level LogLevel) {
	if mask(level, cl.mask) {
		out := fmt.Sprintf("[CONSOLE]: %s", msg)
		fmt.Println(out)
	}
	if cl.next != nil {
		cl.next.Log(msg, level)
	}
}