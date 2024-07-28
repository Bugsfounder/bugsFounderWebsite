package logger

import (
	"fmt"
	"log"
	"os"
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
	Fatal
	Panic
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

// getColorCode returns the ANSI color code for a given log level.
func getColorCode(level LogLevel) string {
	switch level {
	case Debug:
		return "\033[32m" // Green
	case Info:
		return "\033[34m" // Blue
	case Warning:
		return "\033[33m" // Yellow
	case Error:
		return "\033[31m" // Red
	case Fatal:
		return "\033[35m" // Magenta
	case Panic:
		return "\033[41m" // Red Background
	default:
		return "\033[0m" // Reset
	}
}

// customLog formats the log record into the desired format, including the function name, filename, and line number.
func (l *Logger) customLog(level LogLevel, format string, args ...interface{}) {
	funcName, file, line := getCallerInfo()
	msg := fmt.Sprintf(format, args...)
	now := time.Now().Format("2006-01-02:15:04:05,999")
	colorCode := getColorCode(level)
	resetCode := "\033[0m" // Reset color

	logMessage := fmt.Sprintf("%s %s %-8s [%s:%d] %s() %s%s", now, colorCode, strings.ToUpper(level.String()), file, line, funcName, msg, resetCode)
	l.Logger.Output(3, logMessage)

	if level == Fatal {
		// Exit with status code 1 for fatal errors
		os.Exit(1)
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

// Fatal logs a fatal error message and exits the application.
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.customLog(Fatal, format, args...)
}

func (l *Logger) Panic(message string) {
	funcName, file, line := getCallerInfo()
	now := time.Now().Format("2006-01-02:15:04:05,999")
	const (
		RedBG = "\033[41m"
		White = "\033[97m"
		Reset = "\033[0m"
	)
	fmt.Printf("%s %s %-8s %s:%d %s() Panic: %s%s\n", now, RedBG, White, file, line, funcName, message, Reset)
	os.Exit(1)
}

// String converts a LogLevel to its string representation.
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
	case Fatal:
		return "FATAL"
	case Panic:
		return "PANIC"
	default:
		return "UNKNOWN"
	}
}

var logger_singleton *Logger = AppLogger(os.Stdout)

func Logging() *Logger {
	return logger_singleton
}

func Test() {
	logger := AppLogger(os.Stdout)
	logger.Debug("custom message")
	logger.Info("custom message")
	logger.Error("custom message")

	// or
	Logging().Debug("message")
	Logging().Info("message")
	Logging().Error("message")

	Logging().Info("This is an info message.")
	Logging().Error("This is an error message.")
	Logging().Warning("This is a warning message.")
	Logging().Debug("This is a debug message.")
	Logging().Fatal("This is a fatal message.")
	Logging().Panic("This is a panic message.")
}
