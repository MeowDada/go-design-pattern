package logger

type LogLevel uint32

const (
	LogLevelFatal LogLevel = 1 << iota
	LogLevelError
	LogLevelWarn
	LogLevelNotice
	LogLevelInfo
	LogLevelDebug
	LogLevelAll = 0xFFFFFFFF
)

type Logger interface {
	Log(string, LogLevel)
	SetNext(Logger)
}

func mask(level LogLevel, mask LogLevel) bool {
	return level&mask==level
}