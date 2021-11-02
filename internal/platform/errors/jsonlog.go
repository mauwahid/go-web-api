package errors

import (
	"encoding/json"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

type Level int8

const (
	LevelInfo Level = iota
	LevelError
	LevelFatal
	LevelOff
)

func (l Level) string() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	case LevelOff:
		return "OFF"
	default:
		return ""
	}
}

type Logger struct {
	out      io.Writer
	minLevel Level
	mu       sync.Mutex
}

func New(out io.Writer, minLevel Level) *Logger {
	return &Logger{
		out:      out,
		minLevel: minLevel,
	}
}

func (l *Logger) Info(traceID, message string, prop map[string]string) {
	l.print(LevelInfo, traceID, message, prop)
}

func (l *Logger) Fatal(traceID, message string, err error, prop map[string]string) {
	l.print(LevelFatal, traceID, message+"|"+err.Error(), prop)
	os.Exit(1)
}

func (l *Logger) Error(traceID, message string, err error, prop map[string]string) {
	l.print(LevelError, traceID, message+"|"+err.Error(), prop)
}

func (l *Logger) print(level Level, traceID, message string, prop map[string]string) (int, error) {
	if level < l.minLevel {
		return 0, nil
	}
	aux := struct {
		Level      string            `json:"level"`
		Time       string            `json:"time"`
		Message    string            `json:"message"`
		Properties map[string]string `json:"properties,omitempty"`
		TraceID    string            `json:"trace,omitempty"`
		Debug      string            `json:"debug,omitempty"`
	}{
		Level:      level.string(),
		Time:       time.Now().Local().Format(time.RFC3339),
		Properties: prop,
		TraceID:    traceID,
	}

	if level >= LevelError {
		aux.Debug = string(debug.Stack())
	}

	var line []byte
	line, err := json.Marshal(aux)
	if err != nil {
		line = []byte(LevelError.string() + " unable to marshal log message " + err.Error())
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	return l.out.Write(append(line, '\n'))
}

func (l *Logger) Write(message []byte) (n int, err error) {
	return l.print(LevelError, "", string(message), nil)
}
