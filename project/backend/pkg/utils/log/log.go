package log

import (
	"log"
	"os"
)

type environment string

// Development is the const for development environment.
const Development environment = "dev"

// Production is the const for production environment.
const Production environment = "prod"

var (
	debugLog = log.New(os.Stderr, "[DEBUG] ", log.Ldate|log.Ltime)
	infoLog  = log.New(os.Stderr, "[INFO] ", log.Ldate|log.Ltime)
	warnLog  = log.New(os.Stderr, "[WARN] ", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)
	fatalLog = log.New(os.Stderr, "[FATAL] ", log.Ldate|log.Ltime)
)

// Log is the struct for Log.
type Log struct {
	mode environment
}

// Debug print a message in debug level in a non production environment.
func (l *Log) Debug(message ...any) {
	if l.mode != Production {
		debugLog.Println(message...)
	}
}

// Info print a message in info level.
func (l *Log) Info(message ...any) {
	infoLog.Println(message...)
}

// Warn print a message in warn level.
func (l *Log) Warn(message ...any) {
	warnLog.Println(message...)
}

// Error print a message in error level.
func (l *Log) Error(message ...any) {
	errorLog.Println(message...)
}

// Fatal print a message in error level and exit program.
func (l *Log) Fatal(message ...any) {
	fatalLog.Fatalln(message...)
}

// NewLog create a new with the environment.
func NewLog(env environment) *Log {
	return &Log{mode: env}
}
