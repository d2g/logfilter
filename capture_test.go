package logfilter_test

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"

	"github.com/d2g/logfilter"
	"github.com/d2g/logfilter/dummy"
)

func TestStringToLevel(t *testing.T) {

	//String to Level Function
	if logfilter.Undefined != logfilter.StringToLevel("skjcksakj") ||
		logfilter.Trace != logfilter.StringToLevel("trace") ||
		logfilter.Debug != logfilter.StringToLevel("debug") ||
		logfilter.Info != logfilter.StringToLevel("INFO") ||
		logfilter.Warning != logfilter.StringToLevel("Warning") ||
		logfilter.Error != logfilter.StringToLevel("error") ||
		logfilter.Fatal != logfilter.StringToLevel("Fatal") ||
		logfilter.Off != logfilter.StringToLevel("Off") {
		t.Errorf("Error Converting String to Log Level")
	}

}

func TestLevelToString(t *testing.T) {
	//Level to String Function
	if "Undefined" != logfilter.LevelToString(logfilter.Undefined) ||
		"Trace" != logfilter.LevelToString(logfilter.Trace) ||
		"Debug" != logfilter.LevelToString(logfilter.Debug) ||
		"Info" != logfilter.LevelToString(logfilter.Info) ||
		"Warning" != logfilter.LevelToString(logfilter.Warning) ||
		"Error" != logfilter.LevelToString(logfilter.Error) ||
		"Fatal" != logfilter.LevelToString(logfilter.Fatal) ||
		"Off" != logfilter.LevelToString(logfilter.Off) {
		t.Errorf("Error Converting Log Level to String")
	}
}

func TestStdParser(t *testing.T) {
	//Std Parser
	level, message := logfilter.StdParser("Debug: Message")
	if level != logfilter.Debug || message != "Message" {
		t.Errorf("Standard Parsing Error: Expected %v Message got %v %s", logfilter.Debug, level, message)
	}

	level, message = logfilter.StdParser("Message")
	if level != logfilter.Undefined || message != "Message" {
		t.Errorf("Standard Parsing Error: Expected %v Message got %v %s", logfilter.Undefined, level, message)
	}
}

func TestSqrParser(t *testing.T) {
	//Sqr Parser
	level, message := logfilter.SqrParser("[Debug] Message")
	if level != logfilter.Debug || message != "Message" {
		t.Errorf("Square Parsing Error: Expected %v Message got %v %s", logfilter.Debug, level, message)
	}

	level, message = logfilter.SqrParser("Message")
	if level != logfilter.Undefined || message != "Message" {
		t.Errorf("Square Parsing Error: Expected %v Message got %v %s", logfilter.Undefined, level, message)
	}
}

func TestStringToLogLine(t *testing.T) {
	//Our Standard Message should look like
	l := logfilter.StringToLogLine("2009/01/23 01:23:23.123123 /a/b/c/d.go:23: debug: message")

	if l.Level != logfilter.Undefined &&
		l.File != "/a/b/c/d.go" &&
		l.Line != 23 &&
		l.Message != "debug: message" &&
		l.Timestamp.Format("2006/01/02 15:04:05.999999") != "2009/01/23 01:23:23.123123" {
		t.Errorf("Expected 2009/01/23 01:23:23.123123 /a/b/c/d.go:23: debug: message actual %s %s:%d: %s", l.Timestamp.Format("2006/01/02 15:04:05.999999"), l.File, l.Line, l.Message)
	}
}

func TestStdFormat(t *testing.T) {
	l := logfilter.LogLine{
		File:    "/a/b/c/d.go",
		Line:    23,
		Message: "message",
		Level:   logfilter.Debug,
	}
	var err error
	l.Timestamp, err = time.Parse("2006/01/02 15:04:05.999999", "2009/01/23 01:23:23.123123")
	if err != nil {
		t.Errorf("Error Creating Dummy Time How Strange??")
	}

	b := logfilter.StdFormat("", &l, log.Ldate|log.Lmicroseconds|log.Lshortfile)
	e := "2009/01/23 01:23:23.123123 d.go:23: Debug: message"
	if !bytes.Equal([]byte(e), b) {
		t.Errorf("Standard Format Expected \"%s\" Actual \"%s\"", e, string(b))
	}
}

