package logger

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/fatih/color"

	"github.com/avila-r/bitclient/config"
)

// The logger and printer variables are used for logging messages and outputting formatted text.
var (
	// logger is the main logger instance used for logging with prefixes.
	logger = log.New(os.Stdout, "[bitclient]", log.Ldate|log.Ltime|log.Lmsgprefix)
	// printer is used for printing simple output without any prefixes.
	printer = log.New(os.Stdout, "", 0)

	// Predefined colors for colored text output.
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

// Info logs an info message.
func Info(v ...any) {
	logger.SetPrefix("[INFO] ")
	logger.Print(v...)
}

// Infof logs a formatted info message.
func Infof(format string, v ...any) {
	logger.SetPrefix("[INFO] ")
	logger.Printf(format, v...)
}

// Error logs an error message.
func Error(v ...any) {
	logger.SetPrefix(red("[ERROR] "))
	logger.Print(v...)
}

// Errorf logs a formatted error message.
func Errorf(format string, v ...any) {
	logger.SetPrefix(red("[ERROR] "))
	logger.Printf(format, v...)
}

// Fatal logs an error message and exits the program with status 1.
func Fatal(v ...any) {
	logger.SetPrefix(red("[ERROR] "))
	logger.Print(v...)
	os.Exit(1)
}

// Fatalf logs a formatted error message and exits the program with status 1.
func Fatalf(format string, v ...any) {
	logger.SetPrefix(red("[ERROR] "))
	logger.Printf(format, v...)
	os.Exit(1)
}

// Warn logs a warning message.
func Warn(v ...any) {
	logger.SetPrefix(yellow("[WARN] "))
	logger.Print(v...)
}

// Warnf logs a formatted warning message.
func Warnf(format string, v ...any) {
	logger.SetPrefix(yellow("[WARN] "))
	logger.Printf(format, v...)
}

// Debug logs a debug message if debugging is enabled in the configuration.
func Debug(v ...any) {
	if !config.Get().Advanced.Debug {
		return // Do nothing if debug is false
	}
	logger.SetPrefix(fmt.Sprintf("%s %s", caller(), cyan("[DEBUG] ")))
	logger.Print(v...)
}

// Debugf logs a formatted debug message if debugging is enabled in the configuration.
func Debugf(format string, v ...any) {
	if !config.Get().Advanced.Debug {
		return // Do nothing if debug is false
	}
	logger.SetPrefix(fmt.Sprintf("%s %s", caller(), cyan("[DEBUG] ")))
	logger.Printf(format, v...)
}

// Print outputs the message without any prefix.
func Print(v ...any) {
	printer.Print(v...)
}

// Printf outputs a formatted message without any prefix.
func Printf(format string, v ...any) {
	printer.Printf(format, v...)
}

// init initializes the logging system by setting the log level for `slog` if debugging is enabled.
func init() {
	if config.Get().Advanced.Debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
}
