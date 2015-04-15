package logfilter

import (
	"log"
	"os"
)

var std *Logger

var stdFilters filters

func init() {
	// Setup logging the way we expect to receive log mesages
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	// Remove the prefix.
	log.SetPrefix("")
	// Setup the standard logger.
	std = &Logger{
		flag:      log.LstdFlags,
		output:    os.Stderr,
		formatter: StdFormat,
		parsers:   []Parser{StdParser},
		filter:    StdFilter,
	}
	// Set the output of the standard log package to our logger.
	log.SetOutput(std)

	//
	StdFilterReset()
}
