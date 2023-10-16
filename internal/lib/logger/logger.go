package logger

type Logger interface {
    Log(level LogLevel, message string)
}

type LogLevel int

const (
    Info LogLevel = iota
    Debug
    Error
)
