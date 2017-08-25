package logfilter_test

import (
	"log"
	"os"

	"github.com/d2g/logfilter"
	"github.com/d2g/logfilter/dummy"
)

func ExampleLogutils() {
	// Set the log output to just the short filename so the output doesn't contain the date time which changes each test.
	logfilter.SetFlags(log.Lshortfile)

	//Output to std out for testing (You wouldn't usually do this)
	logfilter.SetOutput(os.Stdout)

	// Don't output any messages Fatal or Lower from the package or subpackages github.com/d2g/logfilter.
	logfilter.Exclude("github.com/d2g/logfilter").When(logfilter.Fatal)

	//Output any messages Warning or Above from the package or subpackages github.com/d2g/logfilter/dummy.
	logfilter.Include("github.com/d2g/logfilter/dummy").When(logfilter.Warning)

	//Square Parser to support [Debug] type messages
	logfilter.SetParsers([]logfilter.Parser{logfilter.SqrParser})

	//Output the format in the Square Style
	logfilter.SetFormatter(logfilter.SqrFormat)

	dummy.FatalSqr()    // Create a dummy message.
	dummy.ErrorSqr()    // Create a dummy Error message.
	dummy.WarningSqr()  // Create a dummy Warning message.
	dummy.InfoSqr()     // Create a dummy Info message that should be ignored.
	dummy.DebugSqr()    // Create a dummy degub message that should be ignored.
	dummy.TraceSqr()    // Create a dummy trace message that should be ignored.
	dummy.Unformatted() //Unformatted messages appear as trace messages that should be ignored.

	//Reset the filter for the next test.
	logfilter.StdFilterReset()

	//Output:
	//dummy.go:53: [Fatal] This is a Fatal message
	//dummy.go:45: [Error] This is a Error message
	//dummy.go:37: [Warning] This is a Warning message
}
