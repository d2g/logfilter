package logfilter

import (
	"strings"
)

//filter represents an individual log message filter.
type filter struct {
	find      string
	inclusive bool
	lvl       Level
}

//filters provides a nicer API by allowing us to create the When function.
type filters []*filter

//Default sets the logging level the is output by all packages.
func Default(lvl Level) {
	f := findFilter("")
	f.lvl = lvl
}

//findFilter is a helper function to find a filter by package name.
func findFilter(pkg string) *filter {
	for i := range stdFilters {
		if stdFilters[i].find == pkg {
			return stdFilters[i]
		}
	}
	return nil
}

//Include adds filter object(s) to the stdandard filter and returns pointers to
//the newly created object to allow you to set the required level. As default it
//sets the level to undefined (i.e. lowest).
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

//Exclude adds filter object(s) to the stdandard filter and returns pointers to
//the newly created object to allow you to set the required level. As default it
//sets the level to off (i.e. highest).
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

//When sets the level on the filters provided.
func (f filters) When(l Level) {
	for i := range f {
		f[i].lvl = l
	}
}

//StdFilterReset resets the filters being applied, which is useful for testing.
func StdFilterReset() {
	stdFilters = filters([]*filter{
		{
			find:      "",
			inclusive: true,
			lvl:       Undefined,
		},
	})
}

//StdFilter is the default implementation used by logger for filtering.
//Returns true if the line is written out.
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
