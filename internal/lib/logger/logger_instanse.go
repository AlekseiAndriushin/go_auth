package logger

import (
	"fmt"
	"sync"
)

type DefaultLogger struct {
	Level LogLevel
	mu    sync.Mutex
}

var (
	defaultLogger = &DefaultLogger{
		Level: Info, 
	}
)

func (l *DefaultLogger) Log(level LogLevel, message string) {
	if l.Level > level {
		return 
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Printf("[%v] %s\n", level, message)
}

func SetLogLevel(level LogLevel) {
	defaultLogger.Level = level
}

func LogInfo(message string) {
	defaultLogger.Log(Info, message)
}

func LogDebug(message string) {
	defaultLogger.Log(Debug, message)
}

func LogError(message string) {
	defaultLogger.Log(Error, message)
}
