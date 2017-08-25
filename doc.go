/*
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
		// Except for github.com/d2g/dummy which wil have Info and above.
		log.Println("Debug: Not Displayed")
		dummy.Debug()
		dummy.Info()

		//Output:
		//dummy.go:17: Info: This is a Info message
	}

If you've previously used logutils or a square based convention then look at the
example included in example_logutils_test.go

*/
package logfilter
