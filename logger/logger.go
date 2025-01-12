package logger

import (
	"log"
	"log/slog"
	"os"

	"github.com/fatih/color"

	"github.com/avila-r/bitclient/config"
)

var (
	logger  = log.New(os.Stdout, "[bitclient]", log.Ldate|log.Ltime|log.Lmsgprefix)
	printer = log.New(os.Stdout, "", 0)

	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

func Info(v ...any) {
	logger.SetPrefix("[INFO] ")
	logger.Print(v...)
}

func Infof(format string, v ...any) {
	logger.SetPrefix("[INFO] ")
	logger.Printf(format, v...)
}

func Error(v ...any) {
	logger.SetPrefix(red("[ERROR] "))
	logger.Print(v...)
}

func Errorf(format string, v ...any) {
	logger.SetPrefix(red("[ERROR] "))
	logger.Printf(format, v...)
}

func Warn(v ...any) {
	logger.SetPrefix(yellow("[WARN] "))
	logger.Print(v...)
}

func Warnf(format string, v ...any) {
	logger.SetPrefix(yellow("[WARN] "))
	logger.Printf(format, v...)
}

func Debug(v ...any) {
	if !config.Get().Advanced.Debug {
		return // Do nothing if debug is false
	}
	logger.SetPrefix(cyan("[DEBUG] "))
	logger.Print(v...)
}

func Debugf(format string, v ...any) {
	if !config.Get().Advanced.Debug {
		return // Do nothing if debug is false
	}
	logger.SetPrefix(cyan("[DEBUG] "))
	logger.Printf(format, v...)
}

func Print(v ...any) {
	printer.Print(v...)
}

func Printf(format string, v ...any) {
	printer.Printf(format, v...)
}

func init() {
	if config.Get().Advanced.Debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
}
