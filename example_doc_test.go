package logfilter_test

import (
	"log"
	"os"

	"github.com/d2g/logfilter"
	"github.com/d2g/logfilter/dummy"
)

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

	// Use "[WARN] " style prefixes, like logutils
	// logfilter.SetParsers([]logfilter.Parser{logfilter.SqrParser})

	// Now only log level Warning and above will be written
	// Except for github.com/d2g/dummy which wil have Info and above.
	log.Println("Debug: Not Displayed")
	dummy.Debug()
	dummy.Info()

	//Reset the filter for the next test.
	logfilter.StdFilterReset()

	// Output:
	//dummy.go:25: Info: This is a Info message
}