func TestSqrFormat(t *testing.T) {
	l := logfilter.LogLine{
		File:    "/a/b/c/d.go",
		Line:    23,
		Message: "message",
		Level:   logfilter.Debug,
	}
	var err error
	l.Timestamp, err = time.Parse("2006/01/02 15:04:05.999999", "2009/01/23 01:23:23.123123")
	if err != nil {
		t.Errorf("Error Creating Dummy Time How Strange??")
	}

	b := logfilter.SqrFormat("", &l, log.Ldate|log.Lmicroseconds|log.Lshortfile)
	e := "2009/01/23 01:23:23.123123 d.go:23: [Debug] message"
	if !bytes.Equal([]byte(e), b) {
		t.Errorf("Standard Format Expected \"%s\" Actual \"%s\"", e, string(b))
	}
}

func TestNew(t *testing.T) {
	p := "prefix"
	l := logfilter.New(os.Stdout, p, log.Ldate)

	if l.Output() != os.Stdout {
		t.Errorf("Invalid Output expected %#v, actual %#v", os.Stdout, l.Output())
	}

	if l.Prefix() != p {
		t.Errorf("Invalid Prefix expected %s, actual %s", p, l.Prefix())
	}

	if l.Flags() != log.Ldate {
		t.Errorf("Invalid Flags expected %d, actual %d", log.Ldate, l.Flags())
	}
}

func TestStd(t *testing.T) {

	//Prefix
	p := "Test Prefix"
	logfilter.SetPrefix(p)
	if logfilter.Prefix() != p {
		t.Errorf("Std Prefix expected %s, actual %s", p, logfilter.Prefix())
	}
	logfilter.SetPrefix("")

	//Parsers
	logfilter.SetParsers([]logfilter.Parser{logfilter.SqrParser, logfilter.StdParser})
	ps := logfilter.Parsers()
	if len(ps) != 2 {
		t.Errorf("Error setting parser expected %d, actual %d", 2, len(ps))
	}

	//Flags
	logfilter.SetFlags(log.Lmicroseconds)
	f := logfilter.Flags()
	if f != log.Lmicroseconds {
		t.Errorf("Error setting flags expected %d, actual %d", log.Lmicroseconds, f)
	}
	logfilter.SetFlags(log.LstdFlags)

	//Output
	logfilter.SetOutput(os.Stdout)
	if logfilter.Output() != os.Stdout {
		t.Errorf("Std Output invalid expected %#v, actual %#v", os.Stdout, logfilter.Output())
	}
	logfilter.SetOutput(os.Stderr)

}

func TestFilterFunc(t *testing.T) {
	f := func(l *logfilter.LogLine) bool {
		if l.Message == "true" {
			return true
		}
		return false
	}

	logfilter.SetFilterFunc(f)
	l := logfilter.LogLine{
		Message: "true",
	}

	b := logfilter.FilterFunc()(&l)
	if !b {
		t.Errorf("Error setting filter function expected %t, actual %t", true, b)
	}

	l.Message = "false"
	b = logfilter.FilterFunc()(&l)
	if b {
		t.Errorf("Error setting filter function expected %t, actual %t", false, b)
	}

	logfilter.SetFilterFunc(logfilter.StdFilter)
}

func TestFormatter(t *testing.T) {
	f := func(s string, l *logfilter.LogLine, i int) []byte {
		return []byte("TEST")
	}

	cf := logfilter.Formatter()

	logfilter.SetFormatter(f)
	b := logfilter.Formatter()("", &logfilter.LogLine{}, 0)
	if !bytes.Equal(b, []byte("TEST")) {
		t.Errorf("Error setting formatter function expected %s, actual %s", "TEST", string(b))
	}

	logfilter.SetFormatter(cf)
}

func TestWrite(t *testing.T) {
	var b bytes.Buffer

	logfilter.SetOutput(&b)
	logfilter.SetFlags(0)
	log.Println("Trace: Message")

	if string(b.Bytes()) != "Trace: Message\n" {
		t.Errorf("Mismatch Sent:\"%s\" Received:\"%s\"\n", "Trace: Message", string(b.Bytes()))
	}

	logfilter.SetOutput(os.Stderr)
	logfilter.SetFlags(log.LstdFlags)
}

