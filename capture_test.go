package logfilter_test

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"
	"text/template"
	"time"

	"github.com/d2g/logfilter"
	"github.com/d2g/logfilter/dummy"
)

// capture the content written by f to os.Stderr
func captureStandardErr(f func()) (out string, err error) {
	// Capture stdout.
	stderr := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		return
	}

	os.Stderr = w
	defer func() {
		os.Stderr = stderr
	}()

	capturedOutput := make(chan string)
	errorChannel := make(chan error)

	go func() {
		buf := new(bytes.Buffer)
		_, err := io.Copy(buf, r)
		r.Close()
		if err != nil {
			errorChannel <- err
		}

		capturedOutput <- buf.String()
	}()

	//FUNCTION
	f()
	err = w.Close()
	if err != nil {
		return
	}

	select {
	case out = <-capturedOutput:
	case err = <-errorChannel:
	}
	return
}

func TestParseMessage(test *testing.T) {

	example := logfilter.Capture{}

	reference := logfilter.Line{
		Message: "Message",
		Level:   logfilter.DEBUG,
	}

	line := example.Parse("Debug: Message")

	if !line.Equal(reference) {
		test.Errorf("Parsing Error")
	}

	line = example.Parse("Warning: Message")
	reference.Level = logfilter.WARNING

	if !line.Equal(reference) {
		test.Errorf("Parsing Error")
	}

	line = example.Parse("Error: Message")
	reference.Level = logfilter.ERROR

	if !line.Equal(reference) {
		test.Errorf("Parsing Error")
	}

	line = example.Parse("Fatal: Message")
	reference.Level = logfilter.FATAL

	if !line.Equal(reference) {
		test.Errorf("Parsing Error")
	}
}

func TestParseMessageTime(test *testing.T) {
	example := logfilter.Capture{}
	example.Flags = log.Ltime

	line := example.Parse("14:45:45 INFO: Message")
	t, _ := time.Parse("15:04:05", "14:45:45")

	parsedline := logfilter.Line{
		Timestamp: t,
		Message:   "Message",
		Level:     logfilter.INFO,
	}

	if !line.Equal(parsedline) {
		test.Errorf("Parsing Error Parsed Message:%v\n", line)
	}
}

func TestParseMessageTimeShortFilename(test *testing.T) {
	example := logfilter.Capture{
		Flags: (log.Ltime | log.Lshortfile),
	}

	line := example.Parse("14:45:45 main.go:156: Trace: Message")
	t, _ := time.Parse("15:04:05", "14:45:45")

	parsedline := logfilter.Line{
		Timestamp:   t,
		FileAndLine: "main.go:156",
		Message:     "Message",
		Level:       logfilter.TRACE,
	}

	if !line.Equal(parsedline) {
		test.Errorf("Parsing Error Parsed Message:%v\n", line)
	}
}

func TestParseMessageDateTimeLongFilename(test *testing.T) {

	example := logfilter.Capture{
		Flags: (log.Ldate | log.Ltime | log.Llongfile),
	}

	line := example.Parse("2014/10/03 14:45:45 C:/Go/src/github.com/d2g/logfilter/main.go:155: Trace: Message")
	t, _ := time.Parse("2006/01/02 15:04:05", "2014/10/03 14:45:45")

	parsedline := logfilter.Line{
		Timestamp:   t,
		FileAndLine: "C:/Go/src/github.com/d2g/logfilter/main.go:155",
		Message:     "Message",
		Level:       logfilter.TRACE,
	}

	if !line.Equal(parsedline) {
		test.Errorf("Parsing Error Parsed Message:%v\n", line)
	}

}

func TestWriter(test *testing.T) {

	var b bytes.Buffer

	log.SetFlags(0)

	log.SetOutput(&logfilter.Capture{
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

	log.SetOutput(&logfilter.Capture{
		Flags:  log.Flags(),
		Output: &b,
		Filters: []logfilter.Filter{
			logfilter.Filter{
				Mode:     logfilter.EXCLUDE,
				Filename: "github.com/d2g/logfilter",
				Level:    logfilter.TRACE,
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

	log.SetOutput(&logfilter.Capture{
		Flags:  log.Flags(),
		Output: &b,
		Filters: []logfilter.Filter{
			logfilter.Filter{
				Mode:     logfilter.EXCLUDE,
				Filename: "github.com/d2g/",
				Level:    logfilter.FATAL,
			},
			logfilter.Filter{
				Mode:     logfilter.INCLUDE,
				Filename: "github.com/d2g/logfilter",
				Level:    logfilter.DEBUG,
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
	e1 := logfilter.Line{}
	e2 := logfilter.Line{}

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

	e2.Level = logfilter.INFO

	if e1.Equal(e2) {
		test.Errorf("Equals Failed")
	}

	e2.Level = 0

	if !e1.Equal(e2) {
		test.Errorf("Equals Failed")
	}

}

func TestWriteToOSStderr(t *testing.T) {
	example := logfilter.Capture{}
	message := "Debug: Message\n"

	output, err := captureStandardErr(func() {
		n, err := example.Write([]byte(message))
		if err != nil {
			t.Errorf("Write Failed")
		}

		if n != len([]byte("Debug: Message\n")) {
			t.Errorf("Write Failed, Incorrect Length")
		}
	})

	if err != nil {
		t.Errorf("Error Capturing StandardErr %v\n", err)
	}

	if output != message {
		t.Errorf("StandardErr expect %s got %s\n", message, output)
	}
}

//Examples
func ExampleCapture() {

	// Set our filters, log.Llongfile is important here or we don't get the filename
	// to filter on.
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	// Setup our log filter.
	log.SetOutput(&logfilter.Capture{
		Flags: log.Flags(),
		Filters: []logfilter.Filter{
			logfilter.Filter{ // Exclude all our messages
				Mode:     logfilter.EXCLUDE,
				Filename: "github.com/d2g/logfilter",
				Level:    logfilter.FATAL,
			},
			logfilter.Filter{ // Include Fatal Messages Only
				Mode:     logfilter.INCLUDE,
				Filename: "github.com/d2g/logfilter/dummy",
				Level:    logfilter.WARNING,
			},
		},
		Output: os.Stdout,
		Formatter: func(l *logfilter.Line) []byte {
			t := template.Must(template.New("format").Parse(`{{.Message}}`))
			b := bytes.NewBuffer([]byte{})
			err := t.Execute(b, *l)
			if err != nil {
				panic("Panic Template Error:" + err.Error())
			}

			return b.Bytes()
		},
	})

	dummy.Fatal()   // Create a dummy message.
	dummy.Error()   // Create a dummy Error message.
	dummy.Warning() // Create a dummy Warning message.
	dummy.Info()    // Create a dummy Info message that should be ignored.
	dummy.Debug()   // Create a dummy degub message that should be ignored.
	dummy.Trace()   // Create a dummy trace message that should be ignored.

	dummy.Unformatted() //Unformatted messages appear as trace messages.

	//Output:
	//This is a Fatal message
	//This is a Error message
	//This is a Warning message
	//Some package that doesn't implement the convention.
}
