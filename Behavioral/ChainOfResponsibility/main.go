package main

import (
	logger "github.com/MeowDada/go-design-pattern/Behavioral/ChainOfResponsibility/pkg/logger"
)

func main() {
	
	emailLogger := logger.NewEmailLogger(logger.LogLevelFatal | logger.LogLevelError)
	fileLogger := logger.NewFileLogger(logger.LogLevelFatal | logger.LogLevelDebug)
	consoleLogger := logger.NewConsoleLogger(logger.LogLevelAll)

	emailLogger.SetNext(&fileLogger)
	fileLogger.SetNext(&consoleLogger)

	emailLogger.Log("Hello world" ,logger.LogLevelFatal)
	emailLogger.Log("Hello world" ,logger.LogLevelDebug)
}