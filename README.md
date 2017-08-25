# logfilter [![GoDoc](https://godoc.org/github.com/d2g/logfilter?status.svg)](http://godoc.org/github.com/d2g/logfilter) [![Coverage Status](https://coveralls.io/repos/d2g/logfilter/badge.png?branch=HEAD)](https://coveralls.io/r/d2g/logfilter?branch=HEAD) [![Go Report Card](http://goreportcard.com/badge/d2g/logfilter)](http://goreportcard.com/report/d2g/logfilter) [![Codeship Status for d2g/logfilter](https://codeship.io/projects/a80df9b0-3db4-0132-591a-3a26f38803db/status)](https://codeship.io/projects/43342)
=========

## Logging By Convention Rather Than Configuration
`import "github.com/d2g/logfilter"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
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

log.Println("<Level>:<Message>")

i.e.

log.Println("Trace:Example Message")
is a trace level message.

When logging at trace level it's usual to be very verbose which results in a
large amount of output. Filtering this information down can result in it being
easier to debug issues and system requirements while producing logging
information(i.e. Disk IO).

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
		// Except for github.com/d2g/dummy which will have Info and above.
		log.Println("Debug: Not Displayed")
		dummy.Debug()
		dummy.Info()
	
		//Output:
		//dummy.go:17: Info: This is a Info message
	}

If you've previously used logutils or a square based convention then look at the
example included in example_logutils_test.go




## <a name="pkg-index">Index</a>
* [func Default(lvl Level)](#Default)
* [func Exclude(packagenames ...string) filters](#Exclude)
* [func FilterFunc() func(*LogLine) bool](#FilterFunc)
* [func Flags() int](#Flags)
* [func Include(packagenames ...string) filters](#Include)
* [func LevelToString(l Level) string](#LevelToString)
* [func Output() io.Writer](#Output)
* [func Parsers() []Parser](#Parsers)
* [func Prefix() string](#Prefix)
* [func SetFilterFunc(f func(*LogLine) bool)](#SetFilterFunc)
* [func SetFlags(f int)](#SetFlags)
* [func SetFormatter(f Format)](#SetFormatter)
* [func SetOutput(o io.Writer)](#SetOutput)
* [func SetParsers(p []Parser)](#SetParsers)
* [func SetPrefix(p string)](#SetPrefix)
* [func SqrFormat(prefix string, l *LogLine, f int) []byte](#SqrFormat)
* [func StdFilter(l *LogLine) bool](#StdFilter)
* [func StdFilterReset()](#StdFilterReset)
* [func StdFormat(prefix string, l *LogLine, f int) []byte](#StdFormat)
* [type Format](#Format)
  * [func Formatter() Format](#Formatter)
* [type Level](#Level)
  * [func SqrParser(m string) (Level, string)](#SqrParser)
  * [func StdParser(m string) (Level, string)](#StdParser)
  * [func StringToLevel(sl string) Level](#StringToLevel)
* [type LogLine](#LogLine)
  * [func StringToLogLine(m string) LogLine](#StringToLogLine)
* [type Logger](#Logger)
  * [func New(o io.Writer, p string, f int) *Logger](#New)
  * [func (l *Logger) FilterFunc() func(*LogLine) bool](#Logger.FilterFunc)
  * [func (l *Logger) Flags() int](#Logger.Flags)
  * [func (l *Logger) Formatter() Format](#Logger.Formatter)
  * [func (l *Logger) Output() io.Writer](#Logger.Output)
  * [func (l *Logger) Parsers() []Parser](#Logger.Parsers)
  * [func (l *Logger) Prefix() string](#Logger.Prefix)
  * [func (l *Logger) SetFilterFunc(f func(*LogLine) bool)](#Logger.SetFilterFunc)
  * [func (l *Logger) SetFlags(f int)](#Logger.SetFlags)
  * [func (l *Logger) SetFormatter(f Format)](#Logger.SetFormatter)
  * [func (l *Logger) SetOutput(o io.Writer)](#Logger.SetOutput)
  * [func (l *Logger) SetParsers(p []Parser)](#Logger.SetParsers)
  * [func (l *Logger) SetPrefix(p string)](#Logger.SetPrefix)
  * [func (l *Logger) Write(p []byte) (int, error)](#Logger.Write)
* [type Parser](#Parser)


#### <a name="pkg-files">Package files</a>
[capture.go](/src/github.com/d2g/logfilter/capture.go) [doc.go](/src/github.com/d2g/logfilter/doc.go) [filter.go](/src/github.com/d2g/logfilter/filter.go) [init.go](/src/github.com/d2g/logfilter/init.go) [levels.go](/src/github.com/d2g/logfilter/levels.go) 





## <a name="Default">func</a> [Default](/src/target/filter.go?s=334:357#L8)
``` go
func Default(lvl Level)
```
Default sets the logging level the is output by all packages.



## <a name="Exclude">func</a> [Exclude](/src/target/filter.go?s=1415:1459#L52)
``` go
func Exclude(packagenames ...string) filters
```
Exclude adds filter object(s) to the standard filter and returns pointers to
the newly created object to allow you to set the required level. As default it
sets the level to off (i.e. highest).



## <a name="FilterFunc">func</a> [FilterFunc](/src/target/capture.go?s=4266:4303#L186)
``` go
func FilterFunc() func(*LogLine) bool
```
FilterFunc returns the standard filter function



## <a name="Flags">func</a> [Flags](/src/target/capture.go?s=4386:4402#L191)
``` go
func Flags() int
```
Flags returns the output flags of the std logger.



## <a name="Include">func</a> [Include](/src/target/filter.go?s=816:860#L26)
``` go
func Include(packagenames ...string) filters
```
Include adds filter object(s) to the standard filter and returns pointers to
the newly created object to allow you to set the required level. As default it
sets the level to undefined (i.e. lowest).



## <a name="LevelToString">func</a> [LevelToString](/src/target/levels.go?s=730:764#L36)
``` go
func LevelToString(l Level) string
```
LevelToString converts a Level (i.e. Error) to the corresponding string version (i.e. "Error").



## <a name="Output">func</a> [Output](/src/target/capture.go?s=4836:4859#L212)
``` go
func Output() io.Writer
```
Output returns the output io.writer of the std logger.



## <a name="Parsers">func</a> [Parsers](/src/target/capture.go?s=6133:6156#L272)
``` go
func Parsers() []Parser
```
Parsers returns the current parsers in use by the std logger.



## <a name="Prefix">func</a> [Prefix](/src/target/capture.go?s=5730:5750#L252)
``` go
func Prefix() string
```
Prefix returns the prefix of the std logger.



## <a name="SetFilterFunc">func</a> [SetFilterFunc](/src/target/capture.go?s=4026:4067#L176)
``` go
func SetFilterFunc(f func(*LogLine) bool)
```
SetFilterFunc Set the Filter Function on the default logger.



## <a name="SetFlags">func</a> [SetFlags](/src/target/capture.go?s=4638:4658#L202)
``` go
func SetFlags(f int)
```
SetFlags sets the output flags on the default logger.
The flags match those used by the standard log package.



## <a name="SetFormatter">func</a> [SetFormatter](/src/target/capture.go?s=5511:5538#L242)
``` go
func SetFormatter(f Format)
```
SetFormatter sets the output formatter of the std logger.



## <a name="SetOutput">func</a> [SetOutput](/src/target/capture.go?s=5054:5081#L222)
``` go
func SetOutput(o io.Writer)
```
SetOutput sets the output io.writer of the std logger.



## <a name="SetParsers">func</a> [SetParsers](/src/target/capture.go?s=6357:6384#L282)
``` go
func SetParsers(p []Parser)
```
SetParsers sets the parsers used by the std logger.



## <a name="SetPrefix">func</a> [SetPrefix](/src/target/capture.go?s=5922:5946#L262)
``` go
func SetPrefix(p string)
```
SetPrefix sets the prefix of the std logger.



## <a name="SqrFormat">func</a> [SqrFormat](/src/target/capture.go?s=1546:1601#L61)
``` go
func SqrFormat(prefix string, l *LogLine, f int) []byte
```
SqrFormat generates the square output format
[Level]: Message



## <a name="StdFilter">func</a> [StdFilter](/src/target/filter.go?s=2256:2287#L95)
``` go
func StdFilter(l *LogLine) bool
```
StdFilter is the default implementation used by logger for filtering.
Returns true if the line is written out.



## <a name="StdFilterReset">func</a> [StdFilterReset](/src/target/filter.go?s=2005:2026#L83)
``` go
func StdFilterReset()
```
StdFilterReset resets the filters being applied, which is useful for testing.



## <a name="StdFormat">func</a> [StdFormat](/src/target/capture.go?s=355:410#L7)
``` go
func StdFormat(prefix string, l *LogLine, f int) []byte
```
StdFormat generates the standard output format
Level: Message




## <a name="Format">type</a> [Format](/src/target/capture.go?s=241:287#L3)
``` go
type Format func(string, *LogLine, int) []byte
```
Format is the func type required for formatting the
output of log messages.It allows messages logged in one packages as
`Level: Message` to be output as `[Level] Message`.







### <a name="Formatter">func</a> [Formatter](/src/target/capture.go?s=5279:5302#L232)
``` go
func Formatter() Format
```
Formatter returns the formatter function of the std logger.





## <a name="Level">type</a> [Level](/src/target/levels.go?s=98:112#L1)
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







### <a name="SqrParser">func</a> [SqrParser](/src/target/capture.go?s=3306:3346#L139)
``` go
func SqrParser(m string) (Level, string)
```
SqrParser square convention parser
[Level] Message


### <a name="StdParser">func</a> [StdParser](/src/target/capture.go?s=2906:2946#L120)
``` go
func StdParser(m string) (Level, string)
```
StdParser is the standard convention parser
Level: Message


### <a name="StringToLevel">func</a> [StringToLevel](/src/target/levels.go?s=331:366#L15)
``` go
func StringToLevel(sl string) Level
```
StringToLevel converts a string log level (i.e. "Error") to the corresponding Level (i.e. Error).





## <a name="LogLine">type</a> [LogLine](/src/target/capture.go?s=6781:6890#L303)
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







### <a name="StringToLogLine">func</a> [StringToLogLine](/src/target/levels.go?s=1089:1127#L57)
``` go
func StringToLogLine(m string) LogLine
```
StringToLogLine converts the default log sting sent to the LogLine type.





## <a name="Logger">type</a> [Logger](/src/target/capture.go?s=6587:6724#L292)
``` go
type Logger struct {
    // contains filtered or unexported fields
}
```
Logger used to capture logging output prior to filtering/output.







### <a name="New">func</a> [New](/src/target/capture.go?s=3624:3670#L157)
``` go
func New(o io.Writer, p string, f int) *Logger
```
New creates a new Logger





### <a name="Logger.FilterFunc">func</a> (\*Logger) [FilterFunc](/src/target/capture.go?s=4144:4193#L181)
``` go
func (l *Logger) FilterFunc() func(*LogLine) bool
```
FilterFunc returns the Current Filter Function




### <a name="Logger.Flags">func</a> (\*Logger) [Flags](/src/target/capture.go?s=4475:4503#L196)
``` go
func (l *Logger) Flags() int
```
Flags returns the output flags of the logger




### <a name="Logger.Formatter">func</a> (\*Logger) [Formatter](/src/target/capture.go?s=5390:5425#L237)
``` go
func (l *Logger) Formatter() Format
```
Formatter returns the formatter function of the logger.




### <a name="Logger.Output">func</a> (\*Logger) [Output](/src/target/capture.go?s=4939:4974#L217)
``` go
func (l *Logger) Output() io.Writer
```
Output returns the output io.writer of the logger.




### <a name="Logger.Parsers">func</a> (\*Logger) [Parsers](/src/target/capture.go?s=6244:6279#L277)
``` go
func (l *Logger) Parsers() []Parser
```
Parsers returns the current parsers in use by the logger.




### <a name="Logger.Prefix">func</a> (\*Logger) [Prefix](/src/target/capture.go?s=5820:5852#L257)
``` go
func (l *Logger) Prefix() string
```
Prefix returns the prefix of the logger.




### <a name="Logger.SetFilterFunc">func</a> (\*Logger) [SetFilterFunc](/src/target/capture.go?s=3890:3943#L171)
``` go
func (l *Logger) SetFilterFunc(f func(*LogLine) bool)
```
SetFilterFunc set the filter function on the logger.




### <a name="Logger.SetFlags">func</a> (\*Logger) [SetFlags](/src/target/capture.go?s=4729:4761#L207)
``` go
func (l *Logger) SetFlags(f int)
```
SetFlags sets the default flag on the logger.




### <a name="Logger.SetFormatter">func</a> (\*Logger) [SetFormatter](/src/target/capture.go?s=5621:5660#L247)
``` go
func (l *Logger) SetFormatter(f Format)
```
SetFormatter sets the output formatter of the logger.




### <a name="Logger.SetOutput">func</a> (\*Logger) [SetOutput](/src/target/capture.go?s=5158:5197#L227)
``` go
func (l *Logger) SetOutput(o io.Writer)
```
SetOutput sets the output io.writer of the logger.




### <a name="Logger.SetParsers">func</a> (\*Logger) [SetParsers](/src/target/capture.go?s=6459:6498#L287)
``` go
func (l *Logger) SetParsers(p []Parser)
```
SetParsers sets the parsers used by the logger.




### <a name="Logger.SetPrefix">func</a> (\*Logger) [SetPrefix](/src/target/capture.go?s=6013:6049#L267)
``` go
func (l *Logger) SetPrefix(p string)
```
SetPrefix sets the prefix of the logger.




### <a name="Logger.Write">func</a> (\*Logger) [Write](/src/target/capture.go?s=6977:7022#L313)
``` go
func (l *Logger) Write(p []byte) (int, error)
```
Write is the implement the io.Writer to capture the message being written to log.




## <a name="Parser">type</a> [Parser](/src/target/capture.go?s=2801:2841#L116)
``` go
type Parser func(string) (Level, string)
```
The Parser func type allows you to add additional logging conventions to
interpret different conventions.














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)

