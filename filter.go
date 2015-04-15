package logfilter

import (
	"strings"
)

//
type filter struct {
	find      string
	inclusive bool
	lvl       Level
}

//
type filters []*filter

//
func Default(lvl Level) {
	f := findFilter("")
	f.lvl = lvl
}

//
func findFilter(pkg string) *filter {
	for i := range stdFilters {
		if stdFilters[i].find == pkg {
			return stdFilters[i]
		}
	}
	return nil
}

//
func Include(packagenames ...string) filters {
	n := filters{}

	for i := range packagenames {
		f := findFilter(packagenames[i])
		if f != nil {
			f.inclusive = true
			f.lvl = Undefined
			n = append(n, f)
		} else {
			f = &filter{
				find:      packagenames[i],
				inclusive: true,
				lvl:       Undefined,
			}
			n = append(n, f)
			stdFilters = append(stdFilters, f)
		}
	}

	return n
}

//
func Exclude(packagenames ...string) filters {
	n := filters{}

	for i := range packagenames {
		f := findFilter(packagenames[i])
		if f != nil {
			f.inclusive = false
			f.lvl = Off
			n = append(n, f)
		} else {
			f = &filter{
				find:      packagenames[i],
				inclusive: false,
				lvl:       Off,
			}
			n = append(n, f)
			stdFilters = append(stdFilters, f)
		}
	}

	return n
}

//
func (f filters) When(l Level) {
	for i := range f {
		f[i].lvl = l
	}
}

//
func StdFilterReset() {
	stdFilters = filters([]*filter{
		&filter{
			find:      "",
			inclusive: true,
			lvl:       Undefined,
		},
	})
}

//
func StdFilter(l *LogLine) bool {
	depth := -1
	writeout := false

	//Check for Exclusions / Inclusions
	for i := range stdFilters {

		//Does the filter apply.
		if len(stdFilters[i].find) > depth &&
			strings.Contains(l.File, stdFilters[i].find) &&
			((stdFilters[i].inclusive && stdFilters[i].lvl <= l.Level) ||
				(!stdFilters[i].inclusive && stdFilters[i].lvl >= l.Level)) {
			writeout = stdFilters[i].inclusive
			depth = len(stdFilters[i].find)
		}
	}
	return writeout
}
