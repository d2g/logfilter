package logfilter

import (
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Capture struct {
	Flags   int
	Output  io.Writer
	Filters []Filter
}

func (t *Capture) Parse(p string) (l *Line) {
	l = &Line{}

	remainder := p

	if t.Flags&(log.Ldate) != 0 {

		tmpParts := strings.SplitN(p, " ", 3)

		if t.Flags&(log.Ldate|log.Ltime|log.Lmicroseconds) != 0 ||
			t.Flags&(log.Ldate|log.Lmicroseconds) != 0 ||
			t.Flags&(log.Ldate|log.Ltime) != 0 {

			l.Timestamp, _ = time.Parse("2006/01/02 15:04:05.999999", tmpParts[0]+" "+tmpParts[1])
		}

		remainder = tmpParts[2]

	} else if t.Flags&(log.Ltime|log.Lmicroseconds) != 0 {

		tmpParts := strings.SplitN(p, " ", 2)

		l.Timestamp, _ = time.Parse("15:04:05.999999", tmpParts[0])

		remainder = tmpParts[1]

	}

	if t.Flags&(log.Llongfile|log.Lshortfile) != 0 {
		tmpParts := strings.SplitN(remainder, ": ", 2)
		l.FileAndLine = tmpParts[0]
		remainder = tmpParts[1]
	}

	//Assume the worst For Every Message:
	l.Level = FATAL

	//We need to see if our log files follow the convenstion.
	if len(remainder) >= 6 && strings.ToUpper(remainder[0:6]) == "TRACE:" {
		l.Level = TRACE
		remainder = remainder[6:]
	} else if len(remainder) >= 6 && strings.ToUpper(remainder[0:6]) == "DEBUG:" {
		l.Level = DEBUG
		remainder = remainder[6:]
	} else if len(remainder) >= 5 && strings.ToUpper(remainder[0:5]) == "INFO:" {
		l.Level = INFO
		remainder = remainder[5:]
	} else if len(remainder) >= 8 && strings.ToUpper(remainder[0:8]) == "WARNING:" {
		l.Level = WARNING
		remainder = remainder[8:]
	} else if len(remainder) >= 6 && strings.ToUpper(remainder[0:6]) == "ERROR:" {
		l.Level = ERROR
		remainder = remainder[6:]
	} else if len(remainder) >= 6 && strings.ToUpper(remainder[0:6]) == "FATAL:" {
		l.Level = FATAL
		remainder = remainder[6:]
	}

	l.Message = strings.Trim(remainder, " ")

	return
}

func (t *Capture) Write(p []byte) (int, error) {
	l := t.Parse(string(p))

	//The more exact(Length) the filter the more this applies to this
	depth := 0
	writeout := true

	//Check for Exclusions / Inclusions
	for i := range t.Filters {
		//Does the filter apply.
		if len(t.Filters[i].Filename) > depth &&
			strings.Contains(l.FileAndLine, t.Filters[i].Filename) &&
			((t.Filters[i].Mode == INCLUDE && t.Filters[i].Level <= l.Level) || (t.Filters[i].Mode == EXCLUDE && t.Filters[i].Level >= l.Level)) {

			depth = len(t.Filters[i].Filename)

			switch t.Filters[i].Mode {
			case EXCLUDE:
				writeout = false
			case INCLUDE:
				writeout = true
			}
		}

	}

	if writeout {
		if t.Output != nil {
			return t.Output.Write([]byte(p))
		} else {
			return os.Stderr.Write([]byte(p))
		}
	}

	return 0, io.EOF
}
