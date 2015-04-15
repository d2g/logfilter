# logfilter [![GoDoc](https://godoc.org/github.com/d2g/logfilter?status.svg)](http://godoc.org/github.com/d2g/logfilter) [![Coverage Status](https://coveralls.io/repos/d2g/logfilter/badge.png?branch=HEAD)](https://coveralls.io/r/d2g/logfilter?branch=HEAD) [![Go Report Card](http://goreportcard.com/badge/d2g/logfilter)](http://goreportcard.com/report/d2g/logfilter) [![Codeship Status for d2g/logfilter](https://codeship.io/projects/a80df9b0-3db4-0132-591a-3a26f38803db/status)](https://codeship.io/projects/43342)
=========

## Logging By Convention Rather Than Configuration

Package logfilter implements a filter on the standard log package.

Like most alternative logging packages using this package support the semi
standard logging levels like:


	Trace
	Debug
	Info
	Warning
	Error
	Fatal

Unlike other packages log filter doesn't require the package to specifically
import an additional package over the standard logging package. It does however
require them to follow a convention.

The default convention is:

`log.Println("<Level>:<Message>")`

i.e.

`log.Println("Trace:Example Message")`
is a trace level message.

When logging at trace level it's usual to be very verbose which results in a
large amount of output. Filtering this information down can result in it being easier 
to debug issues and system requirements while producing logging information(i.e. Disk IO).

Logfilter allows you to filter this output based on the package/file location
and the log level type.

For example in your live application you may only want to output warning
messages:


	import(
		"log"
		"github.com/d2g/logfilter"
		"github.com/d2g/logfilter/dummy"
	)
	
	func main() {
		// Change the default filter to warning and above.
		logfilter.Default(logfilter.Warning)
	
		// I want to debug an issue in a particular package so want logging from that package.
		logfilter.Include("github.com/d2g/logfilter/dummy")
	
		// However at this stage I want only Info and above.
		logfilter.Include("github.com/d2g/logfilter/dummy").When(logfilter.Info)
	
		// Now only log level Warning and above will be written
		// Except for github.com/d2g/dummy which wil have Info and above.
		log.Println("Debug: Not Displayed")
		dummy.Debug()
		dummy.Info()
	
		//Output:
		//dummy.go:17: Info: This is a Info message
	}






## func Default
``` go
func Default(lvl Level)
```
Default sets the logging level the is output by all packages.


## func Exclude
``` go
func Exclude(packagenames ...string) filters
```
Exclude adds filter object(s) to the standard filter and returns pointers to
the newly created object to allow you to set the required level. As default it
sets the level to off (i.e. highest).


## func FilterFunc
``` go
func FilterFunc() func(*LogLine) bool
```
FilterFunc returns the standard filter function


## func Flags
``` go
func Flags() int
```
Flags returns the output flags of the std logger.


## func Include
``` go
func Include(packagenames ...string) filters
```
Include adds filter object(s) to the standard filter and returns pointers to
the newly created object to allow you to set the required level. As default it
sets the level to undefined (i.e. lowest).


## func LevelToString
``` go
func LevelToString(l Level) string
```
LevelToString converts a Level (i.e. Error) to the corresponding string version (i.e. "Error").


## func Output
``` go
func Output() io.Writer
```
Output returns the output io.writer of the std logger.


## func Parsers
``` go
func Parsers() []Parser
```
Parsers returns the current parsers in use by the std logger.


## func Prefix
``` go
func Prefix() string
```
Prefix returns the prefix of the std logger.


## func SetFilterFunc
``` go
func SetFilterFunc(f func(*LogLine) bool)
```
SetFilterFunc Set the Filter Function on the default logger.


## func SetFlags
``` go
func SetFlags(f int)
```
SetFlags sets the output flags on the default logger.
The flags match those used by the standard log package.


## func SetFormatter
``` go
func SetFormatter(f Format)
```
SetFormatter sets the output formatter of the std logger.


## func SetOutput
``` go
func SetOutput(o io.Writer)
```
SetOutput sets the output io.writer of the std logger.


## func SetParsers
``` go
func SetParsers(p []Parser)
```
SetParsers sets the parsers used by the std logger.


