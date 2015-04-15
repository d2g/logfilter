package logfilter

import (
	"strconv"
	"strings"
	"time"
)

// Level represents a logging level.
type Level int

// Standard(ish) Logging Levels.
const (
	Undefined Level = iota
	Trace
	Debug
	Info
	Warning
	Error
	Fatal
	Off
)

// StringToLevel converts a string log level (i.e. "Error") to the coresponding Level (i.e. Error).
func StringToLevel(sl string) Level {
	switch strings.ToLower(sl) {
	case "trace":
		return Trace
	case "debug":
		return Debug
	case "info":
		return Info
	case "warning":
		return Warning
	case "error":
		return Error
	case "fatal":
		return Fatal
	case "off":
		return Off
	}
	return Undefined
}

// LevelToString converts a Level (i.e. Error) to the coresponding string version (i.e. "Error").
func LevelToString(l Level) string {
	switch l {
	case Trace:
		return "Trace"
	case Debug:
		return "Debug"
	case Info:
		return "Info"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	case Fatal:
		return "Fatal"
	case Off:
		return "Off"
	}
	return "Undefined"
}

// StringToLogLine converts the default log sting sent to the LogLine type.
func StringToLogLine(m string) LogLine {
	l := LogLine{}
	l.Message = m

	parts := strings.SplitN(l.Message, " ", 3)
	if len(parts) >= 2 {
		t := strings.Join(parts[:2], " ")
		l.Timestamp, _ = time.Parse("2006/01/02 15:04:05.999999", t)
		l.Message = l.Message[len(t):]
	}

	parts = strings.SplitN(l.Message, ": ", 2)
	if len(parts) >= 2 {
		sp := strings.LastIndex(parts[0], ":")
		l.File = parts[0][:sp]
		l.Line, _ = strconv.Atoi(parts[0][sp+1:])
		l.Message = l.Message[len(parts[0]):]
	}

	l.Message = l.Message[2:]

	return l
}
