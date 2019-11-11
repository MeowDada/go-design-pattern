package logger

import (
	"fmt"
)

type EmailLogger struct{
	next Logger
	mask LogLevel
}

func NewEmailLogger(mask LogLevel) EmailLogger {
	return EmailLogger{
		mask: mask,
	}
}

func (el *EmailLogger) SetNext(logger Logger) {
	el.next = logger
}

func (el EmailLogger) Log(msg string, level LogLevel) {
	if mask(level, el.mask) {
		out := fmt.Sprintf("[EMAIL]: %s", msg)
		fmt.Println(out)
	}
	if el.next != nil {
		el.next.Log(msg, level)
	}
}