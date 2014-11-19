package logfilter

import (
	"time"
)

// Logging Level represented as an int value.
type Level int

// Standard(ish) Logging Levels.
const (
	_           = iota
	TRACE Level = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
	OFF
)

// Filter Mode used to switch the Filter mode.
type Mode int

// Available log filter modes.
const (
	INCLUDE Mode = 0
	EXCLUDE Mode = 1
)

// Filter structure represents a single filter applied to the log package output.
// You can apply multipule filters where the length of the filename string is
// used to identify which is more specific.
type Filter struct {
	Mode     Mode
	Filename string
	Level    Level
}

// Line is a structure representing a log package line.
type Line struct {
	Timestamp   time.Time
	FileAndLine string
	Message     string

	Level Level
}

// Used to determin is two log Lines are equal.
func (t *Line) Equal(l Line) bool {
	if !t.Timestamp.Equal(l.Timestamp) {
		return false
	}
	if t.FileAndLine != l.FileAndLine {
		return false
	}
	if t.Message != l.Message {
		return false
	}
	if t.Level != l.Level {
		return false
	}
	return true
}
