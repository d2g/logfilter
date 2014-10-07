package logfilter

import (
	"bytes"
	"log"
	"os"
	"testing"
	"time"
)

func TestParseMessage(test *testing.T) {

	example := Capture{}

	reference := Line{
		Message: "Message",
		Level:   DEBUG,
	}

	line := example.Parse("Debug: Message")

	if !line.Equal(reference) {
		test.Errorf("Parsing Error")
	}

	line = example.Parse("Warning: Message")
	reference.Level = WARNING

	if !line.Equal(reference) {
		test.Errorf("Parsing Error")
	}

	line = example.Parse("Error: Message")
	reference.Level = ERROR

	if !line.Equal(reference) {
		test.Errorf("Parsing Error")
	}

	line = example.Parse("Fatal: Message")
	reference.Level = FATAL

	if !line.Equal(reference) {
		test.Errorf("Parsing Error")
	}
}

func TestParseMessageTime(test *testing.T) {
	example := Capture{}
	example.Flags = log.Ltime

	line := example.Parse("14:45:45 INFO: Message")
	t, _ := time.Parse("15:04:05", "14:45:45")

	parsedline := Line{
		Timestamp: t,
		Message:   "Message",
		Level:     INFO,
	}

	if !line.Equal(parsedline) {
		test.Errorf("Parsing Error Parsed Message:%v\n", line)
	}
}

func TestParseMessageTimeShortFilename(test *testing.T) {
	example := Capture{
		Flags: (log.Ltime | log.Lshortfile),
	}

	line := example.Parse("14:45:45 main.go:156: Trace: Message")
	t, _ := time.Parse("15:04:05", "14:45:45")

	parsedline := Line{
		Timestamp:   t,
		FileAndLine: "main.go:156",
		Message:     "Message",
		Level:       TRACE,
	}

	if !line.Equal(parsedline) {
		test.Errorf("Parsing Error Parsed Message:%v\n", line)
	}
}

func TestParseMessageDateTimeLongFilename(test *testing.T) {

	example := Capture{
		Flags: (log.Ldate | log.Ltime | log.Llongfile),
	}

	line := example.Parse("2014/10/03 14:45:45 C:/Go/src/github.com/d2g/logfilter/main.go:155: Trace: Message")
	t, _ := time.Parse("2006/01/02 15:04:05", "2014/10/03 14:45:45")

	parsedline := Line{
		Timestamp:   t,
		FileAndLine: "C:/Go/src/github.com/d2g/logfilter/main.go:155",
		Message:     "Message",
		Level:       TRACE,
	}

	if !line.Equal(parsedline) {
		test.Errorf("Parsing Error Parsed Message:%v\n", line)
	}

}

func TestWriter(test *testing.T) {

	var b bytes.Buffer

	log.SetFlags(0)

	log.SetOutput(&Capture{
		Flags:  log.Flags(),
		Output: &b,
	})

	log.Println("Trace: Message")

	if string(b.Bytes()) != "Trace: Message\n" {
		test.Errorf("Mismatch Sent:\"%s\" Received:\"%s\"\n", "Trace: Message", string(b.Bytes()))
	}

}

func TestWriterExclude(test *testing.T) {

	var b bytes.Buffer

	log.SetFlags((log.Ldate | log.Ltime | log.Llongfile))

	log.SetOutput(&Capture{
		Flags:  log.Flags(),
		Output: &b,
		Filters: []Filter{
			Filter{
				Mode:     EXCLUDE,
				Filename: "github.com/d2g/logfilter",
				Level:    TRACE,
			},
		},
	})

	//Trace Message Have Been Excluded
	log.Println("Trace: Message")

	//But We should Still Get Debug Messages
	log.Println("Debug: Message")

	if string(b.Bytes())[len(string(b.Bytes()))-15:] != "Debug: Message\n" {
		test.Errorf("Mismatch Sent:\"%s\" Received:\"%s\"\n", "Debug: Message", string(b.Bytes())[len(string(b.Bytes()))-15:])
	}

}

func TestWriterExcludeInclude(test *testing.T) {

	var b bytes.Buffer

	log.SetFlags((log.Ldate | log.Ltime | log.Llongfile))

	log.SetOutput(&Capture{
		Flags:  log.Flags(),
		Output: &b,
		Filters: []Filter{
			Filter{
				Mode:     EXCLUDE,
				Filename: "github.com/d2g/",
				Level:    FATAL,
			},
			Filter{
				Mode:     INCLUDE,
				Filename: "github.com/d2g/logfilter",
				Level:    DEBUG,
			},
		},
	})

	//Trace Message Have Been Excluded
	log.Println("Trace: Message")

	//But We should Still Get Debug Messages
	log.Println("Debug: Message")

	if string(b.Bytes())[len(string(b.Bytes()))-15:] != "Debug: Message\n" {
		test.Errorf("Mismatch Sent:\"%s\" Received:\"%s\"\n", "Debug: Message", string(b.Bytes())[len(string(b.Bytes()))-15:])
	}

}

func TestLineEquals(test *testing.T) {
	e1 := Line{}
	e2 := Line{}

	t, _ := time.Parse("2006/01/02 15:04:05", "2014/10/03 14:45:45")

	e2.Timestamp = t
	if e1.Equal(e2) {
		test.Errorf("Equals Failed")
	}
	e2.Timestamp = time.Time{}

	if !e1.Equal(e2) {
		test.Errorf("Equals Failed")
	}

	e2.FileAndLine = "Y"

	if e1.Equal(e2) {
		test.Errorf("Equals Failed")
	}

	e2.FileAndLine = ""

	if !e1.Equal(e2) {
		test.Errorf("Equals Failed")
	}

	e2.Message = "Y"

	if e1.Equal(e2) {
		test.Errorf("Equals Failed")
	}

	e2.Message = ""

	if !e1.Equal(e2) {
		test.Errorf("Equals Failed")
	}

	e2.Level = INFO

	if e1.Equal(e2) {
		test.Errorf("Equals Failed")
	}

	e2.Level = 0

	if !e1.Equal(e2) {
		test.Errorf("Equals Failed")
	}

}

func TestWriteToOSStderr(t *testing.T) {
	example := Capture{}

	os.Stderr = os.Stdout

	n, err := example.Write([]byte("Debug: Message\n"))
	if err != nil {
		t.Errorf("Write Failed")
	}

	if n != len([]byte("Debug: Message\n")) {
		t.Errorf("Write Failed, Incorrect Length")
	}

	// Output: Debug: Message
}