## func SetPrefix
``` go
func SetPrefix(p string)
```
SetPrefix sets the prefix of the std logger.


## func SqrFormat
``` go
func SqrFormat(prefix string, l *LogLine, f int) []byte
```
SqrFormat generates the square output format
[Level]: Message


## func StdFilter
``` go
func StdFilter(l *LogLine) bool
```
StdFilter is the default implementation used by logger for filtering.
Returns true if the line is written out.


## func StdFilterReset
``` go
func StdFilterReset()
```
StdFilterReset resets the filters being applied, which is useful for testing.


## func StdFormat
``` go
func StdFormat(prefix string, l *LogLine, f int) []byte
```
StdFormat generates the standard output format
Level: Message



## type Format
``` go
type Format func(string, *LogLine, int) []byte
```
Format is the func type required for formatting the
output of log messages.It allows messages logged in one packages as
`Level: Message` to be output as `[Level] Message`.









### func Formatter
``` go
func Formatter() Format
```
Formatter returns the formatter function of the std logger.




## type Level
``` go
type Level int
```
Level represents a logging level.



``` go
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
```
Standard(ish) Logging Levels.







### func SqrParser
``` go
func SqrParser(m string) (Level, string)
```
SqrParser square convention parser
[Level] Message


### func StdParser
``` go
func StdParser(m string) (Level, string)
```
StdParser is the standard convention parser
Level: Message


### func StringToLevel
``` go
func StringToLevel(sl string) Level
```
StringToLevel converts a string log level (i.e. "Error") to the corresponding Level (i.e. Error).




## type LogLine
``` go
type LogLine struct {
    Timestamp time.Time
    File      string
    Line      int
    Message   string

    Level Level
}
```
LogLine struct representing the parsed log message.









### func StringToLogLine
``` go
func StringToLogLine(m string) LogLine
```
StringToLogLine converts the default log sting sent to the LogLine type.




## type Logger
``` go
type Logger struct {
    // contains filtered or unexported fields
}
```
Logger used to capture logging output prior to filtering/output.









### func New
``` go
func New(o io.Writer, p string, f int) *Logger
```
New creates a new Logger




### func (\*Logger) FilterFunc
``` go
func (l *Logger) FilterFunc() func(*LogLine) bool
```
FilterFunc returns the Current Filter Function



### func (\*Logger) Flags
``` go
func (l *Logger) Flags() int
```
Flags returns the output flags of the logger



### func (\*Logger) Formatter
``` go
func (l *Logger) Formatter() Format
```
Formatter returns the formatter function of the logger.



### func (\*Logger) Output
``` go
func (l *Logger) Output() io.Writer
```
Output returns the output io.writer of the logger.



### func (\*Logger) Parsers
``` go
func (l *Logger) Parsers() []Parser
```
Parsers returns the current parsers in use by the logger.



### func (\*Logger) Prefix
``` go
func (l *Logger) Prefix() string
```
Prefix returns the prefix of the logger.



### func (\*Logger) SetFilterFunc
``` go
func (l *Logger) SetFilterFunc(f func(*LogLine) bool)
```
SetFilterFunc set the filter function on the logger.



### func (\*Logger) SetFlags
``` go
func (l *Logger) SetFlags(f int)
```
SetFlags sets the default flag on the logger.



### func (\*Logger) SetFormatter
``` go
func (l *Logger) SetFormatter(f Format)
```
SetFormatter sets the output formatter of the logger.



### func (\*Logger) SetOutput
``` go
func (l *Logger) SetOutput(o io.Writer)
```
SetOutput sets the output io.writer of the logger.



### func (\*Logger) SetParsers
``` go
func (l *Logger) SetParsers(p []Parser)
```
SetParsers sets the parsers used by the logger.



### func (\*Logger) SetPrefix
``` go
func (l *Logger) SetPrefix(p string)
```
SetPrefix sets the prefix of the logger.



### func (\*Logger) Write
``` go
func (l *Logger) Write(p []byte) (int, error)
```
Write is the implement the io.Writer to capture the message being written to log.



## type Parser
``` go
type Parser func(string) (Level, string)
```
The Parser func type allows you to add additional logging conventions to
interpret different conventions.

















- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
