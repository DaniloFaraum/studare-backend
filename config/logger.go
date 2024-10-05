package config

import (
	"io"
	"log"
	"os"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Orange = "\033[38;5;208m"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func messageColor(prefix string, message string, color string) string {
	return Orange + prefix + Reset + color + message + Reset
}

func NewLogger(prefix string) *Logger {
	writer := os.Stdout                                    //Defines the writer as the deafault OS output
	logger := log.New(writer, prefix, log.Ldate|log.Ltime) //Creates a logger with date and time in each log entry

	return &Logger{
		debug:   log.New(writer, messageColor(prefix, " DEBUG: ", Blue), logger.Flags()),
		info:    log.New(writer, messageColor(prefix, " INFO: ", Blue), logger.Flags()),
		warning: log.New(writer, messageColor(prefix, " WARNING: ", Yellow), logger.Flags()),
		err:     log.New(writer, messageColor(prefix, " ERROR: ", Red), logger.Flags()),
		writer:  writer,
	}
}

// Non-formated logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}
func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Formated logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
