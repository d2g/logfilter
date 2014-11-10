// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package d2g/logfilter implements a filter on the standard log package.

Like most alternative logging packages using this package support the semi
standard logging levels like:
	TRACE
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL

Unlike other packages log filter doesn't require the package to specifically
import an additional package over the standard logging package. It does however
require them to follow the same convention.

The convention is:

log.Println("<Level>:<Message>")

i.e.

log.Println("Trace:Example Message")
is a trace level message.

When logging at trace level it's usual to be very verbose which results in a
large ammount of output. Writing this sort of information to a log file can
result in reduce performance (i.e. due to disk IO).

Logfilter allows you to filter this output based on the package/file location
and the log level type.

For example in your live application you may only want to output warning
messages:

func main() {
	// Set the log message output. We can only filter on the fields output so
	// it's important to output the Llongfile.
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	log.SetOutput(&logfilter.Capture{
		Flags:   log.Flags(), // We have to pass the log.Flags() here otherwise we get stuck in a loop.
		Filters: []logfilter.Filter{ //Add our Filters
		logfilter.Filter{ // Filter Out log messages from my packages except warning or above.
			Mode:     logfilter.EXCLUDE,
			Filename: "github.com/d2g/",
			Level:    logfilter.WARNING,
		},
		logfilter.Filter{ // Filter In Trace, Debug, Info from the package github.com/d2g/dhcp4server.
			Mode:     logfilter.INCLUDE,
			Filename: "github.com/d2g/dhcp4server",
			Level:    logfilter.TRACE,
		},
		},
		Output: os.Stderr, //Set the output to an implementaion of io.Writer (Stderr as default)
	})

	// Now only log level WARNING and above will be written for packages staring
	// "github.com/d2g/" except "github.com/d2g/dhcp4server" which will output
	// all messages.
}

In the example above the Output: os.Stderr does nothing (As this is the default)
however it shows how you would set the output should you require.

*/
package logfilter
