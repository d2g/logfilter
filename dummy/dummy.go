/*
Dummy package for making some noise to demonstrate github.com/d2g/logfilter
*/
package dummy

import "log"

func Trace() {
	log.Printf("Trace: This is a Trace message")
}

func TraceSqr() {
	log.Printf("[Trace] This is a Trace message")
}

func Debug() {
	log.Printf("Debug: This is a Debug message")
}

func DebugSqr() {
	log.Printf("[Debug] This is a Debug message")
}

func Info() {
	log.Printf("Info: This is a Info message")
}

func InfoSqr() {
	log.Printf("[Info] This is a Info message")
}

func Warning() {
	log.Printf("Warning: This is a Warning message")
}

func WarningSqr() {
	log.Printf("[Warning] This is a Warning message")
}

func Error() {
	log.Printf("Error: This is a Error message")
}

func ErrorSqr() {
	log.Printf("[Error] This is a Error message")
}

func Fatal() {
	log.Printf("Fatal: This is a Fatal message")
}

func FatalSqr() {
	log.Printf("[Fatal] This is a Fatal message")
}

func Unformatted() {
	log.Printf("Some package that doesn't implement the convention.")
}
