package logfilter

import (
	"io"
	"log"
	"strings"
	"time"
)

// This is the type Format which is the func type reqired for formatting the
// output of log messages.It allows messages logged in one packages as
// `Level: Message` to be output as `[Level] Message`.
type Format func(string, *LogLine, int) []byte

// The Standard Output format
// Level: Message
func StdFormat(prefix string, l *LogLine, f int) []byte {
	var b []byte
	b = append(b, prefix...)

	if f&(log.Ldate|log.Ltime|log.Lmicroseconds) != 0 {
		if f&log.Ldate != 0 {
			year, month, day := l.Timestamp.Date()
			itoa(&b, year, 4)
			b = append(b, '/')
			itoa(&b, int(month), 2)
			b = append(b, '/')
			itoa(&b, day, 2)
			b = append(b, ' ')
		}
		if f&(log.Ltime|log.Lmicroseconds) != 0 {
			hour, min, sec := l.Timestamp.Clock()
			itoa(&b, hour, 2)
			b = append(b, ':')
			itoa(&b, min, 2)
			b = append(b, ':')
			itoa(&b, sec, 2)
			if f&log.Lmicroseconds != 0 {
				b = append(b, '.')
				itoa(&b, l.Timestamp.Nanosecond()/1e3, 6)
			}
			b = append(b, ' ')
		}
	}
	if f&(log.Lshortfile|log.Llongfile) != 0 {
		file := l.File
		if f&log.Lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		b = append(b, file...)
		b = append(b, ':')
		itoa(&b, l.Line, -1)
		b = append(b, ": "...)
	}

	b = append(b, LevelToString(l.Level)...)
	b = append(b, ": "...)
	b = append(b, l.Message...)
	return b
}

//The Square Output Format
//[Level]: Message
func SqrFormat(prefix string, l *LogLine, f int) []byte {
	var b []byte
	b = append(b, prefix...)

	if f&(log.Ldate|log.Ltime|log.Lmicroseconds) != 0 {
		if f&log.Ldate != 0 {
			year, month, day := l.Timestamp.Date()
			itoa(&b, year, 4)
			b = append(b, '/')
			itoa(&b, int(month), 2)
			b = append(b, '/')
			itoa(&b, day, 2)
			b = append(b, ' ')
		}
		if f&(log.Ltime|log.Lmicroseconds) != 0 {
			hour, min, sec := l.Timestamp.Clock()
			itoa(&b, hour, 2)
			b = append(b, ':')
			itoa(&b, min, 2)
			b = append(b, ':')
			itoa(&b, sec, 2)
			if f&log.Lmicroseconds != 0 {
				b = append(b, '.')
				itoa(&b, l.Timestamp.Nanosecond()/1e3, 6)
			}
			b = append(b, ' ')
		}
	}
	if f&(log.Lshortfile|log.Llongfile) != 0 {
		file := l.File
		if f&log.Lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		b = append(b, file...)
		b = append(b, ':')
		itoa(&b, l.Line, -1)
		b = append(b, ": "...)
	}

	b = append(b, '[')
	b = append(b, LevelToString(l.Level)...)
	b = append(b, "] "...)
	b = append(b, l.Message...)
	return b
}

// The Parser func type allows you to add aditional logging conventions to
// interperate.
type Parser func(string) (Level, string)

// The standard convention parser
// Level: Message
func StdParser(m string) (Level, string) {
	c := strings.Index(m, ":")

	// Our Smallest possible level is Info so if it's less than 3 it's not following the standard convension.
	if c > 3 {
		ls := m[:c]
		r := m[c+1:]
		r = strings.TrimLeft(r, " ")

		l := StringToLevel(ls)
		if l != Undefined {
			return l, r
		}
	}
	return Undefined, m
}

// The square convention parser
// [Level] Message
func SqrParser(m string) (Level, string) {
	s := strings.Index(m, "[")
	e := strings.Index(m, "]")

	if s >= 0 && e >= 0 && e > s {
		ls := m[s+1 : e]
		r := m[e+1:]
		r = strings.TrimLeft(r, " ")

		l := StringToLevel(ls)
		if l != Undefined {
			return l, r
		}
	}
	return Undefined, m
}

// Create a new Logger
func New(o io.Writer, p string, f int) *Logger {
	l := &Logger{
		output:    o,
		prefix:    p,
		flag:      f,
		filter:    StdFilter,
		formatter: StdFormat,
		parsers:   []Parser{StdParser},
	}

	return l
}

