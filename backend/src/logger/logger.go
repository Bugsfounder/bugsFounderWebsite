package logger

import (
	"fmt"
	"log"
	"os"

	// "path/filepath"
	"runtime"
	"strings"
	"time"
)

// LogLevel represents different log levels.
type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

// Logger is a custom logger that wraps the standard log.Logger.
type Logger struct {
	*log.Logger
}

// Logging creates a new Logger.
func AppLogger(out *os.File) *Logger {
	return &Logger{
		Logger: log.New(out, "", 0),
	}
}

// getCallerInfo retrieves the function name, filename, and line number where the logger function is called.
func getCallerInfo() (string, string, int) {
	pc, file, line, _ := runtime.Caller(3) // Skip 3 frames to get the caller details
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "<module>", "unknown", 0
	}
	funcName := fn.Name()
	parts := strings.Split(funcName, ".")
	if len(parts) > 1 {
		funcName = parts[len(parts)-1]
	} else {
		funcName = "<module>"
	}
	return funcName, file, line
}

// customLog formats the log record into the desired format, including the function name, filename, and line number.
func (l *Logger) customLog(level LogLevel, format string, args ...interface{}) {
	funcName, file, line := getCallerInfo()
	msg := fmt.Sprintf(format, args...)
	now := time.Now().Format("2006-01-02:15:04:05,999")
	switch level {
	case Debug:
		// For Debug, show full path and function name
		logMessage := fmt.Sprintf("%s DEBUG    [%s:%d] %s() %s", now, file, line, funcName, msg)
		l.Logger.Output(3, logMessage)
	case Info:
		fallthrough
	case Warning:
		fallthrough
	case Error:
		// For Info, Warning, and Error, show relative path (containing only the package name) and function name
		// _, filename := filepath.Split(file)
		logMessage := fmt.Sprintf("%s %-8s [%s:%d] %s() %s", now, strings.ToUpper(level.String()), file, line, funcName, msg)
		l.Logger.Output(3, logMessage)
	}
}

// Debug logs a debug message.
func (l *Logger) Debug(format string, args ...interface{}) {
	l.customLog(Debug, format, args...)
}

// Info logs an info message.
func (l *Logger) Info(format string, args ...interface{}) {
	l.customLog(Info, format, args...)
}

// Warning logs a warning message.
func (l *Logger) Warning(format string, args ...interface{}) {
	l.customLog(Warning, format, args...)
}

// Error logs an error message.
func (l *Logger) Error(format string, args ...interface{}) {
	l.customLog(Error, format, args...)
}

func (l LogLevel) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

var logger_singleton *Logger = AppLogger(os.Stdout)

func Logging() *Logger {
	return logger_singleton
}

// testing logger
func LogInsideFunction() {
	logger := AppLogger(os.Stdout)
	logger.Debug("custom message")
	logger.Info("custom message")
	logger.Error("custom message")

	// or
	Logging().Debug("message")
	Logging().Info("message")
	Logging().Error("message")
}
