package logfilter_test

import (
	"log"
	"os"

	"github.com/d2g/logfilter"
	"github.com/d2g/logfilter/dummy"
)

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

	//Reset the filter for the next test.
	logfilter.StdFilterReset()

	// Output:
	//dummy.go:49: Fatal: This is a Fatal message
	//dummy.go:41: Error: This is a Error message
	//dummy.go:33: Warning: This is a Warning message
}
