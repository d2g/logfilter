package logfilter

import (
	"time"
)

type Level int

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

type Mode int

const (
	INCLUDE Mode = 0
	EXCLUDE Mode = 1
)

type Filter struct {
	Mode     Mode
	Filename string
	Level    Level
}

type Line struct {
	Timestamp   time.Time
	FileAndLine string
	Message     string

	Level Level
}

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