// set the filter function on the logger.
func (l *Logger) SetFilterFunc(f func(*LogLine) bool) {
	l.filter = f
}

// Set the Filter Function on the default logger.
func SetFilterFunc(f func(*LogLine) bool) {
	std.SetFilterFunc(f)
}

// Get the Current Filter Function
func (l *Logger) FilterFunc() func(*LogLine) bool {
	return l.filter
}

// Get the standard filter function
func FilterFunc() func(*LogLine) bool {
	return std.FilterFunc()
}

// Returns the output flags of the std logger.
func Flags() int {
	return std.Flags()
}

// Returns the output flags of the logger
func (l *Logger) Flags() int {
	return l.flag
}

// Sets the output flag on the default logger.
func SetFlags(f int) {
	std.SetFlags(f)
}

// Sets the default flag on the logger.
func (l *Logger) SetFlags(f int) {
	l.flag = f
}

// Returnes the output io.writer of the std logger.
func Output() io.Writer {
	return std.Output()
}

// Returns the output io.writer.
func (l *Logger) Output() io.Writer {
	return l.output
}

// Sets the output io.writer of the std logger.
func SetOutput(o io.Writer) {
	std.SetOutput(o)
}

// Sets the output io.writer of the logger.
func (l *Logger) SetOutput(o io.Writer) {
	l.output = o
}

// Returns the formatter function of the std logger.
func Formatter() Format {
	return std.Formatter()
}

// Returns the formatter function of the logger.
func (l *Logger) Formatter() Format {
	return l.formatter
}

// Sets the output formatter of the std logger.
func SetFormatter(f Format) {
	std.SetFormatter(f)
}

// Sets the output formatter of the logger.
func (l *Logger) SetFormatter(f Format) {
	l.formatter = f
}

// Returns the prefix of the std logger.
func Prefix() string {
	return std.Prefix()
}

// Returns the prefix of the logger.
func (l *Logger) Prefix() string {
	return l.prefix
}

// Sets the prefix of the std logger.
func SetPrefix(p string) {
	std.SetPrefix(p)
}

// Sets the prefix of the logger.
func (l *Logger) SetPrefix(p string) {
	l.prefix = p
}

// Returns the current parsers in use by the std logger.
func Parsers() []Parser {
	return std.Parsers()
}

// Returns the current parsers in use by the logger.
func (l *Logger) Parsers() []Parser {
	return l.parsers
}

// Sets the parsers used by the std logger.
func SetParsers(p []Parser) {
	std.SetParsers(p)
}

// Sets the parsers used by the logger.
func (l *Logger) SetParsers(p []Parser) {
	l.parsers = p
}

// Logger used to capture logging output prior to filtering/output.
type Logger struct {
	flag   int
	output io.Writer
	prefix string

	filter    func(*LogLine) bool
	formatter Format
	parsers   []Parser
}

// The structure representing the parsed log message.
type LogLine struct {
	Timestamp time.Time
	File      string
	Line      int
	Message   string

	Level Level
}

func (l *LogLine) Equal(ll LogLine) bool {
	if l.File == ll.File &&
		l.Level == ll.Level &&
		l.Line == ll.Line &&
		l.Message == ll.Message &&
		l.Timestamp.Equal(ll.Timestamp) {
		return true
	}
	return false
}

// Implement the io.Writer to capture the message being written to log.
func (l *Logger) Write(p []byte) (int, error) {
	log := StringToLogLine(string(p))

	for _, p := range l.parsers {
		lvl, msg := p(log.Message)
		if lvl != Undefined {
			log.Level = lvl
			log.Message = msg
			break
		}
	}

	if l.filter == nil || l.filter(&log) {
		p = l.formatter(l.prefix, &log, l.flag)
		return l.output.Write([]byte(p))
	} else {
		return 0, io.EOF
	}
}

// Cheap integer to fixed-width decimal ASCII.  Give a negative width to avoid zero-padding.
// Knows the buffer has capacity.
func itoa(buf *[]byte, i int, wid int) {
	var u uint = uint(i)
	if u == 0 && wid <= 1 {
		*buf = append(*buf, '0')
		return
	}

	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; u > 0 || wid > 0; u /= 10 {
		bp--
		wid--
		b[bp] = byte(u%10) + '0'
	}
	*buf = append(*buf, b[bp:]...)
}