func TestFilter(t *testing.T) {
	var b bytes.Buffer

	logfilter.SetOutput(&b)
	logfilter.SetFlags(0)
	logfilter.Exclude("github.com/d2g/logfilter").When(logfilter.Trace)

	//Trace Message Have Been Excluded
	log.Println("Trace: Message")

	//But We should Still Get Debug Messages
	log.Println("Debug: Message")

	if string(b.Bytes()) != "Debug: Message\n" {
		t.Errorf("Mismatch Sent:\"%s\" Received:\"%s\"\n", "Debug: Message", string(b.Bytes()))
	}

	//Reset the filter for the next test.
	logfilter.StdFilterReset()
}

func TestFilterDefault(t *testing.T) {
	var b bytes.Buffer

	logfilter.SetOutput(&b)
	logfilter.SetFlags(0)
	logfilter.Default(logfilter.Error)

	//Trace Message Have Been Excluded
	log.Println("Warning: Message")

	//But We should Still Get Debug Messages
	log.Println("Error: Message")

	if string(b.Bytes()) != "Error: Message\n" {
		t.Errorf("Mismatch Sent:\"%s\" Received:\"%s\"\n", "Error: Message", string(b.Bytes()))
	}
	logfilter.Default(logfilter.Undefined)

	//Reset the filter for the next test.
	logfilter.StdFilterReset()
}

func TestFilterExcludeInclude(t *testing.T) {
	var b bytes.Buffer

	logfilter.SetOutput(&b)
	logfilter.SetFlags(0)
	logfilter.Exclude("github.com/d2g").When(logfilter.Fatal)
	logfilter.Include("github.com/d2g/logfilter").When(logfilter.Debug)
	logfilter.Include("github.com/d2g/logfilter/dummy").When(logfilter.Debug)

	//Trace Message Have Been Excluded
	log.Println("Trace: Message")

	//But We should Still Get Debug Messages
	log.Println("Debug: Message")

	if string(b.Bytes()) != "Debug: Message\n" {
		t.Errorf("Mismatch Sent:\"%s\" Received:\"%s\"\n", "Debug: Message", string(b.Bytes()))
	}

	//If we update an existing filter
	logfilter.Exclude("github.com/d2g/logfilter").When(logfilter.Debug)

	b.Reset()

	log.Println("Debug: Message")

	if string(b.Bytes()) != "" {
		t.Errorf("Mismatch Expected:\"%s\" Actual:\"%s\"\n", "", string(b.Bytes()))
	}

	//Reset the filter for the next test.
	logfilter.StdFilterReset()
}

// Example from documentation
func ExampleLogfilterDocumantation() {
	//Remove date time to make testing simpler.
	logfilter.SetFlags(log.Lshortfile)

	//Set the Output to stdout for the example test.
	logfilter.SetOutput(os.Stdout)

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

	//Reset the filter for the next test.
	logfilter.StdFilterReset()
}

func ExampleLogfilter() {

	// Set the log output to just the short filename so the output doesn't contain the date time which changes each test.
	logfilter.SetFlags(log.Lshortfile)

	//Output to std out for testing
	logfilter.SetOutput(os.Stdout)

	// Don't output any messages Fatal or Lower from the package or subpackages github.com/d2g/logfilter.
	logfilter.Exclude("github.com/d2g/logfilter").When(logfilter.Fatal)

	// Output any messages Warning or Above from the package or subpackages github.com/d2g/logfilter/dummy.
	logfilter.Include("github.com/d2g/logfilter/dummy").When(logfilter.Warning)

	dummy.Fatal()       // Create a dummy message.
	dummy.Error()       // Create a dummy Error message.
	dummy.Warning()     // Create a dummy Warning message.
	dummy.Info()        // Create a dummy Info message that should be ignored.
	dummy.Debug()       // Create a dummy degub message that should be ignored.
	dummy.Trace()       // Create a dummy trace message that should be ignored.
	dummy.Unformatted() //Unformatted messages appear as trace messages that should be ignored.

	//Output:
	//dummy.go:29: Fatal: This is a Fatal message
	//dummy.go:25: Error: This is a Error message
	//dummy.go:21: Warning: This is a Warning message

	//Reset the filter for the next test.
	logfilter.StdFilterReset()
}
